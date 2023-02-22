package handlers

import (
	"net/http"
	"strconv"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/julienschmidt/httprouter"
)

// isNotFound helps us match common db error
func isNotFound(err error) bool {
	return err == pgx.ErrNoRows
}

// isDupe helps us match common db error
func isDupe(err error) bool {
	if err, ok := err.(*pgconn.PgError); ok && err.Code == "23505" {
		// integrity violation
		return true
	}

	return false
}

// helpers for easily parsing params
func getInt64(name string, r *http.Request) (out int64, err error) {
	params := httprouter.ParamsFromContext(r.Context())
	arg := params.ByName(name)
	out, err = strconv.ParseInt(arg, 10, 64)
	return
}

func getID(r *http.Request) (out int64, err error) {
	return getInt64("id", r)
}

func getString(name string, r *http.Request) (param string) {
	params := httprouter.ParamsFromContext(r.Context())
	return params.ByName(name)
}
