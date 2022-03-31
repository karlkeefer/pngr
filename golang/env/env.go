package env

import (
	"github.com/QuinnMain/infograph/golang/mail"
	"go.mongodb.org/mongo-driver/mongo"
)

var collectionNames = [4]string{"commodityTables", "commodityCharts", "divisions", "users"}

type Env interface {
	MDB() *mongo.Database
	Mailer() *mail.Mailer
	Collection(string) *mongo.Collection
}

// default impl
type env struct {
	mdb         *mongo.Database
	mail        *mail.Mailer
	collections map[string]*mongo.Collection
}

func (e *env) MDB() *mongo.Database {
	return e.mdb
}
func (e *env) Mailer() *mail.Mailer {
	return e.mail
}
func (e *env) Collection(colname string) *mongo.Collection {
	return e.collections[colname]
}

func NewEnv() (Env, error) {
	mongodb, err := ConnectToMongo()
	if err != nil {
		return nil, err
	}
	var collections map[string]*mongo.Collection
	for _, colname := range collectionNames {
		collections[colname] = mongodb.Collection(colname)
	}

	return &env{ // как это работает?
		mdb:         mongodb,
		mail:        mail.New(),
		collections: collections,
	}, nil
}

// Mock impl
func Mock(mdb *mongo.Database) Env {
	var collections map[string]*mongo.Collection
	for _, colname := range collectionNames {
		collections[colname] = mdb.Collection(colname)
	}
	return &mock{
		mdb:         mdb,
		mail:        nil,
		collections: collections,
	}
}

type mock struct {
	mdb         *mongo.Database
	mail        *mail.Mailer
	collections map[string]*mongo.Collection
}

func (e *mock) MDB() *mongo.Database {
	return e.mdb
}

func (e *mock) Mailer() *mail.Mailer {
	return nil
}

func (e *mock) Collection(colname string) *mongo.Collection {
	return e.collections[colname]
}
