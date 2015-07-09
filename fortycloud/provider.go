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
            "username": &schema.Schema{
                Type:        schema.TypeString,
                Required:    true,
                DefaultFunc: schema.EnvDefaultFunc("FORTYCLOUD_USERNAME", nil),
                Description: "The username for Forty Cloud.",
            },
            "password": &schema.Schema{
                Type:        schema.TypeString,
                Required:    true,
                DefaultFunc: schema.EnvDefaultFunc("FORTYCLOUD_PASSWORD", nil),
                Description: "The password for Forty Cloud.",
            },
            "tenantName": &schema.Schema{
                Type:        schema.TypeString,
                Required:    true,
                DefaultFunc: schema.EnvDefaultFunc("FORTYCLOUD_TENANTNAME", nil),
                Description: "The tenant for Forty Cloud.",
            },
            "formUsername": &schema.Schema{
                Type:        schema.TypeString,
                Required:    true,
                DefaultFunc: schema.EnvDefaultFunc("FORTYCLOUD_FORMUSERNAME", nil),
                Description: "The username for web-based authentication",
            },
            "formPassword": &schema.Schema{
                Type:        schema.TypeString,
                Required:    true,
                DefaultFunc: schema.EnvDefaultFunc("FORTYCLOUD_FORMPASSWORD", nil),
                Description: "The password for web-based authentication",
            },
        },

        ResourcesMap: map[string]*schema.Resource{
            "fortycloud_connection":                resourceFcConnection(),
            "fortycloud_node":                      resourceFcNode(),
            "fortycloud_subnet":                    resourceFcSubnet(),
        },

        ConfigureFunc: providerConfigure,
    }
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
    config := Config{
        Username:      d.Get("username").(string),
        Password:      d.Get("password").(string),
        TenantName:    d.Get("tenantName").(string),
        FormUsername:  d.Get("formUsername").(string),
        FormPassword:  d.Get("formPassword").(string),
    }
    
    log.Println("[INFO] Initializing Forty Cloud service")

    return config.Api()
}