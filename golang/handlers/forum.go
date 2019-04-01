package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func ForumCreate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func ForumCreateThread(w http.ResponseWriter, r *http.Request) {
	threadName := mux.Vars(r)["slug"]
	_ = threadName
	w.WriteHeader(http.StatusOK)
}

func ForumGetInfo(w http.ResponseWriter, r *http.Request) {
	threadName := mux.Vars(r)["slug"]
	_ = threadName
	w.WriteHeader(http.StatusOK)
}

func ForumGetThreads(w http.ResponseWriter, r *http.Request) {
	threadName := mux.Vars(r)["slug"]
	_ = threadName
	w.WriteHeader(http.StatusOK)
}

func ForumGetUsers(w http.ResponseWriter, r *http.Request) {
	threadName := mux.Vars(r)["slug"]
	_ = threadName
	w.WriteHeader(http.StatusOK)
}