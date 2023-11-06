package plugin

import (
	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

const (
	Name    = "onfleet"
	Kind    = "source"
	Team    = "onfleet"
	Version = "development"
)

func Plugin() *plugin.Plugin {
	return plugin.NewPlugin(
		Name,
		Version,
		Configure,
		plugin.WithKind(Kind),
		plugin.WithTeam(Team),
	)
}
