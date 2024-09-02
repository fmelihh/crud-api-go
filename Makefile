build: 
	@go build -o bin/crud cmd/main.go

test:
	@go test -v ./...

run:
	@./bin/crud