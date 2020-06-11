package chatService

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPServer(ctx context.Context, endPoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)
	r.Methods("POST").Path("/user").Handler(httptransport.NewServer(
		endpoints.CreateUser,
		decodeCreateUserRequest,
		encodeCreateUserResponse,
	))
	r.Methods("GET").Path("/user/{id}").Handler(httptransport.NewServer(
		endpoints.GetUser,
		decodeGetUserRequest,
		encodeGetUserResponse,
	))
	r.Methods("PUT").Path("/user/update").Handler(httptransport.NewServer(
		endpoints.UpdateUser,
		decodeUpdateUserRequest,
		encodeUpdateUserResponse,
	))

	return r
}
func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
