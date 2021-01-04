package router

import (
	"github.com/gorilla/mux"
	"github.com/kunalprakash1309/ecommerce/pkg/controllers"

)


// UsersHandler for handling users route in main.go
var ItemsHandler= func(router *mux.Router){

	router.HandleFunc("/item", controllers.PostItem).Methods("POST")
	router.HandleFunc("/item/{id:[0-9a-zA-Z]*}", controllers.GetItem).Methods("GET")
	router.HandleFunc("/item/{id:[0-9a-zA-Z]*}", controllers.DeleteItem).Methods("DELETE")
}