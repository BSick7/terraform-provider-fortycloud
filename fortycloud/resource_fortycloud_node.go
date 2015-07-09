package fortycloud

import (
	"fmt"
	"strconv"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mdl/fortycloud-sdk-go/api"
)

func resourceFcNode() *schema.Resource {
	return &schema.Resource{
		Create: resourceFcNodeCreate,
		Read:   resourceFcNodeRead,
		Update: resourceFcNodeUpdate,
		Delete: resourceFcNodeDelete,
		Exists: resourceFcNodeExists,
		
		Schema: map[string]*schema.Schema{
			"public_ip": &schema.Schema{
				Type:		schema.TypeString,
				Required:	true,
			},
			"peer_id": &schema.Schema{
				Type:		schema.TypeInt,
				Computed:	true,
			},
		},
	}
}

func resourceFcNodeCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceFcNodeRead(d, meta)
}

func resourceFcNodeRead(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*fortycloud.Api)
	node, err := api.Nodes.GetByPublicIP(d.Get("public_ip").(string))
	if err != nil {
		return fmt.Errorf("Error getting node: %s", err)
	}
	if node == nil {
		return fmt.Errorf("Node does not exist.")
	}
	d.Set("peer_id", node.Id)
	d.SetId(strconv.Itoa(node.Id))
	return nil
}

func resourceFcNodeUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceFcNodeRead(d, meta)
}

func resourceFcNodeDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceFcNodeExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	api := meta.(*fortycloud.Api)
	node, err := api.Nodes.GetByPublicIP(d.Get("public_ip").(string))
	if err != nil {
		return false, fmt.Errorf("Error getting node: %s", err)
	}
	return node != nil, nil
}