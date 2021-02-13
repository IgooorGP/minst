GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
BINARY_OUTPUT_NAME=minst

# Commands
default: 
	$(GORUN) cmd/minst/root.go

run: 
	$(GORUN) cmd/minst/root.go

build: 
	$(GOBUILD) -o $(BINARY_OUTPUT_NAME) -v cmd/minst/root.go

tests: 
	$(GOTEST) ./... -cover ./... -v