package onfleet

import (
	"context"
	"errors"

	"github.com/google/go-querystring/query"
)

type TasksService service

type Task struct {
	Id                      string            `json:"id"`
	TimeCreated             int64             `json:"timeCreated"`
	TimeLastModified        int64             `json:"timeLastModified"`
	Organization            string            `json:"organization"`
	ShortID                 string            `json:"shortId"`
	TrackingURL             string            `json:"trackingURL"`
	Worker                  string            `json:"worker"`
	Merchant                string            `json:"merchant"`
	Executor                string            `json:"executor"`
	Creator                 string            `json:"creator"`
	Dependencies            []interface{}     `json:"dependencies"`
	State                   int               `json:"state"`
	CompleteAfter           int64             `json:"completeAfter"`
	CompleteBefore          int64             `json:"completeBefore"`
	ServiceTime             int64             `json:"serviceTime"`
	PickupTask              bool              `json:"pickupTask"`
	Notes                   string            `json:"notes"`
	CompletionDetails       CompletionDetails `json:"completionDetails"`
	Feedback                []interface{}     `json:"feedback"`
	Metadata                []interface{}     `json:"metadata"`
	Overrides               Overrides         `json:"overrides"`
	Container               Container         `json:"container"`
	Recipients              []Recipients      `json:"recipients"`
	EstimatedCompletionTime int64             `json:"estimatedCompletionTime"`
	EstimatedArrivalTime    int64             `json:"estimatedArrivalTime"`
	Destination             Destination       `json:"destination"`
	DidAutoAssign           bool              `json:"didAutoAssign"`
	TrackingViewed          bool              `json:"trackingViewed"`
	Quantity                int               `json:"quantity"`
}

type Metadata struct {
	Name       string      `json:"name"`
	Type       string      `json:"type"`
	Value      interface{} `json:"value"`
	Visibility []string    `json:"visibility"`
}

type CompletionDetails struct {
	Success       bool          `json:"success"`
	Notes         string        `json:"notes"`
	Events        []interface{} `json:"events"`
	Time          interface{}   `json:"time"`
	FailureReason string        `json:"failureReason"`
}

type Overrides struct {
	RecipientSkipSMSNotifications interface{} `json:"recipientSkipSMSNotifications"`
	RecipientNotes                interface{} `json:"recipientNotes"`
	RecipientName                 interface{} `json:"recipientName"`
}

type Container struct {
	Type   string `json:"type"`
	Worker string `json:"worker"`
	Team   string `json:"team"`
}

type Recipients struct {
	Id                   string        `json:"id"`
	Organization         string        `json:"organization"`
	TimeCreated          int64         `json:"timeCreated"`
	TimeLastModified     int64         `json:"timeLastModified"`
	Name                 string        `json:"name"`
	Phone                string        `json:"phone"`
	Notes                string        `json:"notes"`
	SkipSMSNotifications bool          `json:"skipSMSNotifications"`
	Metadata             []interface{} `json:"metadata"`
}

type Destination struct {
	Id               string        `json:"id"`
	TimeCreated      int64         `json:"timeCreated"`
	TimeLastModified int64         `json:"timeLastModified"`
	Address          Address       `json:"address"`
	Notes            string        `json:"notes"`
	Metadata         []interface{} `json:"metadata"`
}

type TasksListState string

var (
	TasksListStateUnassigned TasksListState = "0"
	TasksListStateAssigned   TasksListState = "1"
	TasksListStateActive     TasksListState = "2"
	TasksListStateCompleted  TasksListState = "3"
)

type TasksListOptions struct {
	From   int64            `url:"from"`
	States []TasksListState `url:"state,omitempty"`
	Worker string           `url:"worker,omitempty"`
}

type TaskCreatePayload struct {
	Destination    *Destination  `json:"destination"`
	Recipients     []*Recipients `json:"recipients"`
	CompleteAfter  int64         `json:"completeAfter"`
	CompleteBefore int64         `json:"completeBefore"`
	Notes          string        `json:"notes"`
	Container      Container     `json:"container"`
	Metadata       []Metadata    `json:"metadata"`
	Quantity       int64         `json:"quantity"`
}

