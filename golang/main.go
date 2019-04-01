package main

import (
	"log"
	"net/http"

	"github.com/Xenobyte42/techdb/golang/api"
	"github.com/Xenobyte42/techdb/golang/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func initForumRouting(r *mux.Router) {
	r.HandleFunc("/create", handlers.ForumCreate).Methods("POST")
	r.HandleFunc("/{slug}/create", handlers.ForumCreateThread).Methods("POST")
	r.HandleFunc("/{slug}/details", handlers.ForumGetInfo).Methods("GET")
	r.HandleFunc("/{slug}/threads", handlers.ForumGetThreads).Methods("GET")
	r.HandleFunc("/{slug}/users", handlers.ForumGetUsers).Methods("GET")
}

func initPostRouting(r *mux.Router) {
	r.HandleFunc("/{id}/details", handlers.PostGetDetails).Methods("GET")
	r.HandleFunc("/{id}/details", handlers.PostUpdateDetails).Methods("GET")
}

func initServiceRouting(r *mux.Router) {
	r.HandleFunc("/clear", handlers.ClearDatabase).Methods("POST")
	r.HandleFunc("/status", handlers.GetDatabaseInfo).Methods("GET")
}

func initThreadRouting(r *mux.Router) {
	r.HandleFunc("/{slug_or_id}/create", handlers.ThreadCreate).Methods("POST")
	r.HandleFunc("/{slug_or_id}/details", handlers.ThreadGetDetails).Methods("GET")
	r.HandleFunc("/{slug_or_id}/details", handlers.ThreadUpdateDetails).Methods("POST")
	r.HandleFunc("/{slug_or_id}/posts", handlers.ThreadGetMessages).Methods("GET")
	r.HandleFunc("/{slug_or_id}/vote", handlers.ThreadVote).Methods("POST")
}

func initUserRouting(r *mux.Router) {
	r.HandleFunc("/{nickname}/create", handlers.UserCreate).Methods("POST")
	r.HandleFunc("/{nickname}/profile", handlers.UserGetInfo).Methods("GET")
	r.HandleFunc("/{nickname}/profile", handlers.UserUpdateInfo).Methods("POST")
}

func initRouting() *mux.Router {
	r := mux.NewRouter()

	initForumRouting(r.PathPrefix("/forum").Subrouter())
	initPostRouting(r.PathPrefix("/post").Subrouter())
	initServiceRouting(r.PathPrefix("/service").Subrouter())
	initThreadRouting(r.PathPrefix("/thread").Subrouter())
	initUserRouting(r.PathPrefix("/user").Subrouter())

	return r
}

func main() {
	err := api.Db.Open()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Succesfully connected")
	defer api.Db.Close()

	err = api.Db.Init()
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Succesfully inited. Start listening...")
	router := initRouting()

	http.ListenAndServe(":5000", router)
}
