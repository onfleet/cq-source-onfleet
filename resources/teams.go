package resources

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/onfleet/cq-source-onfleet/client"
	onfleet "github.com/onfleet/gonfleet"
)

func Teams() *schema.Table {
	return &schema.Table{
		Name:        "onfleet_teams",
		Resolver:    fetchTeams,
		Description: "https://docs.onfleet.com/reference/list-teams",
		Transform: transformers.TransformWithStruct(&onfleet.Team{},
			client.SharedTransformers()...),
		Columns: []schema.Column{
			{
				Name:     "organization_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveOrganizationId,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
