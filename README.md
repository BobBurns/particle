# particle
--
    import "github.com/mckee/particle"


## Usage

```go
const URL string = "https://api.particle.io/v1/devices/events/"
```

#### func  Subscribe

```go
func Subscribe(eventPrefix string, token string) <-chan Event
```
subscribes to a particle.io event stream and returns a channel to receive them

#### type Event

```go
type Event struct {
	Name string
	Data struct {
		Data      string `json:"data"`
		TTL       string `json:"ttl"`
		Timestamp string `json:"published_at"`
		CoreID    string `json:"coreid"`
	}
}
```

## This is a Fork of github.com/mckee/particle

Add support to publish event. For example see https://github.com/BobBurns/motorworld

#### func Publish

```go
func (e *Event) Publish (token string) (*Result, error)
```

Slightly changed the Event type

```go
type Data struct {
	Data      string    `json:"data"`
	TTL       uint32    `json:"ttl"`
	Timestamp time.Time `json:"published_at"`
	CoreID    string    `json:"coreid"`
	Private	  bool
}

type Event struct {
	Name string
	Data Data

}
```
