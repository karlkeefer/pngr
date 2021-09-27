package env

import (
	"github.com/karlkeefer/pngr/golang/db/wrapper"
	"github.com/karlkeefer/pngr/golang/mail"
)

type Env interface {
	DB() wrapper.Querier
	Mailer() *mail.Mailer
}

// default impl
type env struct {
	db   wrapper.Querier
	mail *mail.Mailer
}

func (e *env) DB() wrapper.Querier {
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
		db:   wrapper.NewQuerier(db),
		mail: mail.New(),
	}, nil
}

// Mock impl
func Mock(db wrapper.Querier) Env {
	return &mock{
		db: db,
	}
}

type mock struct {
	db wrapper.Querier
}

func (e *mock) DB() wrapper.Querier {
	return e.db
}

func (e *mock) Mailer() *mail.Mailer {
	return nil
}
