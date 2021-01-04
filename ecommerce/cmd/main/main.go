package main

import (
	"log"
	"net/http"
	"context"
	"html/template"

	"github.com/gorilla/mux"
	
	"github.com/kunalprakash1309/ecommerce/pkg/config"
	"github.com/kunalprakash1309/ecommerce/pkg/router"
)

var Tpl *template.Template


func init() {

}

func main() {
	ctx := context.Background()
	config.Setup(ctx, "mongodb://localhost:27017")
	config.Template()

	r := mux.NewRouter()
	router.UsersHandler(r)
	router.ItemsHandler(r)
	srv := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: r,
	}

	log.Fatalln(srv.ListenAndServe())
}



