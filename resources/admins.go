package resources

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/onfleet/cq-source-onfleet/internal/go-onfleet/onfleet"
)

func Admins() *schema.Table {
	return &schema.Table{
		Name:        "onfleet_admins",
		Resolver:    fetchAdmins,
		Description: "https://docs.onfleet.com/reference/list-administrators",
		Transform:   transformers.TransformWithStruct(&onfleet.Admin{}, transformers.WithPrimaryKeys("ID")),
	}
}
