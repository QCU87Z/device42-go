# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=d42-password
BINARY_UNIX=$(BINARY_NAME)_unix
    
all: build
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v cmd/d42-password/main.go
test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v cmd/d42-password/main.go
	./$(BINARY_NAME)
    
    