type TaskUpdatePayload struct {
	Destination       *Destination      `json:"destination,omitempty"`
	Recipients        []*Recipients     `json:"recipients,omitempty"`
	CompleteAfter     int64             `json:"completeAfter,omitempty"`
	Notes             string            `json:"notes,omitempty"`
	Container         Container         `json:"container,omitempty"`
	CompletionDetails CompletionDetails `json:"completionDetails"`
	State             int               `json:"state"`
}

type TaskCompletePayload struct {
	CompletionDetails CompletionDetails `json:"completionDetails"`
}

type TasksCreatePayload struct {
	Tasks []TaskCreatePayload `json:"tasks"`
}

type TaskError struct {
	StatusCode int    `json:"statusCode"`
	Error      int    `json:"error"`
	Message    string `json:"message"`
	Cause      string `json:"cause"`
}

type TasksCreateReturn struct {
	Tasks  []Task `json:"tasks"`
	Errors []struct {
		Error TaskError         `json:"error"`
		Task  TaskCreatePayload `json:"task"`
	} `json:"errors"`
}

type TasksServiceInterface interface {
	List(ctx context.Context, opts *TasksListOptions) ([]Task, error)
	Create(ctx context.Context, payload *TasksCreatePayload) ([]Task, error)
	Update(ctx context.Context, taskId string, payload *TaskUpdatePayload) (*Task, error)
	Complete(ctx context.Context, taskId string, payload *TaskCompletePayload) (*Task, error)
	Get(ctx context.Context, taskId string) (*Task, error)
	Delete(ctx context.Context, taskId string) (*Task, error)
}

// List all tasks
// https://docs.onfleet.com/reference#list-tasks
func (s *TasksService) List(ctx context.Context, opts *TasksListOptions) ([]Task, error) {
	var tasks []Task
	err := s.getMany(ctx, "tasks", opts, &tasks)
	return tasks, err
}

func (s *TasksService) getMany(ctx context.Context, path string, opts interface{}, v interface{}) error {
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

// Create tasks
// https://docs.onfleet.com/reference#create-tasks-in-batch
func (s *TasksService) Create(ctx context.Context, payload *TasksCreatePayload) ([]Task, error) {
	var res TasksCreateReturn
	req, err := s.client.NewRequest("POST", "tasks/batch", payload)
	if err != nil {
		return nil, err
	}

	err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}
	if len(res.Errors) != 0 {
		return nil, errors.New(res.Errors[0].Error.Message)
	}

	return res.Tasks, nil
}

// Update task
// https://docs.onfleet.com/reference#update-task
func (s *TasksService) Update(ctx context.Context, taskId string, payload *TaskUpdatePayload) (*Task, error) {
	var res Task
	req, err := s.client.NewRequest("PUT", "tasks/"+taskId, payload)
	if err != nil {
		return nil, err
	}

	err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *TasksService) Complete(ctx context.Context, taskId string, payload *TaskCompletePayload) (*Task, error) {
	var res Task
	req, err := s.client.NewRequest("POST", "tasks/"+taskId+"/complete", payload)
	if err != nil {
		return nil, err
	}

	err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Get task
// https://docs.onfleet.com/reference#get-single-task
func (s *TasksService) Get(ctx context.Context, taskId string) (*Task, error) {
	var res Task
	req, err := s.client.NewRequest("GET", "tasks/"+taskId, nil)
	if err != nil {
		return nil, err
	}

	err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Delete task
// https://docs.onfleet.com/reference/delete-task
func (s *TasksService) Delete(ctx context.Context, taskId string) (*Task, error) {
	var res Task
	req, err := s.client.NewRequest("DELETE", "tasks/"+taskId, nil)
	if err != nil {
		return nil, err
	}

	err = s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
