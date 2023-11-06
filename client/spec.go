package client

type Spec struct {
	ApiKey      string `json:"api_key"`
	Concurrency int    `json:"concurrency"`
	// By default, only the last 3 months of tasks will be synced.
	// You can specify a different 'from' time here (In RFC3339 '2019-10-12T07:20:50.52Z' format).
	ListTasksFrom string `json:"list_tasks_from"`
}
