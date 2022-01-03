package routes

import (
	"golang-rest-api/controllers"

	"github.com/gorilla/mux"
)

func SetUsersRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/users/api").Subrouter()
	subRoute.HandleFunc("/all", controllers.GetAll).Methods("GET")
	subRoute.HandleFunc("/save", controllers.Save).Methods("POST")
	subRoute.HandleFunc("/delete/{id}", controllers.Delete).Methods("DELETE")
	subRoute.HandleFunc("/find/{id}", controllers.Get).Methods("GET")
}
