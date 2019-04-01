package main

import (
	"net/http"
	"log"
	"os"
	"io"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type LoggerStore struct (
	store map[string]*Logger
)

func (self *LoggerStore) createDefaultLogger(name string) {
	self.store[name] = log.New(os.Stdout, name + ": ", log.Lshortfile || log.Ltime)
}

func (self *LoggerStore) GetLogger(name string) *Logger {
	if _, ok := store[name]; !ok {
		self.createDefaultLogger(name)
	}
	return store[name]
}

var (
	LogStore LoggerStore
)

func initForumRouting(r *Router) {
	r.HandleFunc("/create", ForumCreate).Methods("POST")
	r.HandleFunc("/{slug}/create", ForumCreateThread).Methods("POST")
	r.HandleFunc("/{slug}/details", ForumGetInfo).Methods("GET")
	r.HandleFunc("/{slug}/threads", ForumGetThreads).Methods("GET")
	r.HandleFunc("/{slug}/users"), ForumGetUsers).Methods("GET")
}

func initPostRouting(r *Router) {
	r.HandleFunc("/{id}/details", PostGetDetails).Methods("GET")
	r.HandleFunc("/{id}/details", PostUpdateDetails).Methods("GET")
}

func initServiceRouting(r *Router) {
	r.HandleFunc("/clear", ClearDatabase).Methods("POST")
	r.HandleFunc("/status", GetDatabaseInfo).Methods("GET")
}

func initThreadRouting(r *Router) {
	r.HandleFunc("/{slug_or_id}/create", ThreadCreate).Methods("POST")
	r.HandleFunc("/{slug_or_id}/details", ThreadGetDetails).Methods("GET")
	r.HandleFunc("/{slug_or_id}/details", ThreadUpdateDetails).Methods("POST")
	r.HandleFunc("/{slug_or_id}/posts", ThreadGetMessages).Methods("GET")
	r.HandleFunc("/{slug_or_id}/vote", ThreadVote).Methods("POST")
}

func initUserRouting(r * Router) {
	r.HandleFunc("/{nickname}/create", UserCreate).Methods("POST")
	r.HandleFunc("/{nickname}/profile", UserGetInfo).Methods("GET")
	r.HandleFunc("/{nickname}/profile", UserUpdateInfo).Methods("POST")
}

func initRouting() {
	r := mux.NewRouter()

	initForumRouting(r.PathPrefix("/forum").Subrouter())
	initPostRouting(r.PathPrefix("/post").Subrouter())
	initServiceRouting(r.PathPrefix("/service").Subrouter())
	initThreadRouting(r.PathPrefix("/thread").Subrouter())
	initUserRouting(r.PathPrefix("/user").Subrouter())

	return r
}

func main() {
	logger := LogStore.GetLogger("main")
	logger.Print("Initialize routes...")
	router := initRouting()

	logger.Print("Start listening...")
	http.ListenAndServe(":5000", router)
}
