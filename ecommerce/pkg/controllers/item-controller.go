package controllers


import (
	"encoding/json"
	"log"
	"html/template"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/kunalprakash1309/ecommerce/pkg/models"
	"github.com/kunalprakash1309/ecommerce/pkg/config"
)

var tpl *template.Template

// GetItem to get the items by provided id from item colleciton
func GetItem(w http.ResponseWriter, r *http.Request){

	tpl = config.Tpl
	// get all parameters in form of map
	vars := mux.Vars(r)
	userID := vars["id"]

	item, err := models.GetItemById(userID)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No Such Item"))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response, _ := json.Marshal(item)
		w.Write(response)
	}
}

// PostItem to post the item data into items collection
func PostItem(w http.ResponseWriter, r *http.Request) {


	decoder := json.NewDecoder(r.Body)
	result, err := models.CreateItem(decoder)
	
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Cannot insert user into database"))

		log.Fatalln(err.Error())
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, err := json.Marshal(result)
		if err != nil {
			log.Println("error in converting struct to json", err)
		}

		w.Write(response)
	}

}

// DeleteItem to delete the item by given ID 
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	result, err := models.DeleteItemById(userID)

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