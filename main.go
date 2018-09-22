package main

import (
	"github.com/b-social/terraform-provider-kafkaconnect/kafkaconnect"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return kafkaconnect.Provider()
		},
	})
}
