package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Xenobyte42/techdb/golang/api"
	"github.com/Xenobyte42/techdb/golang/models"
	"github.com/Xenobyte42/techdb/golang/utils"
	"github.com/gorilla/mux"
)

/*
curl -X POST -H "Content-Type: application/json" \
 -d '{"about": "This is the day you will always remember as the day that you almost caught Captain Jack Sparrow!","email": "captaina@blackpearl.sea","fullname": "Captain Jack Sparrow"}' \
 http://localhost:5000/user/vasya/create
*/

func UserCreate(w http.ResponseWriter, r *http.Request) {
	nickname := mux.Vars(r)["nickname"]
	exists, err := api.Db.UserExists(nickname)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if exists {
		utils.RespondError(w, http.StatusConflict, "User already exists")
		return
	}

	requestData := &models.UserUpdate{}
	err = utils.GetData(requestData, r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = api.Db.CreateUser(requestData, nickname)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseData := &models.UserModel{}
	responseData.About = requestData.About
	responseData.Email = requestData.Email
	responseData.Fullname = requestData.Fullname
	responseData.Nickname = nickname
	body, _ := json.Marshal(responseData)
	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}

func UserGetInfo(w http.ResponseWriter, r *http.Request) {
	nickname := mux.Vars(r)["nickname"]

	userData := &models.UserModel{}
	exists, err := api.Db.GetUserData(userData, nickname)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !exists {
		utils.RespondError(w, http.StatusNotFound, "User is gone")
		return
	}
	body, _ := json.Marshal(userData)
	w.Write(body)
}

func UserUpdateInfo(w http.ResponseWriter, r *http.Request) {
	nickname := mux.Vars(r)["nickname"]
	_ = nickname
	w.WriteHeader(http.StatusOK)
}
