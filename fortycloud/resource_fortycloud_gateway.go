package fortycloud

import (
	"fmt"
	fc "github.com/BSick7/fortycloud-sdk-go/api"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"time"
)

type MissingGatewayError struct {
	PublicIP string
}

func (err *MissingGatewayError) Error() string {
	return fmt.Sprintf("Could not find gateway with Public IP=%s.", err.PublicIP)
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
			"enable": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"release": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"allow_ssh_to_everyone": &schema.Schema{
				Type:     schema.TypeBool,
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
			"identity_server_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"security_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"open_vpn_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gateway_as_dns": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"direct_routes_only": &schema.Schema{
				Type:     schema.TypeBool,
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
	retryDelay := 15
	retries := 90 / retryDelay
	gw, err := findGatewayForCreate(api, public_ip, retries, retryDelay)
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

	public_ip := d.Get("public_ip").(string)
	gw, err := findGatewayByPublicIP(api, public_ip)
	if err != nil {
		return fmt.Errorf("Error finding gateway: %s", err)
	}

	if gw != nil {
		d.SetId(gw.Id)
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
	}

	return nil
}

func resourceFcGatewayUpdate(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*fc.Api)

	public_ip := d.Get("public_ip").(string)
	gw, err := findGatewayByPublicIP(api, public_ip)
	if err != nil {
		return fmt.Errorf("Error finding gateway: %s", err)
	}

	id := d.Id()
	if val, ok := d.GetOk("name"); ok {
		gw.Name = val.(string)
	} else {
		d.Set("name", gw.Name)
	}

	if val, ok := d.GetOk("description"); ok {
		gw.Description = val.(string)
	} else {
		d.Set("description", gw.Description)
	}

	if val, ok := d.GetOk("allow_ssh_to_everyone"); ok {
		gw.AllowSSHToEveryone = val.(bool)
	} else {
		d.Set("allow_ssh_to_everyone", gw.AllowSSHToEveryone)
	}

	if val, ok := d.GetOk("gateway_as_dns"); ok {
		gw.GatewayAsDNS = val.(bool)
	} else {
		d.Set("gateway_as_dns", gw.GatewayAsDNS)
	}

	if val, ok := d.GetOk("identity_server_name"); ok {
		gw.IdentityServerName = val.(string)
	} else {
		d.Set("identity_server_name", gw.IdentityServerName)
	}

	if val, ok := d.GetOk("open_vpn_protocol"); ok {
		gw.OpenVPNProtocol = val.(string)
	} else {
		d.Set("open_vpn_protocol", gw.OpenVPNProtocol)
	}

	if val, ok := d.GetOk("enable"); ok {
		gw.Enable = val.(bool)
	} else {
		d.Set("enable", gw.Enable)
	}

	if val, ok := d.GetOk("direct_routes_only"); ok {
		gw.DirectRoutesOnly = val.(bool)
	} else {
		d.Set("direct_routes_only", gw.DirectRoutesOnly)
	}

	if _, err := api.Gateways.Update(id, gw); err != nil {
		return err
	}
	return nil
}

func resourceFcGatewayDelete(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*fc.Api)
	id := d.Id()

	if err := api.Gateways.Delete(id); err != nil {
		return fmt.Errorf("Could not delete gateway: %s", err)
	}
	return nil
}

func resourceFcGatewayExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	api := meta.(*fc.Api)
	gw, err := findGatewayByPublicIP(api, d.Get("public_ip").(string))
	if err != nil {
		return false, err
	}
	return gw != nil, nil
}

func findGatewayByPublicIP(api *fc.Api, public_ip string) (*fc.Gateway, error) {
	gws, err := api.Gateways.All()
	if err != nil {
		return nil, fmt.Errorf("Error getting gateways: %s", err)
	}

	for _, cur := range gws {
		if cur.PublicIP == public_ip {
			return &cur, nil
		}
	}
	return nil, nil
}

func findGatewayForCreate(api *fc.Api, public_ip string, retries int, retryDelay int) (gw fc.Gateway, err error) {
	for i := 0; i < retries; i++ {
		gw, err := findGatewayByPublicIP(api, public_ip)
		if err != nil {
			break
		}
		if gw == nil {
			if (i + 1) == retries {
				err = &MissingGatewayError{PublicIP: public_ip}
				break
			}
			log.Printf("Waiting %d seconds to locate gateway (%s)...\n", retryDelay, public_ip)
			time.Sleep(time.Duration(retryDelay) * time.Second)
			continue
		}
	}
	return
}
