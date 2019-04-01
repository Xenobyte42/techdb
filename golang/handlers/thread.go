package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func ThreadCreate(w http.ResponseWriter, r *http.Request) {
	slugOrId := mux.Vars(r)["slug_or_id"]
	_ = slugOrId
	w.WriteHeader(http.StatusOK)
}

func ThreadGetDetails(w http.ResponseWriter, r *http.Request) {
	slugOrId := mux.Vars(r)["slug_or_id"]
	_ = slugOrId
	w.WriteHeader(http.StatusOK)
}

func ThreadUpdateDetails(w http.ResponseWriter, r *http.Request) {
	slugOrId := mux.Vars(r)["slug_or_id"]
	_ = slugOrId
	w.WriteHeader(http.StatusOK)
}

func ThreadGetMessages(w http.ResponseWriter, r *http.Request) {
	slugOrId := mux.Vars(r)["slug_or_id"]
	_ = slugOrId
	w.WriteHeader(http.StatusOK)
}

func ThreadVote(w http.ResponseWriter, r *http.Request) {
	slugOrId := mux.Vars(r)["slug_or_id"]
	_ = slugOrId
	w.WriteHeader(http.StatusOK)
}
