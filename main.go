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
	http.ListenAndServe("0.0.0.0:80", router)
}
