
# Set the default goal
.DEFAULT_GOAL := build
VERSION := $(shell git describe --tags)
LDFLAGS=-ldflags "-s -w -X=main.version=$(VERSION)"

.PHONY: build
build :
	@echo "Building ...."
	CGO_ENABLED=0 go build $(LDFLAGS) -o ./bin/demo ./server/server.go

test:
	echo "Testing echo"

test-integration:
	echo "Integration Testing echo"

linter:
	 golangci-lint run

gosec:
	gosec  ./...

build-image:
	docker build -t krol/demo:demo .

semgrep:
	semgrep --config=auto .
