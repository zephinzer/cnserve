# Cloud Native Server
A barebones static file server for use in a cloud native environment written in Go.

## Overview
This server includes:
- healthcheck endpoints
- openmetrics support
- opentracing support
- access logs streaming

## Todos
- [x] Serve static files
- [x] Static files directory configuration
- [x] Cross Origin Resource Sharing implementation
- [x] Cross Origin Resource Sharing configuration
- [x] Content Security Policy implementation
- [ ] Content Security Policy configuration
- [x] Prometheus metrics implementation
- [ ] Prometheus metrics endpoint configuration
- [ ] Prometheus metrics endpoint BasicAuth protection
- [ ] Prometheus metrics push gateway implementation
- [ ] Prometheus metrics push gateway configuration
- [x] Zipkin distributed tracing implementation
- [ ] Zipkin distributed tracing configuration
- [ ] Healthcheck implementation
- [ ] Healthcheck configuration

## Usage

### Configuration

| Environment Variable | Description | Example Value |
| --- | --- | --- |
| `CORS_ALLOWED_METHODS` | Comma separated list of methods for Cross-Origin-Resource-Sharing requests | `GET,POST,PUT` |
| `CORS_ALLOWED_ORIGINS` | Comma separated list of origins for Cross-Origin-Resource-Sharing requests | `http://website.com,http://this.goodservice.com` |
| `DEV` | If truthy, runs the server in development | `1` , `true` |
| `PORT` | Port which the server should listen on | `8080` |
| `SERVER_STATIC_PATH` | Path to the directory containing files to serve | `/static` |
| `SECURITY_ALLOWED_HOSTS` | Comma separated list of hostnames which the server should respond to | `hello.com,www.hello.com,www1.hello.com` |

## Development Commands
### Getting Started
Run `make start` to start the application with live-reloading. Live reloading functionality is provided by Realize.

See [the GitHub page for `realize`](https://github.com/oxequa/realize) for more information.

### Getting a Build
Run `make compile` to compile a distributable binary.

### Adding a Dependency
Run `make dep.add DEP=github.com/username/gopkg` to add a dependency.

See [the GitHub page for `dep`](https://github.com/golang/dep) for more information.

## Deployment Commands
Run `make publish` to publish the image