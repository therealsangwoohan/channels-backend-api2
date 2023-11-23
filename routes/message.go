package routes

import (
	"net/http"
)

func (router *Router) SetMessageRoutes() {
	router.ChiMux.Get("/api/messages", router.getAllMessages)
}

func (router *Router) getAllMessages(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("inside getAllMessages"))
}
