package handlers

import (
	"net/http"
)

func UserCreate(w http.ResponseWriter, r *http.Request) {
	nickname := mux.Vars(r)["nickname"]
	w.WriteHeader(http.StatusOK)
}

func UserGetInfo(w http.ResponseWriter, r *http.Request) {
	nickname := mux.Vars(r)["nickname"]
	w.WriteHeader(http.StatusOK)
}

func UserUpdateInfo(w http.ResponseWriter, r *http.Request) {
	nickname := mux.Vars(r)["nickname"]
	w.WriteHeader(http.StatusOK)
}
