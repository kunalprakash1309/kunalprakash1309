package main

import (
	"net/http"
	"log"
	"encoding/json"

	"github.com/gorilla/mux"
	// mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// User to hold the users data
// work on err 
// means it sends output while i provide different data type on POST
type User struct {
	ID        bson.ObjectId    `json:"id" bson:"_id,omitempty"`
	Email     string `json:"email" bson:"email"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`
	Password  string `json:"password" bson:"password"`
	Address []Address `json:"address" bson:"address"`
}

// Address struct to hold the address data
type Address struct {
	State  string `json:"state" bson:"state"`
	City   string `json:"city" bson:"city"`
	Sector string `json:"sector" bson:"sector"`
}

func see(w http.ResponseWriter, r *http.Request){

	contentType := r.Header.Get("content-type")
	log.Println(contentType)

	if contentType == "application/json" {
		var user User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Println("error in encoding the post body", err)
		}
	
		// store id
		user.ID = bson.NewObjectId()
		
		w.Header().Set("Content-type", "application/json")
		response, err := json.Marshal(user)
		if err != nil {
			log.Println("error in convertinf struct to json", err)
		}
	
		w.Write(response)
	} else {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("unsupported format"))
	}

}


func main() {

	mux := mux.NewRouter()
	mux.HandleFunc("/user", see).Methods("POST")
	srv := &http.Server{
		Addr: "127.0.0.1:8080",
		Handler: mux,
	}

	log.Fatalln(srv.ListenAndServe())
}