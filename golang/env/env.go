package env

import (
	"github.com/jmoiron/sqlx"
	"github.com/karlkeefer/pngr/golang/repos/post"
	"github.com/karlkeefer/pngr/golang/repos/user"
)

type env struct {
	db       *sqlx.DB
	userRepo user.Repo
	postRepo post.Repo
}

// helpful interface for testing
type Env interface {
	UserRepo() user.Repo
	PostRepo() post.Repo
}

func New() (Env, error) {
	db, err := NewDB()
	if err != nil {
		return nil, err
	}

	return &env{
		db:       db,
		userRepo: user.NewRepo(db),
		postRepo: post.NewRepo(db),
	}, nil
}

func (e *env) UserRepo() user.Repo {
	return e.userRepo
}

func (e *env) PostRepo() post.Repo {
	return e.postRepo
}

// mock
func Mock(db *sqlx.DB, ur user.Repo, pr post.Repo) Env {
	return &env{
		db:       db,
		userRepo: ur,
		postRepo: pr,
	}
}
