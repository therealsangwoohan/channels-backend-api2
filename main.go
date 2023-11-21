package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"github.com/therealsangwoohan/channels-backend-api2/routes"
)

func main() {
	router := chi.NewRouter()

	connectionString := "postgres://tepudsxx:BfoCW1P4Z3PeIy_mYzWJNnKiliVkgVwu@berry.db.elephantsql.com/tepudsxx"
	database, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	println(database)
	routes.SetChannelRoutes(router)
	routes.SetMessageRoutes(router)
	http.ListenAndServe("127.0.0.1:8080", router)
}
