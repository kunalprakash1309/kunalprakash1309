package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// A sample use
var user = User{
	ID:       1,
	Username: "kunal",
	Password: "password",
}

func CreateToken(userid uint64) (string, error) {
	var err error

	// Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") // This should be in env file
	atClaims := jwt.MapClaims{
		"authorized": true,
		"user_id":    userid,
		"exp":        time.Now().Add(time.Minute * 15).Unix(),
	}

	fmt.Println("exp :- ", atClaims["exp"])

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}

	return token, nil
}

func Login(w http.ResponseWriter, r *http.Request) {
	var u User
	u.ID = 1
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Invalid json provided", http.StatusUnprocessableEntity)
	}

	// compare the user from the request
	if user.Username != u.Username || user.Password != u.Password {
		http.Error(w, "Please provide valid login details", http.StatusUnauthorized)
		return
	}

	token, err := CreateToken(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	result, err := json.Marshal(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	w.Write(result)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", Login).Methods("POST")
	srv := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: r,
	}
	log.Fatalln(srv.ListenAndServe())
}
