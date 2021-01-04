package main

import (
	"html/template"
	"log"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"

)

var tpl *template.Template

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func signin(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
	// w.Header().Set("Content-Type", "application/json")
	tpl.ExecuteTemplate(w, "index.gohtml", nil)


}

func home(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	email := r.FormValue("email")
	password := r.FormValue("password")

	user := &User{
		Email:     email,
		Password: password,
	}

	tpl.ExecuteTemplate(w, "home.gohtml", user)

	fmt.Printf("Your username :- %v \n", user.Email)
	fmt.Printf("Your password :- %v \n", user.Password)
}

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/login", signin).Methods("GET")
	router.HandleFunc("/home", home).Methods("POST")

	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: router,
	}

	log.Fatalln(server.ListenAndServe())
}
