package onfleet

import "context"

type OrganizationService service

type Organization struct {
	Id               string   `json:"id"`
	TimeCreated      int64    `json:"timeCreated"`
	TimeLastModified int64    `json:"timeLastModified"`
	Name             string   `json:"name"`
	Email            string   `json:"email"`
	Image            string   `json:"image"`
	Timezone         string   `json:"timezone"`
	Country          string   `json:"country"`
	Delegatees       []string `json:"delegatees"`
}

type OrganizationServiceInterface interface {
	Get(ctx context.Context) (*Organization, error)
}

func (s *OrganizationService) Get(ctx context.Context) (*Organization, error) {
	req, err := s.client.NewRequest("GET", "organization", nil)
	if err != nil {
		return nil, err
	}

	var org Organization
	err = s.client.Do(ctx, req, &org)
	if err != nil {
		return nil, err
	}

	return &org, nil
}
