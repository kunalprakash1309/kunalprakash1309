package main

import (
	"fmt"
	"github.com/kunalprakash1309/romanNumerals"
	"html"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func handle(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	urlPathElements := strings.Split(req.URL.Path, "/")
	fmt.Println(urlPathElements)

	// if request is GET with correct syntax
	fmt.Println(urlPathElements[1])
	if urlPathElements[1] == "roman_number" {
		number, _ := strconv.Atoi(strings.TrimSpace(urlPathElements[2]))

		// if resource is not in the list, send Not Found status
		if number == 0 || number > 10 {
			res.WriteHeader(http.StatusNotFound)
			res.Write([]byte("404- Not Found"))
		} else {
			fmt.Fprintf(res, "%q", html.EscapeString(romanNumerals.Numerals[number]))
		}
	} else {
		// For all other requests
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("404 - Bad Request"))
	}
}

func main() {
	// http package has methods for dealing with requests
	http.HandleFunc("/", handle)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	s := http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
