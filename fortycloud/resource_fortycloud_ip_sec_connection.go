package fortycloud

import (
	"fmt"
	fc "github.com/bsick7/fortycloud-sdk-go/api"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFcIPSecConnection() *schema.Resource {
	return &schema.Resource{
		Create: resourceFcIPSecConnectionCreate,
		Read:   resourceFcIPSecConnectionRead,
		Update: resourceFcIPSecConnectionUpdate,
		Delete: resourceFcIPSecConnectionDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gateway_a": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"gateway_b": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"enable": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func resourceFcIPSecConnectionCreate(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*fc.Api)
	newConn := &fc.IPSecConnection{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		GatewayA:    d.Get("gateway_a").(string),
		GatewayB:    d.Get("gateway_b").(string),
		Enable:      true,
	}

	if val, ok := d.GetOk("enable"); ok {
		newConn.Enable = val.(bool)
	}

	conn, err := api.IPSecConnections.Create(newConn)
	if err != nil {
		return fmt.Errorf("Error creating connection: %s", err)
	}

	d.SetId(conn.Id)
	return nil
}

func resourceFcIPSecConnectionRead(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*fc.Api)
	id := d.Id()

	conn, err := api.IPSecConnections.Get(id)
	if err != nil {
		return fmt.Errorf("Could not retrieve connection: %s", err)
	}

	d.Set("name", conn.Name)
	d.Set("description", conn.Description)
	d.Set("gateway_a", conn.GatewayA)
	d.Set("gateway_b", conn.GatewayB)
	d.Set("enable", conn.Enable)

	return nil
}

func resourceFcIPSecConnectionUpdate(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*fc.Api)
	id := d.Id()

	conn, err := api.IPSecConnections.Get(id)
	if err != nil {
		return fmt.Errorf("Could not retrieve connection: %s", err)
	}

	conn.Name = d.Get("name").(string)
	conn.Description = d.Get("description").(string)
	conn.GatewayA = d.Get("gateway_a").(string)
	conn.GatewayB = d.Get("gateway_b").(string)

	conn.Enable = true
	if val, ok := d.GetOk("enable"); ok {
		conn.Enable = val.(bool)
	}

	if _, err := api.IPSecConnections.Update(id, conn); err != nil {
		return fmt.Errorf("Could not update connection: %s", err)
	}

	return nil
}

func resourceFcIPSecConnectionDelete(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*fc.Api)
	if d.Id() == "" {
		return nil
	}
	id := d.Id()

	if err := api.IPSecConnections.Delete(id); err != nil {
		return fmt.Errorf("Could not delete connection: %s", err)
	}

	return nil
}
