package main

import (
	"context"
	"flag"
	"log"
	"terraform-provider-aws-profiler/awsprofiler"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "spices.dev/stollenaar/awsprofiler",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), awsprofiler.New, opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
