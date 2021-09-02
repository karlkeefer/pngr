package env

import (
	"github.com/karlkeefer/pngr/golang/db"
	"github.com/karlkeefer/pngr/golang/mail"
)

type Env interface {
	DB() db.Querier
	Mailer() *mail.Mailer
}

// default impl
type env struct {
	db   *db.Queries
	mail *mail.Mailer
}

func (e *env) DB() db.Querier {
	return e.db
}

func (e *env) Mailer() *mail.Mailer {
	return e.mail
}

func New() (Env, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	return &env{
		db:   db,
		mail: mail.New(),
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

func (e *mock) Mailer() *mail.Mailer {
	return nil
}
