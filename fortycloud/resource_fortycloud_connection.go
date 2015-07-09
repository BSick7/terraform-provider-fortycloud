package fortycloud

import (
	"fmt"
	"strconv"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mdl/fortycloud-sdk-go/api"
	"github.com/mdl/fortycloud-sdk-go/forms"
)

func resourceFcConnection() *schema.Resource {
	return &schema.Resource{
		Create: resourceFcConnectionCreate,
		Read:   resourceFcConnectionRead,
		Update: resourceFcConnectionUpdate,
		Delete: resourceFcConnectionDelete,
		
		Schema: map[string]*schema.Schema{
			"peer_a_id": &schema.Schema{
				Type:       schema.TypeInt,
				Required:   true,
			},
			"peer_b_id": &schema.Schema{
				Type:		schema.TypeInt,
				Required:	true,
			},
		},
	}
}

func resourceFcConnectionCreate(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*fortycloud.Api)
	newConn := &forms.Connection{
		PeerA: forms.Peer{
			Id: d.Get("peer_a_id").(int),
		},
		PeerB: forms.Peer{
			Id: d.Get("peer_b_id").(int),
		},
	}
	
	conn, err := api.Connections.Create(newConn)
	if err != nil {
		return fmt.Errorf("Error creating connection: %s", err)
	}
	
	d.SetId(strconv.Itoa(conn.Id))
	return nil
}

func resourceFcConnectionRead(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*fortycloud.Api)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("Could not retrieve id: %s", err)
	}
	
	conn, err2 := api.Connections.Get(id)
	if err2 != nil {
		return fmt.Errorf("Could not retrieve connection: %s", err)
	}
	
	d.Set("peer_a_id", conn.PeerA.Id)
	d.Set("peer_b_id", conn.PeerB.Id)
	
	return nil
}

func resourceFcConnectionUpdate(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*fortycloud.Api)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("Could not retrieve id: %s", err)
	}
	
	conn, err2 := api.Connections.Get(id)
	if err2 != nil {
		return fmt.Errorf("Could not retrieve connection: %s", err2)
	}
	
	conn.PeerA = forms.Peer{Id: d.Get("peer_a_id").(int)}
	conn.PeerB = forms.Peer{Id: d.Get("peer_b_id").(int)}
	_, err3 := api.Connections.Update(conn)
	if err3 != nil {
		return fmt.Errorf("Could not update connection: %s", err3)
	} 
	
	return nil
}

func resourceFcConnectionDelete(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*fortycloud.Api)
	if d.Id() == "" {
		return nil
	}
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("Could not retrieve id: %s", err)
	}
	
	err = api.Connections.Delete(id)
	if err != nil {
		return fmt.Errorf("Could not delete connection: %s", err)
	}
	
	return nil
}