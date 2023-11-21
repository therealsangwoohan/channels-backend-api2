package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetMessageRoutes(router *chi.Mux) {
	router.Get("/api/messages", getAllMessages)
}

func getAllMessages(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("inside getAllMessages"))
}
