GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
BINARY_OUTPUT_NAME=minst

# Commands
default: 
	$(GORUN) cmd/minst/main.go

run: 
	$(GORUN) cmd/minst/main.go

build: 
	$(GOBUILD) -o $(BINARY_OUTPUT_NAME) -v cmd/minst/main.go

tests: 
	$(GOTEST) ./... -cover ./... -v