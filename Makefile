run-tests:
	@echo "Running tests..."
	@go test ./...

install:
	@echo "Installing dependencies..."
	@go mod tidy

run:
	go run ./...