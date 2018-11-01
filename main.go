package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// DefaultPort ...
const DefaultPort = 8080

type mainEnvOptionKeys struct {
	Development string
	Port        string
}

// MainEnvOptionKeys ...
var MainEnvOptionKeys = mainEnvOptionKeys{
	Development: "DEV",
	Port:        "PORT",
}

func main() {
	dev := getDevelopmentModeFromEnvironment()
	fmt.Printf("[main] Development Mode: %v\n", dev)

	server := Server()
	handler := CorsMiddleware(server)
	handler = SecurityMiddleware(handler)
	handler = TracingMiddleware(handler)

	port := getPortFromEnvironment()
	fmt.Printf("[main] Listening on port: %d\n", port)
	serverError := http.ListenAndServe(":"+strconv.Itoa(port), handler)
	if serverError != nil {
		fmt.Println(serverError)
	}
}

func getDevelopmentModeFromEnvironment() bool {
	_development := os.Getenv(MainEnvOptionKeys.Development)
	development := false
	if len(_development) == 0 {
		development = true
	} else {
		developmentText := strings.ToLower(_development)
		if developmentText == "1" || developmentText == "yes" || developmentText == "true" {
			development = true
		}
	}
	return development
}

func getPortFromEnvironment() int {
	var port int
	_port := os.Getenv(MainEnvOptionKeys.Port)
	if len(_port) == 0 {
		port = DefaultPort
	} else {
		portNumber, err := strconv.Atoi(_port)
		if err != nil {
			port = DefaultPort
		} else {
			port = portNumber
		}
	}
	return port
}
