GOBIN=$(shell pwd)/bin
GOFILES=$(wildcard *.go)
GONAME=app
WEBGONAME=web
CLIGONAME=cli
PID=/tmp/go-$(GONAME).pid
export GO111MODULE=on


build-docker:
	@echo "Building app to ./bin/$(GONAME)"
	@CGO_ENABLED=0 GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build -v -o bin/$(GONAME) ./$(GOFILES)
