.PHONY: all

all: dep test bin

bin:
	@go build -v -o bitmap $(PWD)/cmd

dep:
	@go mod vendor
	@go mod tidy

test:
	./test.sh

fake:
	@go generate ./...

help:
	@echo "Usage:"
	@echo "     bin  ................ build the binary (goes to ./bitmap)"
	@echo "     dep  ................ update dependencies"
	@echo "     test ................ run all tests (requires 'ginkgo')"
	@echo "     fake ................ regenerate interface mocks for testing"
