test:
	@go test -v ./...

coverage:
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o ./tests/coverage.html
	@rm coverage.out

build-so:
	go build -o dist/library.so -buildmode=c-shared main.go
