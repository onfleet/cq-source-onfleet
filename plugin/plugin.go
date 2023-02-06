package plugin

import (
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/onfleet/cq-source-onfleet/client"
	"github.com/onfleet/cq-source-onfleet/resources"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"cloudquery-onfleet",
		Version,
		schema.Tables{
			resources.Admins(),
			resources.Hubs(),
			resources.Workers(),
			resources.Tasks(),
			resources.Teams(),
		},
		client.New,
	)
}
