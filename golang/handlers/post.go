package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func PostGetDetails(w http.ResponseWriter, r *http.Request) {
	messageId := mux.Vars(r)["id"]
	_ = messageId
	w.WriteHeader(http.StatusOK)
}

func PostUpdateDetails(w http.ResponseWriter, r *http.Request) {
	messageId := mux.Vars(r)["id"]
	_ = messageId
	w.WriteHeader(http.StatusOK)
}
