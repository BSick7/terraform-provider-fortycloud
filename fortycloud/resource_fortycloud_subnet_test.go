package fortycloud

import (
	"fmt"
	fc "github.com/bsick7/fortycloud-sdk-go/api"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"testing"
)

func TestAccFCSubnet_basic(t *testing.T) {
	var subnet fc.Subnet

	testCheck := func(s *terraform.State) error {
		if subnet.Cidr != "10.1.0.0/16" {
			t.Errorf("Cidr (Expected=%s, Actual=%s)\n", "10.1.0.0/16", subnet.Cidr)
		}
		return nil
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFCSubnetDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccFortyCloudSubnetBasicConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFCSubnetExists("fortycloud_subnet.subnet1", &subnet),
					testCheck,
				),
			},
		},
	})
}

func testAccCheckFCSubnetExists(n string, res *fc.Subnet) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s\n", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		api := testAccProvider.Meta().(*fc.Api)
		subnet, err := api.Subnets.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		if subnet == nil {
			return fmt.Errorf("Subnet not found")
		}

		*res = *subnet

		return nil
	}
}

func testAccCheckFCSubnetDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortycloud_subnet" {
			continue
		}

		api := testAccProvider.Meta().(*fc.Api)
		subnets, err := api.Subnets.All()
		if err != nil {
			return err
		}

		for _, subnet := range subnets {
			if subnet.Id == rs.Primary.ID {
				return fmt.Errorf("Subnet was not destroyed. %+v", subnet)
			}
		}
	}

	return nil
}

var testAccFortyCloudSubnetBasicConfig = fmt.Sprint(`
resource "fortycloud_subnet" "subnet1" {
  name = "subnet1-testacc"
  description = "Subnet 1 (Acceptance Test)"
  cidr = "10.1.0.0/16"
}
`)
