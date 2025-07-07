build:
	@go build -o bin/social-gopher cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/social-gopher
