# export env variables
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

.PHONY: default run

# vars
APP_NAME=softutils_golanf

default: run

install:
	@echo "Installing dependencies"
	@go install github.com/cosmtrek/air@latest
	@go mod tidy

run:
	@echo "Running the application"
	@go mod tidy 
	@go run main.go

build:
	@echo "Building the application"

clean:
	@echo "Cleaning the application"
	@rm -rf bin
