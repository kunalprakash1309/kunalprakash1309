package models

import (
	"context"
	"encoding/json"
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/kunalprakash1309/ecommerce/pkg/config"
)


type Item struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Category    string             `json:"category,omitempty" bson:"category,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Price       int                `json:"price,omitempty" bson:"price,omitempty"`
}

func GetItemById(userId string) (Item, error) {
	DB = config.DB

	var item Item

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection := DB.Collection("items")
	ObjectID, _ := primitive.ObjectIDFromHex(userId)
	statement := bson.M{"_id": ObjectID}
	err := collection.FindOne(ctx, statement).Decode(&item)

	return item, err
}

func CreateItem(decoder *json.Decoder) (*mongo.InsertOneResult, error) {
	DB = config.DB

	var item Item

	err := decoder.Decode(&item)
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection := DB.Collection("items")
	result, err := collection.InsertOne(ctx, item)

	return result, err
}

func DeleteItemById(userID string) (*mongo.DeleteResult, error){
	DB := config.DB
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection := DB.Collection("items")
	ObjectID, _ := primitive.ObjectIDFromHex(userID) 
	statement := bson.M{"_id": ObjectID}

	result, err := collection.DeleteOne(ctx, statement)

	return result, err
}