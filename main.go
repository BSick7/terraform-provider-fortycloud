package main

import (
	"github.com/bsick7/terraform-provider-fortycloud/fortycloud"
	"github.com/hashicorp/terraform/plugin"
)

var Version string

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: fortycloud.Provider,
	})
}
