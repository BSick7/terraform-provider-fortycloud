package fortycloud

import (
	"fmt"
	"strconv"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mdl/fortycloud-sdk-go/api"
	"github.com/mdl/fortycloud-sdk-go/forms"
)

func resourceFcSubnet() *schema.Resource {
	return &schema.Resource{
		Create: resourceFcSubnetCreate,
		Read:   resourceFcSubnetRead,
		Update: resourceFcSubnetUpdate,
		Delete: resourceFcSubnetDelete,
		
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:		schema.TypeString,
				Required:	true,
			},
			"description": &schema.Schema{
				Type:		schema.TypeString,
				Optional:	true,
			},
			"source": &schema.Schema{
				Type:		schema.TypeString,
				Optional:	true,
			},
			"subnet": &schema.Schema{
				Type:		schema.TypeString,
				Required:	true,
			},
			"actual_subnet": &schema.Schema{
				Type:		schema.TypeString,
				Optional:	true,
			},
			"nat_disabled": &schema.Schema{
				Type:		schema.TypeBool,
				Optional:	true,
			},
		},
	}
}

func resourceFcSubnetCreate(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*fortycloud.Api)
	
	newSubnet := &forms.PrivateSubnet{
		Name: d.Get("name").(string),
		Description: d.Get("description").(string),
		Source: d.Get("source").(string),
		Subnet: d.Get("subnet").(string),
		ActualSubnet: d.Get("actual_subnet").(string),
		NatDisabled: d.Get("nat_disabled").(bool),
	}
	
	subnet, err := api.PrivateSubnets.Create(newSubnet)
	if err != nil {
		return fmt.Errorf("Error creating subnet: %s", err)
	}
	
	d.SetId(strconv.Itoa(subnet.Id))
	return nil
}

func resourceFcSubnetRead(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*fortycloud.Api)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("Could not retrieve id: %s", err)
	}
	
	subnet, err2 := api.PrivateSubnets.Get(id)
	if err2 != nil {
		return fmt.Errorf("Could not retrieve subnet: %s", err2)
	}
	
	d.Set("name", subnet.Name)
	d.Set("description", subnet.Description)
	d.Set("source", subnet.Source)
	d.Set("subnet", subnet.Subnet)
	d.Set("actual_subnet", subnet.ActualSubnet)
	d.Set("nat_disabled", subnet.NatDisabled)

	return nil
}

func resourceFcSubnetUpdate(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*fortycloud.Api)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("Could not retrieve id: %s", err)
	}
	
	subnet, err2 := api.PrivateSubnets.Get(id)
	if err2 != nil {
		return fmt.Errorf("Could not retrieve subnet: %s", err2)
	}
	
	subnet.Name = d.Get("name").(string)
	subnet.Description = d.Get("description").(string)
	subnet.Source = d.Get("source").(string)
	subnet.Subnet = d.Get("subnet").(string)
	subnet.ActualSubnet = d.Get("actual_subnet").(string)
	subnet.NatDisabled = d.Get("nat_disabled").(bool)
	
	_, err3 := api.PrivateSubnets.Update(subnet)
	if err3 != nil {
		return fmt.Errorf("Could not update subnet: %s", err3)
	} 
	
	return nil
}

func resourceFcSubnetDelete(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*fortycloud.Api)
	if d.Id() == "" {
		return nil
	}
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("Could not retrieve id: %s", err)
	}
	
	err = api.PrivateSubnets.Delete(id)
	if err != nil {
		return fmt.Errorf("Could not delete subnet: %s", err)
	}
	
	return nil
}