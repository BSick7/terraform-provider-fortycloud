package fortycloud

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"user": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("FORTY_CLOUD_USER", nil),
				Description: "The user name for Forty Cloud.",
			},

			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("FORTY_CLOUD_PASSWORD", nil),
				Description: "The user password for Forty Cloud.",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			
		},

		//ConfigureFunc: providerConfigure,
	}
}

/*
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		User:          d.Get("user").(string),
		Password:      d.Get("password").(string),
	}

	return config.Client()
}
*/