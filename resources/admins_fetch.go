package resources

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/onfleet/cq-source-onfleet/client"
)

func fetchAdmins(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)

	result, err := cqClient.OnfleetClient.Administrators.List()
	if err != nil {
		return err
	}

	res <- result

	return nil
}
