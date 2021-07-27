default: build

build:
	@CGO_ENABLED=1 go build -o bin/rest-api

dev:
	@go run main.go
