package env

import (
	"github.com/jmoiron/sqlx"
	"github.com/karlkeefer/pngr/golang/db"
	"github.com/karlkeefer/pngr/golang/models/post"
	"github.com/karlkeefer/pngr/golang/models/user"
)

type Env struct {
	db       *sqlx.DB
	userRepo *user.Repo
	postRepo *post.Repo
}

func New() (*Env, error) {
	db, err := db.New()
	if err != nil {
		return nil, err
	}

	return &Env{
		db:       db,
		userRepo: user.NewRepo(db),
		postRepo: post.NewRepo(db),
	}, nil
}

func (e *Env) UserRepo() *user.Repo {
	return e.userRepo
}

func (e *Env) PostRepo() *post.Repo {
	return e.postRepo
}
