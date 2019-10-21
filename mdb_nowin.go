// +build !windows

package mdb

import (
	"database/sql"
)

//
// Implementing new open function
//
func Open(driver, filen string) (*sql.DB, error) {
	var err error
	var db *sql.DB

	db, err = sql.Open(driver, filen)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
