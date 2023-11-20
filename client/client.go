package client

import (
	"fmt"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	onfleetclient "github.com/onfleet/gonfleet/client"
	"github.com/rs/zerolog"
)

type Client struct {
	OnfleetClient *onfleetclient.API

	Spec Spec

	// Goes into the `organization_id` column, to help distinguish orgs in case of multiple syncs from different orgs.
	OrganizationId string
	ListTasksFrom  time.Time

	Logger zerolog.Logger
}

func (c *Client) ID() string {
	return "org:" + c.OrganizationId
}

func New(logger zerolog.Logger, pluginSpec Spec) (schema.ClientMeta, error) {
	if pluginSpec.ApiKey == "" {
		return nil, fmt.Errorf("API key not provided. Please provided it in the 'api_key' configuration option - See (https://github.com/onfleet/cq-source-onfleet/blob/master/README.md)")
	}

	onfleetClient, err := onfleetclient.New(pluginSpec.ApiKey, nil)

	if err != nil {
		return nil, fmt.Errorf("failed to create onfleet client: %s", err)
	}

	orgId, err := getOrganizationId(onfleetClient)
	if err != nil {
		return nil, fmt.Errorf("failed to get org details: %s", err)
	}

	var listTasksFrom time.Time
	if pluginSpec.ListTasksFrom == "" {
		listTasksFrom = time.Now().AddDate(
			/* years: */ 0,
			/* months: */ -3,
			/* days: */ 0,
		)
	} else {
		listTasksFrom, err = time.Parse(time.RFC3339, pluginSpec.ListTasksFrom)
		if err != nil {
			return nil, fmt.Errorf("failed to parse 'list_tasks_from' (format should be RFC3339 '2002-10-02T10:00:00Z'): %s", err)
		}
	}

	listTasksFrom = listTasksFrom.UTC()

	return &Client{
		OnfleetClient:  onfleetClient,
		Spec:           pluginSpec,
		OrganizationId: orgId,
		ListTasksFrom:  listTasksFrom,
		Logger:         logger,
	}, nil
}

func getOrganizationId(onfleetClient *onfleetclient.API) (string, error) {

	org, err := onfleetClient.Organizations.Get()
	if err != nil {
		return "", err
	}

	if org.ID == "" {
		return "", fmt.Errorf("organization id is empty")
	}

	return org.ID, nil
}
