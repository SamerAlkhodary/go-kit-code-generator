package chatService

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type ChatService interface {
	CreateUser(user User, ctx context.Context) (string, error)
	GetUser(id string, ctx context.Context) (User, error)
	UpdateUser(id string, profilePic string, ctx context.Context) (string, error)
}
type chatService struct {
	logger     log.Logger
	repository Repository
}

func NewService(logger log.Logger, repository Repository) ChatService {
	return &chatService{
		logger:     logger,
		repository: repository,
	}
}
func (s *chatService) CreateUser(user User, ctx context.Context) (string, error) {
	logger := log.With(s.logger, "method", "CreateUser")
	id, err := s.repository.CreateUser(user, ctx)

	if err != nil {
		level.Error(logger).Log("err", err)

		//TODO: fix return
		return nil, err

	}

	logger.Log("CreateUser")

	return id, nil
}
func (s *chatService) GetUser(id string, ctx context.Context) (User, error) {
	logger := log.With(s.logger, "method", "GetUser")
	user, err := s.repository.GetUser(id, ctx)

	if err != nil {
		level.Error(logger).Log("err", err)

		//TODO: fix return
		return nil, err

	}

	logger.Log("GetUser")

	return user, nil
}
func (s *chatService) UpdateUser(id string, profilePic string, ctx context.Context) (string, error) {
	logger := log.With(s.logger, "method", "UpdateUser")
	message, err := s.repository.UpdateUser(id, profilePic, ctx)

	if err != nil {
		level.Error(logger).Log("err", err)

		//TODO: fix return
		return nil, err

	}

	logger.Log("UpdateUser")

	return message, nil
}
