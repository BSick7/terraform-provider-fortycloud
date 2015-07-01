package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/mdl/terraform-provider-fortycloud/fortycloud"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: fortycloud.Provider,
	})
}