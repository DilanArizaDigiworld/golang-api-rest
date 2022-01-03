package routes

import (
	"golang-rest-api/controllers"

	"github.com/gorilla/mux"
)

func SetAuthRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/auth/api").Subrouter()
	subRoute.HandleFunc("/login", controllers.Login).Methods("POST")
}
