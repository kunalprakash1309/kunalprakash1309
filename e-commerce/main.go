package main

import (
	"encoding/json"
	"log"
	"net/http"
	"context"
	"time"

	"github.com/gorilla/mux"
	// mgo "gopkg.in/mgo.v2"
	// _ "gopkg.in/mgo.v2/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DB struct to initialize the database
type DB struct {
	Database *mongo.Database
}


// User to hold the users data
type User struct {
	ID        primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
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

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := db.Database.Collection("users")
	ObjectID, _ := primitive.ObjectIDFromHex(vars["id"]) 
	statement := bson.M{"_id": ObjectID}
	err := collection.FindOne(ctx, statement).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No Such User"))
		// log.Fatalln("error in retrieving the data")
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response, _ := json.Marshal(user)
		// if err != nil {
		// 	w.Write([]byte("error in converting struct to json "))
		// 	log.Fatalln("error in converting struct to json ", err)
		// }
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
		// user.ID = bson.NewObjectId()
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		
		collection := db.Database.Collection("users")
		result, err := collection.InsertOne(ctx, user)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Cannot insert user into database"))

			log.Fatalln(err.Error())
		} else {
			w.Header().Set("Content-Type", "application/json")
			response, err := json.Marshal(result)
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

// UpdateUser to update the user data from databse
func (db *DB) UpdateUser(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection := db.Database.Collection("users")
	ObjectID, _:= primitive.ObjectIDFromHex(vars["id"])
	statement := bson.M{"_id": ObjectID}
	update := bson.M{}

	result, err := collection.UpdateOne(ctx, statement)
}

// DeleteUser to delete user data from database
func (db *DB) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection := db.Database.Collection("users")
	ObjectID, _ := primitive.ObjectIDFromHex(vars["id"]) 
	statement := bson.M{"_id": ObjectID}

	result, err := collection.DeleteOne(ctx, statement)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-type", "text/plain")
		w.Write([]byte("Deleted Succesfully!"))
		log.Printf("Delete 1 %v", result)
		log.Printf("Delete 1 %v", result.DeletedCount)
	}
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Println(err)
	}
	err = client.Connect(ctx)
	err = client.Ping(context.TODO(), nil)
	log.Println("Database connected")

	database := client.Database("Ecommerce")
	db := &DB{
		Database: database,
	}
	

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
