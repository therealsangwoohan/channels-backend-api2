package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/therealsangwoohan/channels-backend-api2/routes"
)

func main() {
	connectionString := "postgres://tepudsxx:BfoCW1P4Z3PeIy_mYzWJNnKiliVkgVwu@berry.db.elephantsql.com/tepudsxx"
	database, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
	router := routes.NewRouter(database)
	http.ListenAndServe("127.0.0.1:8080", router.ChiMux)
}
