package main

import (
	"github.com/cloudquery/plugin-sdk/serve"
	"github.com/onfleet/cq-source-onfleet/plugin"
)

func main() {
	serve.Source(plugin.Plugin())
}
