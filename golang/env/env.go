package env

import (
	"github.com/karlkeefer/pngr/golang/db"
)

type Env interface {
	DB() db.Querier
}

// default impl
type env struct {
	db *db.Queries
}

func (e *env) DB() db.Querier {
	return e.db
}

func New() (Env, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	return &env{
		db: db,
	}, nil
}

// Mock impl
func Mock(db db.Querier) Env {
	return &mock{
		db: db,
	}
}

type mock struct {
	db db.Querier
}

func (e *mock) DB() db.Querier {
	return e.db
}
