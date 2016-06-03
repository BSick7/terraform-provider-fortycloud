package fortycloud

import (
	"fmt"
	fc "github.com/BSick7/fortycloud-sdk-go/api"
	"github.com/hashicorp/terraform/helper/schema"
	"strings"
)

func resourceFcSubnet() *schema.Resource {
	return &schema.Resource{
		Create: resourceFcSubnetCreate,
		Read:   resourceFcSubnetRead,
		Update: resourceFcSubnetUpdate,
		Delete: resourceFcSubnetDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cidr": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"disable_auto_nat": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"gateway_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_group_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceFcSubnetCreate(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*fc.Api)

	newSubnet := &fc.Subnet{
		Name:           d.Get("name").(string),
		Description:    d.Get("description").(string),
		Cidr:           d.Get("cidr").(string),
		DisableAutoNAT: d.Get("disable_auto_nat").(bool),
	}

	subnet, err := api.Subnets.Create(newSubnet)
	if err != nil {
		return fmt.Errorf("Error creating subnet: %s", err)
	}
	d.SetId(subnet.Id)

	// Assign gateway to subnet
	gw_id := d.Get("gateway_id").(string)
	res_gw_id, err := assignSubnetGateway(api, subnet.Id, gw_id)
	if err != nil {
		return fmt.Errorf("Could not assign gateway to subnet: %s", err)
	}
	d.Set("gateway_id", res_gw_id)

	// Assign resource group to subnet
	rg_id := d.Get("resource_group_id").(string)
	res_rg_id, err := assignSubnetResourceGroup(api, subnet.Id, rg_id)
	if err != nil {
		return fmt.Errorf("Could not assign resource group to subnet: %s", err)
	}
	d.Set("resource_group_id", res_rg_id)

	return nil
}

func resourceFcSubnetRead(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*fc.Api)
	id := d.Id()

	subnet, err := api.Subnets.Get(id)
	if err != nil {
		return fmt.Errorf("Could not retrieve subnet: %s", err)
	}

	d.Set("name", subnet.Name)
	d.Set("description", subnet.Description)
	d.Set("cidr", subnet.Cidr)
	d.Set("disable_auto_nat", subnet.DisableAutoNAT)
	d.Set("gateway_id", parseGatewayIdFromRef(subnet.GatewayRef))
	d.Set("resource_group_id", parseResourceGroupIdFromRef(subnet.ResourceGroupRef))

	return nil
}

func resourceFcSubnetUpdate(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*fc.Api)
	id := d.Id()

	subnet, err := api.Subnets.Get(id)
	if err != nil {
		return fmt.Errorf("Could not retrieve subnet: %s", err)
	}

	subnet.Name = d.Get("name").(string)
	subnet.Description = d.Get("description").(string)
	subnet.Cidr = d.Get("cidr").(string)
	subnet.DisableAutoNAT = d.Get("disable_auto_nat").(bool)

	_, err3 := api.Subnets.Update(id, subnet)
	if err3 != nil {
		return fmt.Errorf("Could not update subnet: %s", err3)
	}

	// Update gateway on subnet
	cur_gw_id := parseGatewayIdFromRef(subnet.GatewayRef)
	new_gw_id := d.Get("gateway_id").(string)
	if cur_gw_id != new_gw_id {
		gw_id, err := assignSubnetGateway(api, id, new_gw_id)
		if err != nil {
			return fmt.Errorf("Could not update subnet gateway: %s", err)
		}
		d.Set("gateway_id", gw_id)
	}

	// Update resource group on subnet
	cur_rg_id := parseResourceGroupIdFromRef(subnet.ResourceGroupRef)
	new_rg_id := d.Get("resource_group_id").(string)
	if cur_rg_id != new_rg_id {
		rg_id, err := assignSubnetResourceGroup(api, id, new_rg_id)
		if err != nil {
			return fmt.Errorf("Could not update subnet resource group: %s", err)
		}
		d.Set("resource_group_id", rg_id)
	}

	return nil
}

func resourceFcSubnetDelete(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*fc.Api)
	id := d.Id()
	if len(id) == 0 {
		return nil
	}

	if err := api.Subnets.Delete(id); err != nil {
		return fmt.Errorf("Could not delete subnet: %s", err)
	}

	return nil
}

func assignSubnetGateway(api *fc.Api, subnet_id string, gw_id string) (string, error) {
	if len(gw_id) <= 0 {
		return "", nil
	}
	mod, err := api.Subnets.AssignGateway(subnet_id, gw_id)
	if err != nil {
		return "", fmt.Errorf("Could not update subnet's gateway: %s", err)
	}
	return parseGatewayIdFromRef(mod.GatewayRef), nil
}

func assignSubnetResourceGroup(api *fc.Api, subnet_id string, rg_id string) (string, error) {
	if len(rg_id) <= 0 {
		return "", nil
	}
	mod, err := api.Subnets.AssignResourceGroup(subnet_id, rg_id)
	if err != nil {
		return "", fmt.Errorf("Could not update subnet's resource group: %s", err)
	}
	return parseResourceGroupIdFromRef(mod.ResourceGroupRef), nil
}

func parseGatewayIdFromRef(ref string) string {
	// Expected input: https://api.fortycloud.net/restapi/v0.4/gateways/4013
	tokens := strings.Split(ref, "/")
	if len(tokens) < 2 || tokens[len(tokens)-2] != "gateways" {
		return ""
	}
	return tokens[len(tokens)-1]
}

func parseResourceGroupIdFromRef(ref string) string {
	// Expected input: https://api.fortycloud.net/restapi/v0.4/resource-groups/2069
	tokens := strings.Split(ref, "/")
	if len(tokens) < 2 || tokens[len(tokens)-2] != "resource-groups" {
		return ""
	}
	return tokens[len(tokens)-1]
}
