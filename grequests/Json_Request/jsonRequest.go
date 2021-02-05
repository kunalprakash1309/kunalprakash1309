package main

import (
	"log"

	"github.com/levigross/grequests")


func main() {
	resp, err := grequests.Get("http://httpbin.org/get", nil)

	if err != nil {
		log.Fatalln("Unable to make requests : ", err)
	}

	var returnData map[string]interface{}
	// json populate the struct or map with json provided from response
	resp.JSON(&returnData)
	log.Println(returnData)
}