package fortycloud

import (
	"fmt"
	fc "github.com/BSick7/fortycloud-sdk-go/api"
	"github.com/hashicorp/terraform/helper/schema"
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

	subnet := &fc.Subnet{
		Name:           d.Get("name").(string),
		Description:    d.Get("description").(string),
		Cidr:           d.Get("cidr").(string),
		DisableAutoNAT: d.Get("disable_auto_nat").(bool),
	}
	subnet.SetGatewayID(d.Get("gateway_id").(string))
	subnet.SetResourceGroupID(d.Get("resource_group_id").(string))

	existing, err := api.FindSubnet(subnet.Cidr, subnet.GatewayID())
	if err != nil {
		return fmt.Errorf("error matching with existing subnet: %s", err)
	}
	if existing != nil {
		// if we find an existing subnet matched on cidr and gateway id
		// it was probably created by the gateway when registered
		d.SetId(existing.Id)
	} else {
		newSubnet, err := api.Subnets.Create(subnet)
		if err != nil {
			return fmt.Errorf("error creating subnet: %s", err)
		}
		d.SetId(newSubnet.Id)
	}

	return resourceFcSubnetUpdate(d, meta)
}

func resourceFcSubnetRead(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*fc.Api)

	if d.Id() == "" {
		return nil
	}

	subnet, err := api.Subnets.Get(d.Id())
	if err != nil {
		return fmt.Errorf("error retrieving subnet: %s", err)
	}
	if subnet == nil {
		d.MarkNewResource()
		d.SetId("")
		return nil
	}

	d.Set("name", subnet.Name)
	d.Set("description", subnet.Description)
	d.Set("cidr", subnet.Cidr)
	d.Set("disable_auto_nat", subnet.DisableAutoNAT)
	d.Set("gateway_id", subnet.GatewayID())
	d.Set("resource_group_id", subnet.ResourceGroupID())

	return nil
}

func resourceFcSubnetUpdate(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*fc.Api)

	if d.Id() == "" {
		return nil
	}

	subnet := &fc.Subnet{
		Id:             d.Id(),
		Name:           d.Get("name").(string),
		Description:    d.Get("description").(string),
		Cidr:           d.Get("cidr").(string),
		DisableAutoNAT: d.Get("disable_auto_nat").(bool),
	}
	subnet.SetGatewayID(d.Get("gateway_id").(string))
	subnet.SetResourceGroupID(d.Get("resource_group_id").(string))

	if _, err3 := api.Subnets.Update(subnet.Id, subnet); err3 != nil {
		return fmt.Errorf("error updating subnet: %s", err3)
	}

	// update gateway on subnet
	if _, err := api.Subnets.AssignGateway(subnet.Id, subnet.GatewayID()); err != nil {
		return fmt.Errorf("error assigning gateway to subnet: %s", err)
	}

	// update resource group on subnet
	if _, err := api.Subnets.AssignResourceGroup(subnet.Id, subnet.ResourceGroupID()); err != nil {
		return fmt.Errorf("error assigning resource group to subnet: %s", err)
	}

	return nil
}

func resourceFcSubnetDelete(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*fc.Api)

	if d.Id() == "" {
		return nil
	}

	if err := api.Subnets.Delete(d.Id()); err != nil {
		return fmt.Errorf("error deleting subnet: %s", err)
	}

	return nil
}
