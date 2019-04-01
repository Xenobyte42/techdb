package handlers

import (
	"net/http"

	_ "github.com/gorilla/mux"
)

func ClearDatabase(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func GetDatabaseInfo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
