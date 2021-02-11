package main

// Our service starts with our business logic.

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// <<<<<<<<<------- Buisness Logic -------->>>>>>>>>>

// In Go kit, we model a service as an interface.

// StringService provides operations on strings.
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}


// <<<<<<<<<------- Implementation -------->>>>>>>>>>

// StringService will have an implementation

type stringService struct{}

func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
	return len(s)
}

// ErrEmpty is returned when input string is empty
var ErrEmpty = errors.New("Empty string")

// <<<<<<<<<------- Requests and Response -------->>>>>>>>>>

// In go-kit, the primary messaging pattern is RPC.
// So, every method in our interface will be modeled as a remote procedure call
// For each method, we define request and response structs

type uppercaseRequest struct {
	S string `json:"s"`
}

type uppercaseResponse struct {
	V string `json:"V"`
	Err string `json:"err,omitempty"`
}

type countRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	V int `json:"V"`
}

// <<<<<<<<<------- Endpoints -------->>>>>>>>>>

// Definition of endpoint
// type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)
// It represents single RPC. It means a single method in our service interface.
// We’ll write simple adapters to convert each of our service’s methods into an endpoint
// Each adapters takes a StringService and returns an endpoint that corresponds to one of the methods

func makeUppercaseEndpoint(svc StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(uppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return uppercaseResponse{v, err.Error()}, nil
		}
		return uppercaseResponse{v, ""}, nil
	}
}

func makeCountEndpoint(svc StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req := request.(countRequest)
		v := svc.Count(req.S)
		return countResponse{v}, nil
	}
}

// <<<<<<<<<------- Transports -------->>>>>>>>>>

// Now we need to expose our service to the outside world
// Our organisation may use any JSON, RPC, Thrift
// Let use JSON

func main() {

	svc := stringService{}
	uppercaseHandler := httptransport.NewServer(
		makeUppercaseEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
	)

	countHadler := httptransport.NewServer(
		makeCountEndpoint(svc),
		decodeCountRequest,
		encodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHadler)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func decodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request uppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request countRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}