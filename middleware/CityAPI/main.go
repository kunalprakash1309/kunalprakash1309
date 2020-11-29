package main

import (
	"encoding/json"
	"net/http"
	"log"
	"strconv"
	"time"
)

type city struct {
	Name string
	Area uint64
}

// Middleware to check content type as a json
func filterContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Currently in the check content type middleware")

		if r.Header.Get("Content-type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("415- Unsupported Media Type. Please send JSON"))
		}
	})
}

// Middleware to add server timestamp for response cookie
func setServerTimeCookie(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// handler.ServeHTTP(w, r)
		log.Println("currently in set server time middleware")

		//setting  cookie to each and every response
		cookie := http.Cookie{
			Name: "Server-Time(UTC)",
			Value: strconv.FormatInt(time.Now().Unix(), 10),
		}

		http.SetCookie(w, &cookie)
	})
}

func mainLogic(w http.ResponseWriter, r *http.Request) {
	// Check if method is POST
	if r.Method == "POST"{

		w.Write([]byte("Hello"))
		var tempCity city

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&tempCity)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()

		log.Printf("Got %s city with area of %d sq miles! \n", tempCity.Name, tempCity.Area)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("201-created"))

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405-Method not allowed"))
	}
}

func main() {
	mainLoginHandler := http.HandlerFunc(mainLogic)
	
	http.Handle("/city", filterContentType(setServerTimeCookie(mainLoginHandler)))
	http.ListenAndServe(":8080", nil)
}