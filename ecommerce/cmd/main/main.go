package main

import (
	"log"
	"net/http"
	"context"

	"github.com/gorilla/mux"
	
	"github.com/kunalprakash1309/ecommerce/pkg/config"
	"github.com/kunalprakash1309/ecommerce/pkg/router"
)


func main() {
	ctx := context.Background()
	config.Setup(ctx, "mongodb://localhost:27017")

	r := mux.NewRouter()
	router.UsersHandler(r)
	router.ItemsHandler(r)
	srv := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: r,
	}

	log.Fatalln(srv.ListenAndServe())
}



