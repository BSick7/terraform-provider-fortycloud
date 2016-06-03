package fortycloud

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"access_key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("FORTYCLOUD_ACCESS_KEY", nil),
				Description: "The access key for FortyCloud.",
			},
			"secret_key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("FORTYCLOUD_SECRET_KEY", nil),
				Description: "The secret key for FortyCloud.",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"fortycloud_gateway":           resourceFcGateway(),
			"fortycloud_ip_sec_connection": resourceFcIPSecConnection(),
			"fortycloud_subnet":            resourceFcSubnet(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		AccessKey: d.Get("access_key").(string),
		SecretKey: d.Get("secret_key").(string),
	}

	log.Println("[INFO] Initializing Forty Cloud service")

	return config.Api()
}
