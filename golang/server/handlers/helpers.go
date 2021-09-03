package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/lib/pq"
)

// isNotFound helps us match common db error
func isNotFound(err error) bool {
	return err == sql.ErrNoRows
}

// isDupe helps us match common db error
func isDupe(err error) bool {
	if err, ok := err.(*pq.Error); ok && err.Code.Class() == "23" {
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
