package handlers

import (
	"database/sql"

	"github.com/lib/pq"
)

func isNotFound(err error) bool {
	return err == sql.ErrNoRows
}

func isDupe(err error) bool {
	if err, ok := err.(*pq.Error); ok && err.Code.Class() == "23" {
		// integrity violation
		return true
	}

	return false
}
