package resources

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/onfleet/cq-source-onfleet/client"
	onfleet "github.com/onfleet/gonfleet"
)

func Workers() *schema.Table {
	return &schema.Table{
		Name:        "onfleet_workers",
		Resolver:    fetchWorkers,
		Description: "https://docs.onfleet.com/reference/list-workers",
		Transform: transformers.TransformWithStruct(&onfleet.Worker{},
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
