package env

import (
	"github.com/QuinnMain/infograph/golang/mail"
	"go.mongodb.org/mongo-driver/mongo"
)

type Env interface {
	MDB() *mongo.Database
	Mailer() *mail.Mailer
}

// default impl
type env struct {
	mdb  *mongo.Database
	mail *mail.Mailer
}

// accesors (зачем, я не заметил вроде чтобы их вызывали)

func (e *env) MDB() *mongo.Database {
	return e.mdb
}
func (e *env) Mailer() *mail.Mailer {
	return e.mail
}

func NewEnv() (Env, error) {
	mongodb, err := ConnectToMongo()
	if err != nil {
		return nil, err
	}

	return &env{ // как это работает?
		mdb:  mongodb,
		mail: mail.New(),
	}, nil
}

// Mock impl
func Mock(mdb *mongo.Database) Env {
	return &mock{
		mdb: mdb,
	}
}

type mock struct {
	mdb *mongo.Database
}

func (e *mock) MDB() *mongo.Database {
	return e.mdb
}

func (e *mock) Mailer() *mail.Mailer {
	return nil
}
