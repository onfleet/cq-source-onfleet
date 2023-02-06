# go-onfleet

Go client for accessing [Onfleet API](https://docs.onfleet.com/reference)

## Usage

```bash
go get github.com/onfleet/cq-source-onfleet/internal/go-onfleet
```

```go
import (
    "context"
    "fmt"
    "github.com/keplr-team/go-onfleet/onfleet"
)

client := onfleet.NewClient("API_KEY")
workers, err := client.Workers.List(context.Background(), nil)
if err != nil {
    // do stuff
}
fmt.Printf("%+v\n", workers)
```

### What's inside

- [] Admins API
  - [x] List admins 

- [ ] Workers API
  - [ ] Create worker
  - [x] List workers
  - [x] Get workers by location
  - [x] Get single worker
  - [ ] Update worker
  - [ ] Delete worker

- [ ] Hubs API
  - [ ] Create hub
  - [x] List hubs
  - [ ] Update hub

- [ ] Teams API
  - [ ] Create team
  - [x] List team
  - [ ] Update team
  - [ ] Get single team
  - [ ] Delete team

- [ ] Tasks API
  - [x] Create tasks
  - [x] List tasks
  - [x] Update task
  - [x] Get single task
  - [ ] Delete task

### What's next

- [ ] unit testing
- [ ] handle errors
- [ ] handle rate limit
- [ ] add more endpoints
