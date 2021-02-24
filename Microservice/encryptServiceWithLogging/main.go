package main

import (
	"log"
	"net/http"
	"os"

	kitlog "github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/kunalprakash1309/Microservice/encryptServiceWithLogging/helpers"
)

func main() {
	logger := kitlog.NewLogfmtLogger(os.Stderr)
	var svc helpers.EncryptService
	svc = helpers.EncryptServiceInstance{}
	svc = helpers.LoggingMiddleware{Logger: logger, Next: svc}

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

	log.Fatalln(http.ListenAndServe(":8000", nil))
}