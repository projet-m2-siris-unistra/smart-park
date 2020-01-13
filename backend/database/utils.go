package database

import (
	"database/sql"
	"errors"
	"fmt"
)


func setOrderBy(id string) (string) {
	result := " ORDER BY " + id + " ASC "
	return result
}

func checkDeletion(result sql.Result) (bool, error) {
	rows, err := result.RowsAffected()

	if err != nil {
		return false, err
	}

	if rows > 1 {
		return false, fmt.Errorf("too many rows deleted: %d, expected 1", rows)
	} else if rows == 0 {
		return false, errors.New("no row deleted")
	}

	return true, nil
}

// Scannable is an interface implemented for sql.Row and sql.Rows
type Scannable interface {
	Scan(...interface{}) error
}
