package eventService

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type EventService interface {
	GetEvents(ctx context.Context) ([]Event, error)
	AddEvents(event []Event, ctx context.Context) (string, error)
}
type eventService struct {
	logger     log.Logger
	repository Repository
}

func NewService(logger log.Logger, repository Repository) EventService {
	return &eventService{
		logger:     logger,
		repository: repository,
	}
}

func (s *eventService) GetEvents(ctx context.Context) ([]Event, error) {
	logger := log.With(s.logger, "method", "GetEvents")
	events, err := s.repository.GetEvents(ctx)

	if err != nil {
		level.Error(logger).Log("err", err)

		//TODO: fix return
		return nil, err

	}

	logger.Log("GetEvents")

	return events, nil
}
func (s *eventService) AddEvents(event []Event, ctx context.Context) (string, error) {
	logger := log.With(s.logger, "method", "AddEvents")
	message, err := s.repository.AddEvents(event, ctx)

	if err != nil {
		level.Error(logger).Log("err", err)

		//TODO: fix return
		return nil, err

	}

	logger.Log("AddEvents")

	return message, nil
}
