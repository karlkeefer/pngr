package env

import (
	"github.com/karlkeefer/pngr/golang/db"
)

type env struct {
	db *db.Queries
}

// helpful interface for testing
type Env interface {
	DB() *db.Queries
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

func (e *env) DB() *db.Queries {
	return e.db
}
