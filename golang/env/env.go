package env

import (
	"github.com/QuinnMain/infograph/golang/db/postgres/db/wrapper"
	"github.com/QuinnMain/infograph/golang/mail"
	"go.mongodb.org/mongo-driver/mongo"
)

type Env interface {
	DB() wrapper.Querier
	MDB() *mongo.Database
	Mailer() *mail.Mailer
}

// default impl
type env struct {
	db   wrapper.Querier
	mdb  *mongo.Database
	mail *mail.Mailer
}

// accesors (зачем, я не заметил вроде чтобы их вызывали)
func (e *env) DB() wrapper.Querier {
	return e.db
}
func (e *env) MDB() *mongo.Database {
	return e.mdb
}
func (e *env) Mailer() *mail.Mailer {
	return e.mail
}

func NewEnv() (Env, error) {
	db, err := ConnectToPostgres()
	if err != nil {
		return nil, err
	}

	mongodb, err := ConnectToMongo()
	if err != nil {
		return nil, err
	}

	return &env{ // как это работает?
		db:   wrapper.NewQuerier(db),
		mdb:  mongodb,
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

func (e *mock) MDB() *mongo.Database {
	return nil
}

func (e *mock) Mailer() *mail.Mailer {
	return nil
}
