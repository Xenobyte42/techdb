package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Xenobyte42/techdb/golang/models"
)

func RespondError(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)
	response := &models.ErrorResponse{}
	response.Message = msg
	body, _ := json.Marshal(response)
	w.Write(body)
}

func GetData(data interface{}, r *http.Request) error {
	body := r.Body
	defer body.Close()

	byteBody, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteBody, data)
	return err
}
