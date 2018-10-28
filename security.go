package main

import (
	"fmt"
	"github.com/unrolled/secure"
	"net/http"
	"os"
	"strings"
)

var DefaultSecurityAllowedHosts = []string{}

type securityEnvOptionKeys struct {
	AllowedHosts string
}

var SecurityEnvOptionKeys = securityEnvOptionKeys{
	AllowedHosts: "SECURITY_ALLOWED_HOSTS",
}

func SecurityMiddleware(handler http.Handler) http.Handler {
	allowedHosts := getAllowedHostsFromEnvironment()
	fmt.Printf("[security] Allowed Hosts: %v\n", allowedHosts)

	development := getDevelopmentModeFromEnvironment()
	fmt.Printf("[security} Development Mode: %v\n", development)

	return secure.New(secure.Options{
		AllowedHosts:            allowedHosts,
		IsDevelopment:           development,
		HostsProxyHeaders:       []string{"X-Forwarded-Host"},
		STSSeconds:              315360000,
		STSIncludeSubdomains:    true,
		STSPreload:              true,
		FrameDeny:               true,
		CustomFrameOptionsValue: "SAMEORIGIN",
		ContentTypeNosniff:      true,
		BrowserXssFilter:        true,
		ContentSecurityPolicy:   "default-src 'self'",
		ReferrerPolicy:          "same-origin",
	}).Handler(handler)
}

func getAllowedHostsFromEnvironment() []string {
	_allowedHosts := os.Getenv(SecurityEnvOptionKeys.AllowedHosts)
	var allowedHosts []string
	if len(_allowedHosts) == 0 {
		allowedHosts = DefaultSecurityAllowedHosts
	} else {
		allowedHosts = strings.Split(_allowedHosts, ",")
	}
	return allowedHosts
}
