package main

import (
	"fmt"
	"log"
	"net/http"
	_ "time"
	"github.com/gorilla/mux"
)

func QueryHandler(w http.ResponseWriter, r *http.Request) {
	//Fetch query parameters as a map
	queryParams := r.URL.Query()

	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Got paramter id: %s\n", queryParams["id"][0])
	fmt.Fprintf(w, "Got paramter category: %s", queryParams["category"])
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/articles", QueryHandler)
	// server := &http.Server{
	// 	Handler: r,
	// 	Addr: "127.0.0.1:8080",
	// }

	log.Fatal(http.ListenAndServe(":8080", r))
}