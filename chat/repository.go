package chatService

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-kit/kit/log"
)

var RepoError = errros.New("Unable to handle repo request")

type Repository interface {
	CreateUser(user User, ctx context.Context) (string, error)
	GetUser(id string, ctx context.Context) (User, error)
	UpdateUser(id string, profilePic string, ctx context.Context) (string, error)
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

func (repo *repository) CreateUser(user User, ctx context.Context) (string, error) {
}

func (repo *repository) GetUser(id string, ctx context.Context) (User, error) {
}

func (repo *repository) UpdateUser(id string, profilePic string, ctx context.Context) (string, error) {
}
