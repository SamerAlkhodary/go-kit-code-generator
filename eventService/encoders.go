package eventService

import (
	"context"
	"encoding/json"
	"net/http"
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
func decodeGetEventsRequest(ctx context.Context, r *http.Request) (interface{}, error) {

	var request GetEventsRequest
	request = GetEventsRequest{}
	return request, nil
}

func decodeAddEventsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request AddEventsRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}
	return request, nil
}
