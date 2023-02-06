package onfleet

import (
	"context"
)

type HubsService service

type Address struct {
	Number     string `json:"number"`
	Street     string `json:"street"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	PostalCode string `json:"postalCode"`
	Apartment  string `json:"apartment"`
	Unparsed   string `json:"unparsed"`
}

type Hub struct {
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Location []float64 `json:"location"`
	Address  Address   `json:"address"`
	Teams    []string  `json:"teams,omitempty"`
}

type HubsServiceInterface interface {
	List(ctx context.Context) ([]Hub, error)
}

func (s *HubsService) List(ctx context.Context) ([]Hub, error) {
	req, err := s.client.NewRequest("GET", "hubs", nil)
	if err != nil {
		return nil, err
	}

	var hubs []Hub
	err = s.client.Do(ctx, req, &hubs)
	if err != nil {
		return nil, err
	}

	return hubs, nil
}
