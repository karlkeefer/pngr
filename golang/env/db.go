package env

import (
	"context"
	"os"
	"time"

	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectToMongo() (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // не знаю что делает WithTimeout
	defer cancel()
	// connection := options.Client().ApplyURI(os.Getenv("MONGO_HOST"))
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_HOST")))
	if err != nil {
		return nil, err
	}
	// пингуем на всякий
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	return client.Database(os.Getenv("MONGO_DB")), err
}
