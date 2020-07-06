package eventService

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
)

type GetEventsRequest struct {
}
type GetEventsResponse struct {
	Events []Event `json:"events"`
}

func makeGetEventsEndpoint(s EventService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		events, error := s.GetEvents(ctx)
		return GetEventsResponse{Events: events}, error
	}
}

type AddEventsRequest struct {
	Event []Event `json:"event"`
}
type AddEventsResponse struct {
	Message string `json:"message"`
}

func makeAddEventsEndpoint(s EventService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddEventsRequest)

		message, error := s.AddEvents(req.Event, ctx)
		return AddEventsResponse{Message: message}, error
	}
}

type Endpoints struct {
	GetEvents endpoint.Endpoint
	AddEvents endpoint.Endpoint
}

func MakeEndpoints(s EventService) Endpoints {
	return Endpoints{
		GetEvents: makeGetEventsEndpoint(s),
		AddEvents: makeAddEventsEndpoint(s),
	}
}
func (e GetEventsRequest) Hashcode() string {
	json, _ := json.Marshal(e)
	hasher := md5.New()
	hasher.Write(json)
	code := hex.EncodeToString(hasher.Sum(nil))
	return code
}
func (e AddEventsRequest) Hashcode() string {
	json, _ := json.Marshal(e)
	hasher := md5.New()
	hasher.Write(json)
	code := hex.EncodeToString(hasher.Sum(nil))
	return code
}
