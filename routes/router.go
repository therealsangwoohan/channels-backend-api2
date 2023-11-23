package routes

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
)

type Router struct {
	ChiMux   *chi.Mux
	Database *sql.DB
}

func NewRouter(database *sql.DB) *Router {
	router := &Router{
		ChiMux:   chi.NewRouter(),
		Database: database,
	}
	router.SetChannelRoutes()
	router.SetMessageRoutes()
	return router
}
