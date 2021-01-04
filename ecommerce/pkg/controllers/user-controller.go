package controllers


import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/kunalprakash1309/ecommerce/pkg/models"
)

// GetUser to get the user from database
func GetUser(w http.ResponseWriter, r *http.Request){
	// get all parameters in form of map
	vars := mux.Vars(r)
	userID := vars["id"]

	user, err := models.GetUserById(userID)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No Such User"))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response, _ := json.Marshal(user)
		w.Write(response)
	}
}

// PostUser to post the user into database
func PostUser(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	result, err := models.CreateUser(decoder)
	
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
	
}

// UpdateUser to update the user data from databse
func UpdateUser(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	userID := vars["id"]

	decoder := json.NewDecoder(r.Body)

	result, err := models.UpdateUserById(userID, decoder)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Cannot update user into database"))

		log.Fatalln(err.Error())
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, err := json.Marshal(result)
		if err != nil {
			log.Println("error in converting struct to json", err)
		}

		w.Write(response)
	}
}

// DeleteUser to delete user data from database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	result, err := models.DeleteUserById(userID)

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