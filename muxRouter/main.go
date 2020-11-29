package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

// ArticleHandler is a funciton handler
func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	// mux.vars returns all path parameters as a map
	vars := mux.Vars(r)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category is :%v\n", vars["category"])
	fmt.Fprintf(w, "ID is: %v\n ", vars["id"])
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to page")
}

func main() {
	// create a new router
	r := mux.NewRouter()
	// attach a path with handler
	r.HandleFunc("/", index)
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler).Name("articleRoute")

	// url, _ := r.Get("articleRoute").URL("category", "books", "id", "123")
	// fmt.Print(url)

	// srv := &http.Server{
	// 	Handler: r,
	// 	Addr: "127.0.0.1:8000",
	// 	// Good practice: enforce timeouts for servers you create!
	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout: 15 * time.Second,
	// }
	log.Fatal(http.ListenAndServe(":8080", r))
}

