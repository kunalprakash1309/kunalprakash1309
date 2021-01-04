package models


import (
	"context"
	"time"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/kunalprakash1309/ecommerce/pkg/config"
)

var DB *mongo.Database

// User to hold the users data
type User struct {
	ID        primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Email     string        `json:"email,omitempty" bson:"email,omitempty"`
	// FirstName string        `json:"firstName,omitempty" bson:"firstName,omitempty"`
	// LastName  string        `json:"lastName,omitempty" bson:"lastName,omitempty"`
	Password  string        `json:"password,omitempty" bson:"password,omitempty"`
	// Item      Item          `json:"item,omitempty" bson:"item,omitempty"`
}

// Address struct to hold the address data
type Address struct {
	State  string `json:"state" bson:"state"`
	City   string `json:"city" bson:"city"`
	Sector string `json:"sector" bson:"sector"`
}

func GetUserById(userId string) (User, error) {
	DB = config.DB

	var user User
	
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection := DB.Collection("users")
	ObjectID, _ := primitive.ObjectIDFromHex(userId) 
	statement := bson.M{"_id": ObjectID}
	err := collection.FindOne(ctx, statement).Decode(&user)

	return user, err
}

func CreateUser(decoder *json.Decoder) (*mongo.InsertOneResult, error) {
	DB = config.DB

	var user User

	err := decoder.Decode(&user)
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	
	collection := DB.Collection("users")
	result, err := collection.InsertOne(ctx, user)

	return result, err
}

func UpdateUserById(userID string, decoder *json.Decoder) (*mongo.UpdateResult, error) {
	DB := config.DB

	var user User

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err := decoder.Decode(&user)
	if err != nil {
		return nil, err
	}

	collection := DB.Collection("users")
	ObjectID, _:= primitive.ObjectIDFromHex(userID)
	filter := bson.M{"_id": ObjectID}
	update := bson.M{"$set": &user}

	result, err := collection.UpdateOne(ctx, filter, update)

	return result, err
}


func DeleteUserById(userID string) (*mongo.DeleteResult, error){
	DB := config.DB
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection := DB.Collection("users")
	ObjectID, _ := primitive.ObjectIDFromHex(userID) 
	statement := bson.M{"_id": ObjectID}

	result, err := collection.DeleteOne(ctx, statement)

	return result, err
}