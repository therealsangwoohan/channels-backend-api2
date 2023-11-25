package main

import (
	"net/http"

	_ "github.com/lib/pq"
	"github.com/therealsangwoohan/channels-backend-api2/database"
	"github.com/therealsangwoohan/channels-backend-api2/router"
)

func main() {
	database.ConnectDB()
	router := router.NewRouter()
	http.ListenAndServe("127.0.0.1:8080", router)
}
