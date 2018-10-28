package main

import (
	"fmt"
	"github.com/rs/cors"
	"net/http"
	"os"
	"strings"
)

var DefaultCorsAllowedMethods = []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}
var DefaultCorsAllowedOrigins = []string{}

type corsEnvOptionKeys struct {
	AllowedMethods string
	AllowedOrigins string
}

// EnvOptions defines the name of the environment variables used to
// retrieve the values used to configure CORS
var CorsEnvOptionKeys = corsEnvOptionKeys{
	AllowedMethods: "CORS_ALLOWED_METHODS",
	AllowedOrigins: "CORS_ALLOWED_ORIGINS",
}

// CorsMiddleware returns a middleware that adds cross-origin-
// resource-sharing capabilities to a mux
func CorsMiddleware(server *http.ServeMux) http.Handler {
	allowedMethods := getAllowedMethodsFromEnvironment()
	fmt.Printf("[cors] Allowed Methods: %v\n", allowedMethods)

	allowedOrigins := getAllowedOriginsFromEnvironment()
	fmt.Printf("[cors] Allowed Origins: %v\n", allowedOrigins)

	return cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   allowedMethods,
		AllowCredentials: true,
		Debug:            true,
	}).Handler(server)
}

func getAllowedMethodsFromEnvironment() []string {
	var allowedMethods []string
	_allowedMethods := os.Getenv(CorsEnvOptionKeys.AllowedMethods)
	if _allowedMethods == "" {
		allowedMethods = DefaultCorsAllowedMethods
	} else {
		allowedMethods = strings.Split(_allowedMethods, ",")
	}
	for i := 0; i < len(allowedMethods); i++ {
		allowedMethods[i] = strings.ToUpper(allowedMethods[i])
	}

	return allowedMethods
}

func getAllowedOriginsFromEnvironment() []string {
	var allowedOrigins []string
	_allowedOrigins := os.Getenv(CorsEnvOptionKeys.AllowedOrigins)
	if _allowedOrigins == "" {
		allowedOrigins = DefaultCorsAllowedOrigins
	} else {
		allowedOrigins = strings.Split(_allowedOrigins, ",")
	}

	return allowedOrigins
}
