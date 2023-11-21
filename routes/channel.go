package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetChannelRoutes(router *chi.Mux) {
	router.Get("/api/channels", getAllChannels)
}

func getAllChannels(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("inside getAllChannels"))
}
