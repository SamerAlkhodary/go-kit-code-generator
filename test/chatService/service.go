package chatService

import (
	"context"

	"github.com/go-kit/kit/log"
)

type ChatService interface {
	CreateUser(firstName string, lastName string, email string, profilePic string, ctx context.Context) (string, error)
	GetUser(id string, ctx context.Context) (string, string, string, string, error)
	UpdateUser(id string, profilePic string, ctx context.Context) (string, error)
}
type chatService struct {
	logger log.Logger
}

func NewService(logger log.Logger) ChatService {
	return &chatService{
		logger: logger,
	}
}
func (s *chatService) CreateUser(firstName string, lastName string, email string, profilePic string, ctx context.Context) (string, error) {
	Logger := log.With(s.logger, "method", "CreateUser")
	Logger.Log(firstName)
	//TODO: implement
	return "", nil
}
func (s *chatService) GetUser(id string, ctx context.Context) (string, string, string, string, error) {
	Logger := log.With(s.logger, "method", "GetUser")
	Logger.Log(id)
	//TODO: implement
	return "", "", "", "", nil
}
func (s *chatService) UpdateUser(id string, profilePic string, ctx context.Context) (string, error) {
	Logger := log.With(s.logger, "method", "UpdateUser")
	Logger.Log(id)
	//TODO: implement
	return "", nil
}
