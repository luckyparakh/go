.DEFAULT_GOAL := run
.PHONY: fmt build vet clean

fmt: 
	go fmt ./...
build: fmt
	go build ./...
run: build
	go run main.go

clean:
	go clean