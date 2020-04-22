package alicloud

import (
	"strconv"
	"time"

	"strings"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/edas"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/terraform-providers/terraform-provider-alicloud/alicloud/connectivity"
)

func resourceAlicloudEdasApplicationPackageAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlicloudEdasApplicationPackageAttachmentCreate,
		Read:   resourceAlicloudEdasApplicationPackageAttachmentRead,
		Delete: resourceAlicloudEdasApplicationPackageAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"app_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"group_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"package_version": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"war_url": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"last_package_version": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAlicloudEdasApplicationPackageAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	edasService := EdasService{client}

	regionId := client.RegionId
	appId := d.Get("app_id").(string)
	packageVersion := d.Get("package_version").(string)
	if len(packageVersion) == 0 {
		packageVersion = strconv.Itoa(time.Now().Second())
	}
	groupId := d.Get("group_id").(string)
	warUlr := d.Get("war_url").(string)

	if version, err := edasService.GetLastPackgeVersion(appId, groupId); err != nil {
		return err
	} else {
		d.Set("last_package_version", version)
	}

	request := edas.CreateDeployApplicationRequest()
	request.RegionId = regionId
	request.AppId = appId
	request.PackageVersion = packageVersion
	request.DeployType = "url"
	request.WarUrl = warUlr
	request.GroupId = groupId

	raw, err := edasService.client.WithEdasClient(func(edasClient *edas.Client) (interface{}, error) {
		return edasClient.DeployApplication(request)
	})
	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), request.GetActionName(), AlibabaCloudSdkGoERROR)
	}
	addDebug(request.GetActionName(), raw, request.RoaRequest, request)

	response, _ := raw.(*edas.DeployApplicationResponse)
	changeOrderId := response.ChangeOrderId
	if response.Code != 200 {
		return Error("deploy application failed for " + response.Message)
	}

	if len(changeOrderId) > 0 {
		stateConf := BuildStateConf([]string{"0", "1"}, []string{"2"}, d.Timeout(schema.TimeoutCreate), 5*time.Second, edasService.EdasChangeOrderStatusRefreshFunc(changeOrderId, []string{"3", "6", "10"}))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
		}
	}
	d.SetId(appId + ":" + packageVersion)

	return resourceAlicloudEdasApplicationPackageAttachmentRead(d, meta)
}

func resourceAlicloudEdasApplicationPackageAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	edasService := EdasService{client}

	regionId := client.RegionId
	appId := strings.Split(d.Id(), ":")[0]

	request := edas.CreateQueryApplicationStatusRequest()
	request.RegionId = regionId
	request.AppId = appId

	raw, err := edasService.client.WithEdasClient(func(edasClient *edas.Client) (interface{}, error) {
		return edasClient.QueryApplicationStatus(request)
	})
	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_edas_application_package_attachment", request.GetActionName(), AlibabaCloudSdkGoERROR)
	}
	addDebug(request.GetActionName(), raw, request.RoaRequest, request)
	response, _ := raw.(*edas.QueryApplicationStatusResponse)

	if response.Code != 200 {
		return Error("QueryApplicationStatus failed for " + response.Message)
	}

	groupId := d.Get("group_id").(string)
	for _, group := range response.AppInfo.GroupList.Group {
		if group.GroupId == groupId {
			d.SetId(appId + ":" + group.PackageVersionId)
		}
	}

	return nil
}

func resourceAlicloudEdasApplicationPackageAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	edasService := EdasService{client}

	regionId := client.RegionId
	appId := d.Get("app_id").(string)
	packageVersion := d.Get("last_package_version").(string)
	groupId := d.Get("group_id").(string)

	if len(packageVersion) == 0 {
		return nil
	}

	request := edas.CreateRollbackApplicationRequest()
	request.RegionId = regionId
	request.AppId = appId
	request.HistoryVersion = packageVersion
	request.GroupId = groupId

	raw, err := edasService.client.WithEdasClient(func(edasClient *edas.Client) (interface{}, error) {
		return edasClient.RollbackApplication(request)
	})
	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), request.GetActionName(), AlibabaCloudSdkGoERROR)
	}
	addDebug(request.GetActionName(), raw, request.RoaRequest, request)

	response, _ := raw.(*edas.RollbackApplicationResponse)
	changeOrderId := response.ChangeOrderId
	if response.Code != 200 && !strings.Contains(response.Message, "ex.app.deploy.group.empty") {
		return Error("deploy application failed for " + response.Message)
	}

	if len(changeOrderId) > 0 {
		stateConf := BuildStateConf([]string{"0", "1"}, []string{"2"}, d.Timeout(schema.TimeoutCreate), 5*time.Second, edasService.EdasChangeOrderStatusRefreshFunc(changeOrderId, []string{"3", "6", "10"}))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
		}
	}

	return nil
}
