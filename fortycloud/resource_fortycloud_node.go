package fortycloud

import (
	"fmt"
	"log"
	"strconv"
	"time"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mdl/fortycloud-sdk-go/api"
)

type MissingNodeError struct {
	PublicIP string
}
func (err *MissingNodeError) Error() string {
	return fmt.Sprintf("Could not find node with Public IP=%s.", err.PublicIP)
}

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
	retryDelay := 15
	retries := 90 / retryDelay
	publicip := d.Get("public_ip").(string)
	for i :=0; i < retries; i++ {
		err := resourceFcNodeRead(d, meta)
		if _, ok := err.(*MissingNodeError); ok {
			if (i+1) == retries {
				return err
			}
			log.Printf("Waiting %d seconds to locate node (%s)...\n", retryDelay, publicip)
			time.Sleep(time.Duration(retryDelay) * time.Second)
			continue
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func resourceFcNodeRead(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*fortycloud.Api)
	node, err := api.Nodes.GetByPublicIP(d.Get("public_ip").(string))
	if err != nil {
		return fmt.Errorf("Error getting node: %s", err)
	}
	if node == nil {
		return &MissingNodeError{PublicIP:d.Get("public_ip").(string)}
	}
	d.Set("peer_id", node.Id)
	d.SetId(strconv.Itoa(node.Id))
	return nil
}

func resourceFcNodeUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceFcNodeRead(d, meta)
}

func resourceFcNodeDelete(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*fortycloud.Api)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("Could not get node id: %s", err)
	}
	
	if err := api.Nodes.Delete(id); err != nil {
		return fmt.Errorf("Could not delete node: %s", err)
	}
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