package router

import (
	"github.com/gorilla/mux"
	"github.com/kunalprakash1309/ecommerce/pkg/controllers"

)


// UsersHandler for handling users route in main.go
var UsersHandler= func(router *mux.Router){

	router.HandleFunc("/user", controllers.PostUser).Methods("POST")
	router.HandleFunc("/user/{id:[0-9a-zA-Z]*}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/user/{id:[0-9a-zA-Z]*}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id:[0-9a-zA-Z]*}", controllers.DeleteUser).Methods("DELETE")
}