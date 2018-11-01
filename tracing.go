package main

import (
	zipkin "github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"
	logreporter "github.com/openzipkin/zipkin-go/reporter/log"
	"log"

	"net/http"
	"os"
)

// TracingMiddleware ...
func TracingMiddleware(handler http.Handler) http.Handler {
	reporter := logreporter.NewReporter(log.New(os.Stderr, "", log.LstdFlags))
	defer reporter.Close()
	endpoint, endpointErr := zipkin.NewEndpoint("cnserve", "localhost:9411")
	if endpointErr != nil {
		log.Fatalf("Unable to create local endpoint: %v\n", endpointErr)
	}
	tracer, tracerErr := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
	if tracerErr != nil {
		log.Fatalf("Unable to create tracer: %v\n", tracerErr)
	}
	zipkinMiddleware := zipkinhttp.NewServerMiddleware(tracer, zipkinhttp.TagResponseSize(true))
	return zipkinMiddleware(handler)
}
