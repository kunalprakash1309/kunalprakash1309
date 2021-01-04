package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gorilla/mux"   
)

var secretKey = []byte("secret_key")
var users = map[string]string{"kunal": "passme", "admin": "password"}

// Response is a representation of json response for jwt
type Response struct {
	Token string `json:"token"`
	Status string `json:"status"`
}

// HealthcheckHandler returns date and time
func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	tokenString, err := request.HeaderExtractor{"access_token"}.ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
		// Don't forget to validate the alg in what you expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Access Denied; Please check the access token"))
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		response := make(map[string]string)
		
		response["time"] = time.Now().String()
		response["user"] = claims["username"].(string)
		responseJSON, _ := json.Marshal(response)
		w.Write(responseJSON) 
	} else{
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(err.Error()))
	}	
}

// getTokenHandler validates the user credentials
func getTokenHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Please pass the data as URL form encoded", http.StatusBadRequest)
		return
	}

	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	fmt.Println(username)
	fmt.Println(password)

	fmt.Println("Users :- ", users[username])

	if originalPassword, ok := users[username]; ok {
		if password == originalPassword {
			// Create a claims map
			claims := jwt.MapClaims{
				"username": username,
				"ExpiresAt": 15000,
				"IssuedAt": time.Now().Unix(),
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString(secretKey)
			if err != nil {
				w.WriteHeader(http.StatusBadGateway)
				w.Write([]byte(err.Error()))
			}

			response := Response{Token: tokenString, Status: "success"}
			responseJSON, _ := json.Marshal(response)
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write(responseJSON)
		} else {
			http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
			return
		}
	} else {
		http.Error(w, "user is not found", http.StatusNotFound)
		return
	}

}



func main() {
	r := mux.NewRouter()
	r.HandleFunc("/getToken", getTokenHandler)
	r.HandleFunc("/healthcheck", HealthcheckHandler)
	http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr: "127.0.0.1:8000",
	}
	log.Fatalln(srv.ListenAndServe())
}