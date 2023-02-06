package onfleet

import "context"

type AdminsService service

type Admin struct {
	ID               string        `json:"id"`
	TimeCreated      int64         `json:"timeCreated"`
	TimeLastModified int64         `json:"timeLastModified"`
	Organization     string        `json:"organization"`
	Email            string        `json:"email"`
	Type             string        `json:"type"`
	Name             string        `json:"name"`
	IsActive         bool          `json:"isActive"`
	IsReadOnly       bool          `json:"isReadOnly"`
	IsAccountOwner   bool          `json:"isAccountOwner"`
	Phone            string        `json:"phone"`
	Teams            []string      `json:"teams"`
	Metadata         []interface{} `json:"metadata"`
}

type AdminsServiceInterface interface {
	List(ctx context.Context) ([]Admin, error)
}

func (s *AdminsService) List(ctx context.Context) ([]Admin, error) {
	req, err := s.client.NewRequest("GET", "admins", nil)
	if err != nil {
		return nil, err
	}

	var admins []Admin
	err = s.client.Do(ctx, req, &admins)
	if err != nil {
		return nil, err
	}

	return admins, nil
}
