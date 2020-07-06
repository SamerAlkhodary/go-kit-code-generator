package eventService

import (
	"context"
	"database/sql"
	"github.com/go-kit/kit/log"
)

type Repository interface {
	GetEvents(ctx context.Context) ([]Event, error)
	AddEvents(event []Event, ctx context.Context) (string, error)
}
type repository struct {
	db     *sql.DB
	logger log.Logger
}

func MakeNewRepository(db *sql.DB, logger log.Logger) Repository {

	return &repository{
		db:     db,
		logger: log.With(logger, "repository", "sql"),
	}

}

func (repo *repository) GetEvents(ctx context.Context) ([]Event, error) {
}

func (repo *repository) AddEvents(event []Event, ctx context.Context) (string, error) {
}
