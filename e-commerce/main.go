package main

import (
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"encoding/json"
	"log"
	"net/http"
	"context"
	"time"

	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB struct to initialize the database
type DB struct {
	Database *mgo.Database
}

// User to hold the users data
type User struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Email     string        `json:"email" bson:"email"`
	FirstName string        `json:"firstName" bson:"firstName"`
	LastName  string        `json:"lastName" bson:"lastName"`
	Password  string        `json:"password,omitempty" bson:"password,omitempty"`
	Address   []Address     `json:"address" bson:"address"`
}

// Address struct to hold the address data
type Address struct {
	State  string `json:"state" bson:"state"`
	City   string `json:"city" bson:"city"`
	Sector string `json:"sector" bson:"sector"`
}

// GetUser to get the user using the user ID
func (db *DB) GetUser(w http.ResponseWriter, r *http.Request) {
	// get all parameters in form of map
	vars := mux.Vars(r)

	var user User

	collection := db.Database.C("users")
	statement := bson.M{"_id": bson.ObjectIdHex(vars["id"])}
	err := collection.Find(statement).One(&user)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No Such User"))
		// log.Fatalln("error in retrieving the data")
	} else {
		w.WriteHeader(http.StatusOK)
		response, err := json.Marshal(user)
		if err != nil {
			w.Write([]byte("error in converting struct to json "))
			log.Fatalln("error in converting struct to json ", err)
		}
		w.Header().Set("Content-type", "application/json")
		w.Write(response)
	}

}

// PostUser to post the user into database
func (db *DB) PostUser(w http.ResponseWriter, r *http.Request) {

	contentType := r.Header.Get("content-type")

	if contentType == "application/json" {
		var user User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Println("error in encoding the post body", err)
		}

		// store id
		user.ID = bson.NewObjectId()
		
		collection := db.Database.C("users")
		err = collection.Insert(user)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Cannot insert user into database"))

			log.Fatalln(err.Error())
		} else {
			w.Header().Set("Content-type", "application/json")
			response, err := json.Marshal(user)
			if err != nil {
				log.Println("error in converting struct to json", err)
			}
	
			w.Write(response)
		}
	} else {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("unsupported format"))
	}

}

// DeleteUser to delete user data from database
func (db *DB) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	collection := db.Database.C("users")
	statement := bson.M{"_id": bson.ObjectIdHex(vars["id"])}

	err := collection.Remove(statement)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-type", "text/plain")
		w.Write([]byte("Deleted Succesfully!"))
	}
}

func main() {
	// session, err := mgo.Dial("127.0.0.1")
	// if err != nil {
	// 	log.Println("session creatin error", err)
	// }

	// database := session.DB("Ecommerce")
	// db := &DB{
	// 	Database: database,
	// }

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.NewClient(options.Client().ApplyURI(""))
	if err != nil {
		fmt.Println(err)
	}
	err = client.Connect(ctx)
	err = client.Ping(context.TODO(), nil)
	fmt.Println("Database connected")
	

	mux := mux.NewRouter()
	mux.HandleFunc("/user", db.PostUser).Methods("POST")
	mux.HandleFunc("/user/{id:[0-9a-zA-Z]*}", db.GetUser).Methods("GET")
	mux.HandleFunc("/user/{id:[0-9a-zA-Z]*}", db.DeleteUser).Methods("DELETE")
	srv := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	log.Fatalln(srv.ListenAndServe())
}
