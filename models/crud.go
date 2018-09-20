package models

import (
	"database/sql"
)

func open() *sql.DB {
	driver := ""
	dataSourceName := ""

	db, err := sql.Open(driver, dataSourceName)

	if err != nil {

	}

	return db
}

func execute() {

}

func query(q QuerySet) {

}

// func create(query string) {
//
// }
//
// func read(query string) {
//
// }
//
// func update(query string) {
//
// }

// func delete(query string) {
//
// }
