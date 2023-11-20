package resources

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/onfleet/cq-source-onfleet/client"
	onfleet "github.com/onfleet/gonfleet"
)

func fetchTasks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)

	params := onfleet.TaskListQueryParams{
		From: cqClient.ListTasksFrom.UnixMilli(),
	}

	for {
		response, err := cqClient.OnfleetClient.Tasks.List(params)

		if err != nil {
			return err
		}

		res <- response.Tasks

		if response.LastId == "" {
			break
		}

		params.LastId = response.LastId
	}

	return nil
}
