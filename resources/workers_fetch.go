package resources

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/onfleet/cq-source-onfleet/client"
	"github.com/onfleet/cq-source-onfleet/internal/go-onfleet/onfleet"
)

func fetchWorkers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)

	workers, err := cqClient.OnfleetClient.Workers().List(ctx, &onfleet.WorkersListOptions{})
	if err != nil {
		return err
	}

	res <- workers

	return nil
}
