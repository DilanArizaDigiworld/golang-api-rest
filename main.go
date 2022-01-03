package main

import (
	"golang-rest-api/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	routes.SetUsersRoutes(router)
	routes.SetAuthRoutes(router)

	server := http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	log.Println("Server on running in http://localhost:8000")
	log.Println(server.ListenAndServe())

}
