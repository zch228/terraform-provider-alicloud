package alicloud

import (
	"strings"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/edas"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/terraform-providers/terraform-provider-alicloud/alicloud/connectivity"
)

func resourceAlicloudEdasDeployGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlicloudEdasDeployGroupCreate,
		Read:   resourceAlicloudEdasDeployGroupRead,
		Delete: resourceAlicloudEdasDeployGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"app_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"group_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"group_type": {
				Type:     schema.TypeInt,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAlicloudEdasDeployGroupCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	edasService := EdasService{client}

	appId := d.Get("app_id").(string)
	regionId := client.RegionId
	groupName := d.Get("group_name").(string)

	request := edas.CreateInsertDeployGroupRequest()
	request.RegionId = regionId
	request.AppId = appId
	request.GroupName = groupName

	err := resource.Retry(5*time.Minute, func() *resource.RetryError {
		raw, err := edasService.client.WithEdasClient(func(edasClient *edas.Client) (interface{}, error) {
			return edasClient.InsertDeployGroup(request)
		})
		if err != nil {
			if IsExpectedErrors(err, []string{ThrottlingUser}) {
				time.Sleep(10 * time.Second)
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		response := raw.(*edas.InsertDeployGroupResponse)
		deployGroup := response.DeployGroupEntity
		d.SetId(appId + ":" + groupName + ":" + deployGroup.Id)
		addDebug(request.GetActionName(), raw, request.RoaRequest, request)
		return nil
	})

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_edas_deploy_group", request.GetActionName(), AlibabaCloudSdkGoERROR)
	}

	return resourceAlicloudEdasDeployGroupRead(d, meta)
}

func resourceAlicloudEdasDeployGroupRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	edasService := EdasService{client}

	strs := strings.Split(d.Id(), ":")

	if len(strs) != 3 {
		return WrapError(Error("resource id decode failed: " + d.Id()))
	}

	appId := strs[0]
	groupId := strs[2]

	deployGroup, err := edasService.GetDeployGroup(appId, groupId)
	if err != nil {
		return WrapError(err)
	}

	d.Set("group_type", deployGroup.GroupType)
	d.Set("app_id", deployGroup.AppId)
	d.Set("group_name", deployGroup.GroupName)

	return nil
}

func resourceAlicloudEdasDeployGroupDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	edasService := EdasService{client}

	appId := d.Get("app_id").(string)
	regionId := client.RegionId
	groupName := d.Get("group_name").(string)

	request := edas.CreateDeleteDeployGroupRequest()
	request.RegionId = regionId
	request.AppId = appId
	request.GroupName = groupName

	err := resource.Retry(5*time.Minute, func() *resource.RetryError {
		raw, err := edasService.client.WithEdasClient(func(edasClient *edas.Client) (interface{}, error) {
			return edasClient.DeleteDeployGroup(request)
		})
		if err != nil {
			if IsExpectedErrors(err, []string{ThrottlingUser}) {
				time.Sleep(10 * time.Second)
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		addDebug(request.GetActionName(), raw, request.RoaRequest, request)
		return nil
	})
	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), request.GetActionName(), AlibabaCloudSdkGoERROR)
	}

	return nil
}
