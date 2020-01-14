.PHONY: build get-deps update-deps help default

default: help

BIN_NAME="mc-aws"

help:
	@echo 'Directives for MC-AWS:'
	@echo
	@echo 'Usage:'
	@echo '    make build           Compile the project.'
	@echo '    make get-deps        runs mod download to download all the dependencies'
	@echo '    make update-deps     runs mod tidy to update dependencies'
	@echo '    make test            Run tests on a compiled project.'
	@echo

build:
	go build -o bin/${BIN_NAME}
	
get-deps:
	@echo "Downloading dependencies..."
	go mod download

update-deps:
	@echo "Updating dependencies..."
	go mod tidy
