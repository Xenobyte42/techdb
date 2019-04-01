package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Xenobyte42/techdb/golang/api"
)

func ClearDatabase(w http.ResponseWriter, r *http.Request) {
	err := api.Db.ClearDatabase()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetDatabaseInfo(w http.ResponseWriter, r *http.Request) {
	response, err := api.Db.GetInfo()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	body, _ := json.Marshal(response)
	w.Write(body)
}
