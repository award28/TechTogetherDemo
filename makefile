# Go parameters
GOBUILD=go build 
GOCLEAN=go clean
GOTEST=go test
BINARY=app
APP=./cmd/app/$(BINARY).go

.PHONY: clean test

all: build run

build: 
	$(GOBUILD) $(APP)

test: 
	$(GOTEST) -tags unit -v ./internal/...

clean: 
	rm -f ./$(BINARY)

run:	
	./$(BINARY)
