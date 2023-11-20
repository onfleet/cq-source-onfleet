package resources

import (
	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Name:       "organization_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveOrganizationId,
				PrimaryKey: true,
			},
		},
	}
}
