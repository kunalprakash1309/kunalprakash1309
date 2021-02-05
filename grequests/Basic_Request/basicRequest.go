package main

import (
	"log"

	"github.com/levigross/grequests"
)


func main() {
	resp, err := grequests.Get("http://httpbin.org/get", nil)
	if err != nil {
		log.Fatalln("Unable to make request", err)
	}

	log.Println(resp.String())
	log.Println(resp.Header)

}