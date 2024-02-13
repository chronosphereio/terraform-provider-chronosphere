package main

import (
	"flag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere"
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{
		ProviderFunc: chronosphere.Provider,
		ProviderAddr: chronosphere.LocalName,
	}

	if debugMode {
		opts.Debug = true
	}

	plugin.Serve(opts)
}
