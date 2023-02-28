package onfleet

import "context"

type TeamService service

type Team struct {
	Id                   string   `json:"id"`
	TimeCreated          int64    `json:"timeCreated"`
	TimeLastModified     int64    `json:"timeLastModified"`
	EnableSelfAssignment bool     `json:"enableSelfAssignment"`
	Name                 string   `json:"name"`
	Hub                  *string  `json:"hub"`
	Workers              []string `json:"workers"`
	Managers             []string `json:"managers"`
	Tasks                []string `json:"tasks"`
}

type TeamsServiceInterface interface {
	List(ctx context.Context) ([]Team, error)
}

func (s *TeamService) List(ctx context.Context) ([]Team, error) {
	req, err := s.client.NewRequest("GET", "teams", nil)
	if err != nil {
		return nil, err
	}

	var teams []Team
	err = s.client.Do(ctx, req, &teams)
	if err != nil {
		return nil, err
	}

	return teams, nil
}
