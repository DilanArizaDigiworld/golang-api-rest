package controllers

import (
	"encoding/json"
	"golang-rest-api/common"
	"golang-rest-api/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAll(writer http.ResponseWriter, request *http.Request) {
	personas := []models.Persona{}
	db := common.GetConnection()
	defer db.Close()

	db.Find(&personas)
	json, _ := json.Marshal(personas)
	common.SendResponse(writer, http.StatusOK, json)
}

func Get(writer http.ResponseWriter, request *http.Request) {
	persona := models.Persona{}
	id := mux.Vars(request)["id"]

	db := common.GetConnection()
	defer db.Close()

	db.Find(&persona, id)

	if persona.ID > 0 {
		json, _ := json.Marshal(persona)
		common.SendResponse(writer, http.StatusOK, json)
	} else {
		common.SendError(writer, http.StatusNotFound)
	}
}

func Save(writer http.ResponseWriter, request *http.Request) {
	persona := models.Persona{}

	db := common.GetConnection()
	defer db.Close()

	error := json.NewDecoder(request.Body).Decode(&persona)

	if error != nil {
		log.Fatal(error)
		common.SendError(writer, http.StatusBadRequest)
		return
	}

	error = db.Save(&persona).Error

	if error != nil {
		log.Fatal(error)
		common.SendError(writer, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(persona)

	common.SendResponse(writer, http.StatusCreated, json)
}

func Delete(writer http.ResponseWriter, request *http.Request) {
	persona := models.Persona{}
	id := mux.Vars(request)["id"]

	db := common.GetConnection()
	defer db.Close()

	db.Find(persona, id)

	if persona.ID > 0 {
		db.Delete(persona)
		common.SendResponse(writer, http.StatusOK, []byte(`{}`))
	} else {
		common.SendError(writer, http.StatusBadRequest)
	}
}
