package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/therealsangwoohan/channels-backend-api2/handlers"
)

func NewRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/api/channels", handlers.GetAllChannels)
	router.Get("/api/channels/{id}", handlers.GetChannel)
	router.Post("/api/channels", handlers.CreateChannel)

	router.Get("/api/messages", handlers.GetAllMessages)
	return router
}
