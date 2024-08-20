ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif

install:
	go get ./...

dev:
	$(GOPATH)/bin/air -c .air.toml

test:
	go test -cover -v ./...

test-json:
	go test -cover -json -v ./...

coverage:
	go test -coverprofile=tmp/coverage.out ./...
	go tool cover -html=ctmp/overage.out

lint:
	test -z $(gofmt -l .)

build:
	go build -o bin/main cmd/api/main.go

run:
	go run cmd/api/main.go



.PHONY: install dev test test-json coverage lint build run