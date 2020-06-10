package chatservice

import (
	"context"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPServer(ctx context.Context, endPoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)
	r.Method( /*TODO:choose request*/ ).Path( /*TODO: choose path*/ ).Handler(httptransport.NewServer(
		endpoints.SendMessage,
		decodeSendMessageRequest,
		encodeSendMessageResponse,
	))
	r.Method( /*TODO:choose request*/ ).Path( /*TODO: choose path*/ ).Handler(httptransport.NewServer(
		endpoints.SendImage,
		decodeSendImageRequest,
		encodeSendImageResponse,
	))
	r.Method( /*TODO:choose request*/ ).Path( /*TODO: choose path*/ ).Handler(httptransport.NewServer(
		endpoints.SendLocation,
		decodeSendLocationRequest,
		encodeSendLocationResponse,
	))
	r.Method( /*TODO:choose request*/ ).Path( /*TODO: choose path*/ ).Handler(httptransport.NewServer(
		endpoints.SendAudio,
		decodeSendAudioRequest,
		encodeSendAudioResponse,
	))

	return r
}
