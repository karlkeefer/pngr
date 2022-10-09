package env

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func Connect() (*pgxpool.Pool, error) {
	conn, err := pgxpool.Connect(context.Background(), buildConnectionString())
	if err != nil {
		return nil, err
	}

	err = conn.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func buildConnectionString() string {
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASSWORD")
	if user == "" || pass == "" {
		log.Fatalln("You must include POSTGRES_USER and POSTGRES_PASSWORD environment variables")
	}
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	dbname := os.Getenv("POSTGRES_DB")
	if host == "" || port == "" || dbname == "" {
		log.Fatalln("You must include POSTGRES_HOST, POSTGRES_PORT, and POSTGRES_DB environment variables")
	}

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, dbname)
}
