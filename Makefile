default: build

build:
	@CGO_ENABLED=0 go build -o bin/rest-api 

dev:
	@go run main.go
