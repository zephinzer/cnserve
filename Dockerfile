FROM golang:1.11-alpine as development
ARG PROJECT_NAME="app"
RUN apk add --no-cache git curl bash
RUN go get -v github.com/oxequa/realize
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN mkdir -p /.cache/go-build && chmod 777 -R /.cache
WORKDIR /go/src/${PROJECT_NAME}
ENTRYPOINT [ "realize", "start", "--run", "main.go" ]

FROM golang:1.11-alpine as compile
ARG PROJECT_NAME="app"
COPY . /go/src/${PROJECT_NAME}
WORKDIR /go/src/${PROJECT_NAME}
RUN go build \
  && ln -s /go/src/${PROJECT_NAME}/${PROJECT_NAME} /bin/start
ENTRYPOINT [ "start" ]

FROM alpine:3.8 as production
ARG PROJECT_NAME="app"
RUN apk --no-cache update \
  && apk add --no-cache ca-certificates
COPY --from=compile /go/src/${PROJECT_NAME}/${PROJECT_NAME} /bin/start
WORKDIR /
ENTRYPOINT [ "start" ]
