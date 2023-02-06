package onfleet

import (
	"context"

	"github.com/google/go-querystring/query"
)

type WorkersService service

type UserData struct {
	Platform          string  `json:"platform"`
	DeviceDescription string  `json:"deviceDescription"`
	BatteryLevel      float64 `json:"batteryLevel"`
	AppVersion        string  `json:"appVersion"`
}

type Vehicle struct {
	Id               string  `json:"id"`
	Type             string  `json:"type"`
	Description      *string `json:"description"`
	LicensePlate     *string `json:"licensePlate"`
	Color            *string `json:"color"`
	TimeLastModified int64   `json:"timeLastModified"`
}

type Worker struct {
	Id               string        `json:"id"`
	TimeCreated      int64         `json:"timeCreated"`
	TimeLastModified int64         `json:"timeLastModified"`
	Organization     string        `json:"organization"`
	Name             string        `json:"name"`
	DisplayName      string        `json:"displayName"`
	Phone            string        `json:"phone"`
	ActiveTask       *string       `json:"activeTask"`
	Tasks            []string      `json:"tasks"`
	OnDuty           bool          `json:"onDuty"`
	IsResponding     bool          `json:"isResponding"`
	TimeLastSeen     *int64        `json:"timeLastSeen"`
	Capacity         int           `json:"capacity"`
	AccountStatus    string        `json:"accountStatus"`
	ImageUrl         *string       `json:"imageUrl"`
	Location         []float64     `json:"location"`
	DelayTime        *float64      `json:"delayTime"`
	Timezone         *string       `json:"timezone"`
	Teams            []string      `json:"teams"`
	UserData         UserData      `json:"userData"`
	Vehicle          *Vehicle      `json:"vehicle"`
	Metadata         []interface{} `json:"metadata"`

	HasRecentlyUsedSpoofedLocations bool `json:"hasRecentlyUsedSpoofedLocations"`
}

type WorkersListState string

var (
	WorkersListStateOffDuty WorkersListState = "0"
	WorkersListStateIdle    WorkersListState = "1"
	WorkersListStateActive  WorkersListState = "2"
)

type WorkersListOptions struct {
	Teams  []string           `url:"teams,omitempty"`
	States []WorkersListState `url:"states,omitempty"`
	Phones []string           `url:"phones,omitempty"`
}

type WorkersLocationOptions struct {
	Longitude float64  `url:"longitude,omitempty"`
	Latitude  float64  `url:"latitude,omitempty"`
	Radius    *float64 `url:"radius,omitempty"`
}

type WorkersServiceInterface interface {
	List(ctx context.Context, opts *WorkersListOptions) ([]Worker, error)
	Location(ctx context.Context, opts *WorkersLocationOptions) ([]Worker, error)
	Get(ctx context.Context, workerId string) (*Worker, error)
}

// List list all workers
// https://docs.onfleet.com/reference#list-workers
func (s *WorkersService) List(ctx context.Context, opts *WorkersListOptions) ([]Worker, error) {
	var workers []Worker
	err := s.getMany(ctx, "workers", opts, &workers)
	return workers, err
}

// Location get workers by location
// https://docs.onfleet.com/reference#get-workers-by-location
func (s *WorkersService) Location(ctx context.Context, opts *WorkersLocationOptions) ([]Worker, error) {
	var res struct {
		Workers []Worker `json:"workers"`
	}

	err := s.getMany(ctx, "workers/location", opts, &res)
	return nil, err
}

// Get worker
// https://docs.onfleet.com/reference#get-single-worker
func (s *WorkersService) Get(ctx context.Context, workerId string) (*Worker, error) {
	var res Worker
	req, err := s.client.NewRequest("GET", "workers/"+workerId, nil)
	if err != nil {
		return nil, err
	}

	err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *WorkersService) getMany(ctx context.Context, path string, opts interface{}, v interface{}) error {
	if opts != nil {
		v, err := query.Values(opts)
		if err != nil {
			return err
		}
		path += "?" + v.Encode()
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return err
	}

	err = s.client.Do(ctx, req, v)
	if err != nil {
		return err
	}

	return nil
}
