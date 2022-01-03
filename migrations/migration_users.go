package main

import (
	"golang-rest-api/common"
	"golang-rest-api/models"
	"log"
)

func main() {
	db := common.GetConnection()
	defer db.Close()

	log.Print(`Migrando tabla "Users"...`)

	if !db.HasTable(&models.Users{}) {
		db.AutoMigrate(&models.Users{})
	} else {
		log.Print(`La tabla "Users" ya existe, eliminela de la base de datos para continuar la migración total de la tabla`)
	}
}
