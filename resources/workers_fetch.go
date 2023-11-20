package resources

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/onfleet/cq-source-onfleet/client"
)

func fetchWorkers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)

	workers, err := cqClient.OnfleetClient.Workers.List()
	if err != nil {
		return err
	}

	res <- workers

	return nil
}
