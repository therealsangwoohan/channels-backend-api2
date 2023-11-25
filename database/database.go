package database

import (
	"database/sql"
	"log"
)

var Database *sql.DB

func ConnectDB() {
	connectionString := "postgres://tepudsxx:BfoCW1P4Z3PeIy_mYzWJNnKiliVkgVwu@berry.db.elephantsql.com/tepudsxx"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	Database = db
}
