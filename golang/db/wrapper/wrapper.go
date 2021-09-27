package wrapper

import (
	"context"
	"database/sql"
	"log"

	"github.com/karlkeefer/pngr/golang/db"
)

type Querier interface {
	db.Querier
	WithTx(context.Context, func(q db.Querier) error) error
}

type Queries struct {
	*db.Queries
	conn *sql.DB
}

func (q *Queries) WithTx(ctx context.Context, fn func(q db.Querier) error) error {
	tx, err := q.conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	err = fn(&Queries{
		Queries: q.Queries.WithTx(tx),
	})

	if err != nil {
		rollBackErr := tx.Rollback()
		if rollBackErr != nil {
			log.Println(rollBackErr)
		}

		return err
	}

	return tx.Commit()
}

func NewQuerier(conn *sql.DB) Querier {
	return &Queries{
		Queries: db.New(conn),
		conn:    conn,
	}
}
