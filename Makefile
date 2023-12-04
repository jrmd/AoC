build:
	@go build main.go

run:
	@go run main.go run --all

test:
	go test -v ./...
