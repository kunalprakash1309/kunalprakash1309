package main

import (
	"log"
	"net/http"
	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/kunalprakash1309/Microservice/encryptService/helpers"
)

func main() {
	svc := helpers.EncryptServiceInstance{}

	encrypthandler := httptransport.NewServer(
		helpers.MakeEncryptEndpoint(svc),
		helpers.DecodeEncryptRequest,
		helpers.EncodeResponse,
	)

	decryptHandler := httptransport.NewServer(
		helpers.MakeDecryptEndpoint(svc),
		helpers.DecodeDecryptRequest,
		helpers.EncodeResponse,
	)

	http.Handle("/encrypt", encrypthandler)
	http.Handle("/decrypt", decryptHandler)

	log.Fatalln(http.ListenAndServe(":8080", nil))
}