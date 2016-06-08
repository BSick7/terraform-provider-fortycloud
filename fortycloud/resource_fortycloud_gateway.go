package fortycloud

import (
	"fmt"
	fc "github.com/BSick7/fortycloud-sdk-go/api"
	"github.com/hashicorp/terraform/helper/schema"
)

type MissingGatewayError struct {
	PublicIP string
}

func (err MissingGatewayError) Error() string {
	return fmt.Sprintf("could not find gateway with Public IP=%s.", err.PublicIP)
}

func resourceFcGateway() *schema.Resource {
	return &schema.Resource{
		Create: resourceFcGatewayCreate,
		Read:   resourceFcGatewayRead,
		Update: resourceFcGatewayUpdate,
		Delete: resourceFcGatewayDelete,
		Exists: resourceFcGatewayExists,

		Schema: map[string]*schema.Schema{
			"public_ip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"allow_ssh_to_everyone": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
				Optional: true,
			},
			"gateway_as_dns": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
				Optional: true,
			},
			"identity_server_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"open_vpn_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"enable": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
				Optional: true,
			},
			"direct_routes_only": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
				Optional: true,
			},
			"fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"overlay_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vpn_users_subnet": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"release": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"route_all_traffic_via_gw": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"private_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"security_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ha_state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceFcGatewayCreate(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*fc.Api)
	public_ip := d.Get("public_ip").(string)

	// Gateway is created by the box registering with FortyCloud
	// Keep trying to find the registration
	gw, err := api.FindGatewayByPublicIP(public_ip, true)
	if err != nil {
		return err
	}
	d.SetId(gw.Id)

	// Merge computed properties from retrieval with modified properties locally
	if err := resourceFcGatewayUpdate(d, meta); err != nil {
		return err
	}

	return resourceFcGatewayRead(d, meta)
}

func resourceFcGatewayRead(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*fc.Api)

	gw, err := api.Gateways.Get(d.Id())
	if err != nil {
		return fmt.Errorf("error retrieving gateway [%s]: %s", d.Id(), err)
	}

	d.Set("fqdn", gw.Fqdn)
	d.Set("overlay_address", gw.OverlayAddress)
	d.Set("vpn_users_subnet", gw.VpnUsersSubnet)
	d.Set("region", gw.Region)
	d.Set("enable", gw.Enable)
	d.Set("release", gw.Release)
	d.Set("allow_ssh_to_everyone", gw.AllowSSHToEveryone)
	d.Set("route_all_traffic_via_gw", gw.RouteAllTrafficViaGW)
	d.Set("private_ip", gw.PrivateIP)
	d.Set("identity_server_name", gw.IdentityServerName)
	d.Set("state", gw.State)
	d.Set("name", gw.Name)
	d.Set("description", gw.Description)
	d.Set("security_group", gw.SecurityGroup)
	d.Set("open_vpn_protocol", gw.OpenVPNProtocol)
	d.Set("gateway_as_dns", gw.GatewayAsDNS)
	d.Set("direct_routes_only", gw.DirectRoutesOnly)
	d.Set("ha_state", gw.HaState)
	return nil
}

func resourceFcGatewayUpdate(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*fc.Api)

	gw, err := api.Gateways.Get(d.Id())
	if err != nil {
		return fmt.Errorf("error retrieving gateway [%s]: %s", d.Id(), err)
	}

	if val, ok := d.GetOk("name"); ok {
		gw.Name = val.(string)
	}

	if val, ok := d.GetOk("description"); ok {
		gw.Description = val.(string)
	}

	if val, ok := d.GetOk("allow_ssh_to_everyone"); ok {
		gw.AllowSSHToEveryone = val.(bool)
	}

	if val, ok := d.GetOk("gateway_as_dns"); ok {
		gw.GatewayAsDNS = val.(bool)
	}

	if val, ok := d.GetOk("identity_server_name"); ok {
		gw.IdentityServerName = val.(string)
	}

	if val, ok := d.GetOk("open_vpn_protocol"); ok {
		gw.OpenVPNProtocol = val.(string)
	}

	if val, ok := d.GetOk("enable"); ok {
		gw.Enable = val.(bool)
	}

	if val, ok := d.GetOk("direct_routes_only"); ok {
		gw.DirectRoutesOnly = val.(bool)
	}

	if _, err := api.Gateways.Update(d.Id(), gw); err != nil {
		return fmt.Errorf("error updating gateway: %s", err)
	}
	return nil
}

func resourceFcGatewayDelete(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*fc.Api)
	id := d.Id()

	if err := api.Gateways.Delete(id); err != nil {
		return fmt.Errorf("error deleting gateway: %s", err)
	}
	return nil
}

func resourceFcGatewayExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	api := meta.(*fc.Api)
	gw, err := api.FindGatewayByPublicIP(d.Get("public_ip").(string), false)
	if err != nil {
		return false, err
	}
	return gw != nil, nil
}
