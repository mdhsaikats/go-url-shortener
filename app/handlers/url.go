package handler

import (
	"encoding/json"
	"net/http"

	"main.go/app/models"
)

func GenerateShortCode(w http.ResponseWriter,r *http.Request){
	var url models.Url
	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil{
		http.Error(w,"Invalid json request",http.StatusBadRequest)
		return
	}
	//TODO: need to work from here
}