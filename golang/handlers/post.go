package handlers

import (
	"net/http"
)

func PostGetDetails(w http.ResponseWriter, r *http.Request) {
	messageId := mux.Vars(r)["id"]
	w.WriteHeader(http.StatusOK)
}

func PostUpdateDetails(w http.ResponseWriter, r *http.Request) {
	messageId := mux.Vars(r)["id"]
	w.WriteHeader(http.StatusOK)
}

