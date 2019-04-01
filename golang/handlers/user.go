package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func UserCreate(w http.ResponseWriter, r *http.Request) {
	nickname := mux.Vars(r)["nickname"]
	_ = nickname
	w.WriteHeader(http.StatusOK)
}

func UserGetInfo(w http.ResponseWriter, r *http.Request) {
	nickname := mux.Vars(r)["nickname"]
	_ = nickname
	w.WriteHeader(http.StatusOK)
}

func UserUpdateInfo(w http.ResponseWriter, r *http.Request) {
	nickname := mux.Vars(r)["nickname"]
	_ = nickname
	w.WriteHeader(http.StatusOK)
}
