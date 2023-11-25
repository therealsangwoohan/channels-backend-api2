package handlers

import (
	"net/http"
)

func GetAllMessages(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("inside getAllMessages"))
}
