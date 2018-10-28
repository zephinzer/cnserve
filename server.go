package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"os"
	"path"
)

// DefaultServerStaticPath ...
const DefaultServerStaticPath = "/static"

type serverEnvOptionKeys struct {
	StaticPath string
}

// ServerEnvOptionKeys ...
var ServerEnvOptionKeys = serverEnvOptionKeys{
	StaticPath: "SERVER_STATIC_PATH",
}

func Server() *http.ServeMux {
	server := http.NewServeMux()
	staticPath := getStaticPathFromEnvironment()

	server.Handle("/metrics", promhttp.Handler())
	server.Handle("/", prometheus.InstrumentHandler(
		"/", http.FileServer(http.Dir(staticPath))))
	fmt.Println("[server] Static files path: ", staticPath)
	return server
}

func getStaticPathFromEnvironment() string {
	var staticPath string
	_staticPath := os.Getenv(ServerEnvOptionKeys.StaticPath)
	curentDirectory, err := os.Getwd()
	if err != nil {
		panic(err)
	} else if len(_staticPath) == 0 {
		staticPath = path.Join(curentDirectory, DefaultServerStaticPath)
	} else {
		staticPath = path.Join(curentDirectory, _staticPath)
	}
	return staticPath
}
