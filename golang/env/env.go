package env

import (
	"github.com/jmoiron/sqlx"

	"github.com/karlkeefer/pngr/golang/db"
	"github.com/karlkeefer/pngr/golang/models/user"
)

type Env struct {
	db       *sqlx.DB
	userRepo *user.Repo
}

func (e *Env) UserRepo() *user.Repo {
	return e.userRepo
}

func New() (*Env, error) {
	db, err := db.New()
	if err != nil {
		return nil, err
	}

	return &Env{
		db:       db,
		userRepo: user.NewRepo(db),
	}, nil
}
