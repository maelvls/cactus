.PHONY: all

all: dep test build

bin:
	@go build -v -o bitmap $(PWD)/cmd

dep:
	@go mod vendor
	@go mod tidy

test:
	@ginkgo -mod vendor -r
