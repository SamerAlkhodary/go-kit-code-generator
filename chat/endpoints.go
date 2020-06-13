package chatService

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type CreateUserRequest struct {
	User User `json:"user"`
}
type CreateUserResponse struct {
	Id string `json:"id"`
}

func makeCreateUserEndpoint(s ChatService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		id, error := s.CreateUser(req.User, ctx)
		return CreateUserResponse{Id: id}, error
	}
}

type GetUserRequest struct {
	Id string `json:"id"`
}
type GetUserResponse struct {
	User User `json:"user"`
}

func makeGetUserEndpoint(s ChatService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)
		user, error := s.GetUser(req.Id, ctx)
		return GetUserResponse{User: user}, error
	}
}

type UpdateUserRequest struct {
	Id         string `json:"id"`
	ProfilePic string `json:"profilePic"`
}
type UpdateUserResponse struct {
	Message string `json:"message"`
}

func makeUpdateUserEndpoint(s ChatService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateUserRequest)
		message, error := s.UpdateUser(req.Id, req.ProfilePic, ctx)
		return UpdateUserResponse{Message: message}, error
	}
}

type Endpoints struct {
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
	UpdateUser endpoint.Endpoint
}

func MakeEndpoints(s ChatService) Endpoints {
	return Endpoints{
		CreateUser: makeCreateUserEndpoint(s),
		GetUser:    makeGetUserEndpoint(s),
		UpdateUser: makeUpdateUserEndpoint(s),
	}
}
