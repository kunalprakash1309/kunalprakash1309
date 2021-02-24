package main

import (
	// "log"
	//"context"
	"net/http"
	"os"

	//"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

// <<<<<<<<<------- Transports -------->>>>>>>>>>

// Now we need to expose our service to the outside world
// Our organisation may use any JSON, RPC, Thrift
// Let use JSON

func main() {

	logger := log.NewLogfmtLogger(os.Stderr)

	var svc StringService

	svc = stringService{}
	svc = loggingMiddleware{logger, svc}

	// var uppercase endpoint.Endpoint
	// uppercase = makeUppercaseEndpoint(svc)
	// uppercase = loggingMiddleware(log.With(logger, "method", "uppercase"))(uppercase)

	// var count endpoint.Endpoint
	// count = makeCountEndpoint(svc)
	// count = loggingMiddleware(log.With(logger, "method", "count"))(count)

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
	http.ListenAndServe(":8080", nil)
}

// type Middleware func(endpoint.Endpoint) endpoint.Endpoint

// func loggingMiddleware(logger log.Logger) Middleware {
// 	return func(next endpoint.Endpoint) endpoint.Endpoint {
// 		return func(ctx context.Context, request interface{}) (interface{}, error) {
// 			logger.Log("msg", "calling endpoint")
// 			defer logger.Log("msg", "called endpoint")
// 			return next(ctx, request)
// 		}
// 	}
// }