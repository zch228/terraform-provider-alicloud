package alicloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/terraform-providers/terraform-provider-alicloud/alicloud/connectivity"
)

func TestAccAlicloudEdasDeployGroupDataSource(t *testing.T) {
	rand := acctest.RandIntRange(1000, 9999)
	resourceId := "data.alicloud_edas_deploy_groups.default"
	name := fmt.Sprintf("tf-testacc-edas-deploy-groups%v", rand)

	testAccConfig := dataSourceTestAccConfigFunc(resourceId, name, dataSourceEdasDeployGroupConfigDependence)

	allConf := dataSourceTestAccConfig{
		existConfig: testAccConfig(map[string]interface{}{
			"ids":    []string{"${alicloud_edas_deploy_group.default.id}"},
			"app_id": "${alicloud_edas_application.default.id}",
		}),
		fakeConfig: testAccConfig(map[string]interface{}{
			"ids":    []string{"${alicloud_edas_deploy_group.default.id}_fake"},
			"app_id": "${alicloud_edas_application.default.id}_fake",
		}),
	}
	var existEdasDeployGroupsMapFunc = func(rand int) map[string]string {
		return map[string]string{
			"groups.#":             "1",
			"groups.0.group_name":  fmt.Sprintf("tf-testacc-edas-deploy-groups%v", rand),
			"groups.0.app_id":      CHECKSET,
			"groups.0.group_type":  CHECKSET,
			"groups.0.cluster_id":  CHECKSET,
			"groups.0.create_time": CHECKSET,
			"groups.0.update_time": CHECKSET,
			"groups.0.group_id":    CHECKSET,
		}
	}

	var fakeEdasDeployGroupsMapFunc = func(rand int) map[string]string {
		return map[string]string{
			"groups.#": "0",
		}
	}

	var edasApplicationCheckInfo = dataSourceAttr{
		resourceId:   resourceId,
		existMapFunc: existEdasDeployGroupsMapFunc,
		fakeMapFunc:  fakeEdasDeployGroupsMapFunc,
	}

	preCheck := func() {
		testAccPreCheckWithRegions(t, true, connectivity.EdasSupportedRegions)
	}

	edasApplicationCheckInfo.dataSourceTestCheckWithPreCheck(t, rand, preCheck, allConf)
}

func dataSourceEdasDeployGroupConfigDependence(name string) string {
	return fmt.Sprintf(`
		variable "name" {
		 default = "%v"
		}

		resource "alicloud_vpc" "default" {
		  cidr_block = "172.16.0.0/12"
		  name       = "${var.name}"
		}

		resource "alicloud_edas_cluster" "default" {
		  cluster_name = "${var.name}"
		  cluster_type = 2
		  network_mode = 2
		  vpc_id       = "${alicloud_vpc.default.id}"
		}

		resource "alicloud_edas_application" "default" {
		  application_name = "${var.name}"
		  cluster_id = alicloud_edas_cluster.default.id
		  package_type = "WAR"
		}
		
		resource "alicloud_edas_deploy_group" "default" {
		  app_id = alicloud_edas_application.default.id
		  group_name = "${var.name}"
		}		
		`, name)
}
