package db

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
	_ "github.com/lib/pq"
)

func New() (database *sqlx.DB, err error) {
	str := buildConnectionString()

	database, err = sqlx.Open("postgres", str)
	if err != nil {
		return
	}

	err = database.Ping()
	if err != nil {
		return
	}

	// this lets us use the json tags for finding postgres column names
	database.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)

	return
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
