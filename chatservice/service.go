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
	repository Repository
	logger     log.Logger
}

func NewService(rep Repository, logger log.Logger) chatService {
	return &chatService{
		repository: rep,
		logger:     logger,
	}
}
func (s *chatService) CreateUser(firstName string, lastName string, email string, profilePic string, ctx context.Context) (string, error) {
	Logger := log.With(s.logger, "method", CreateUser)
	//TODO: implement
}
func (s *chatService) GetUser(id string, ctx context.Context) (string, string, string, string, error) {
	Logger := log.With(s.logger, "method", GetUser)
	//TODO: implement
}
func (s *chatService) UpdateUser(id string, profilePic string, ctx context.Context) (string, error) {
	Logger := log.With(s.logger, "method", UpdateUser)
	//TODO: implement
}
