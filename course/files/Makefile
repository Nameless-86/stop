# Set the default goal
.DEFAULT_GOAL := build

# Format the Go code
fmt:
	go fmt ./...
.PHONY: fmt

# Lint the Go code
lint: fmt
	golint ./...
.PHONY: lint

# Vet the Go code
vet: fmt
	go vet ./...
.PHONY: vet

# Build the Go code
build: vet
	go build main.go
.PHONY: build
