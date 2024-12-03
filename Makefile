build:
	@go build -o bin/EcomGo cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/EcomGo
	