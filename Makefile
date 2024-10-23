all: test lint run

test:
	go test ./... --cover 

lint:
	@golangci-lint run

run:
	go run cmd/main.go

