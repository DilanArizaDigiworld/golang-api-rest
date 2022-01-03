package controllers

import (
	"encoding/json"
	"golang-rest-api/common"
	"golang-rest-api/models"
	"golang-rest-api/secure"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAll(writer http.ResponseWriter, request *http.Request) {
	users := []models.Users{}
	db := common.GetConnection()
	defer db.Close()

	db.Find(&users)
	json, _ := json.Marshal(users)
	common.SendResponse(writer, http.StatusOK, json)
}

func Get(writer http.ResponseWriter, request *http.Request) {
	user := models.Users{}
	id := mux.Vars(request)["id"]

	db := common.GetConnection()
	defer db.Close()

	db.Find(&user, id)

	if user.ID > 0 {
		json, _ := json.Marshal(user)
		common.SendResponse(writer, http.StatusOK, json)
	} else {
		common.SendError(writer, http.StatusNotFound)
	}
}

func Save(writer http.ResponseWriter, request *http.Request) {
	user := models.Users{}

	db := common.GetConnection()
	defer db.Close()

	error := json.NewDecoder(request.Body).Decode(&user)

	if error != nil {
		log.Fatal(error)
		common.SendError(writer, http.StatusBadRequest)
		return
	}

	// Hash password
	userPassword := user.Password
	hash, _ := secure.HashPassword(userPassword)
	user.Password = hash

	error = db.Save(&user).Error

	if error != nil {
		log.Fatal(error)
		common.SendError(writer, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(user)

	common.SendResponse(writer, http.StatusCreated, json)
}

func Delete(writer http.ResponseWriter, request *http.Request) {
	user := models.Users{}
	id := mux.Vars(request)["id"]

	db := common.GetConnection()
	defer db.Close()

	db.Find(&user, id)

	if user.ID > 0 {
		db.Delete(user)
		common.SendResponse(writer, http.StatusOK, []byte(`{}`))
	} else {
		common.SendError(writer, http.StatusBadRequest)
	}
}
