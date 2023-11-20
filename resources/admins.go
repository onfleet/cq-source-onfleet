package resources

import (
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/onfleet/cq-source-onfleet/client"
	onfleet "github.com/onfleet/gonfleet"
)

func Admins() *schema.Table {
	return &schema.Table{
		Name:        "onfleet_admins",
		Resolver:    fetchAdmins,
		Description: "https://docs.onfleet.com/reference/list-administrators",
		Transform: transformers.TransformWithStruct(&onfleet.Admin{},
			transformers.WithPrimaryKeys("ID", "Organization"),
			transformers.WithResolverTransformer(client.ResolverTransformer),
			transformers.WithTypeTransformer(client.TypeTransformer),
		),
	}
}
