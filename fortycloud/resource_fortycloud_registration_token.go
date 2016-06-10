package fortycloud

import (
	"fmt"
	fc "github.com/BSick7/fortycloud-sdk-go/api"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFcRegistrationToken() *schema.Resource {
	return &schema.Resource{
		Create: resourceFcRegistrationTokenCreate,
		Read:   resourceFcRegistrationTokenRead,
		Delete: resourceFcRegistrationTokenDelete,

		Schema: map[string]*schema.Schema{
			// platform types
			//     AwsBilledAMI - AWS AMI with installed Fortycloud gateway, paid via the AWS marketplace
			//     AwsByolAMI – AWS AMI with installed Fortycloud gateway, license purchased via Fortycloud
			//     Rackspace
			//     IBM
			//     Azure
			//     CenturyLink
			//     OtherPlatform (for any cloud platform, not pre-installed with Fortycloud)
			"platform": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			// license types
			//     LargeGW
			"license": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"token": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceFcRegistrationTokenCreate(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*fc.Api)

	platform := d.Get("platform").(string)
	license := d.Get("license").(string)

	token, err := api.Gateways.GetRegistrationToken(platform, license)
	if err != nil {
		return fmt.Errorf("error generating registration token: %s", err)
	}

	d.Set("token", token)
	d.SetId(token)
	return nil
}

func resourceFcRegistrationTokenRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceFcRegistrationTokenDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
