GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
BINARY_OUTPUT_NAME=terminstall

# Commands
default: 
	$(GORUN) cmd/terminstall/main.go

run: 
	$(GORUN) cmd/terminstall/main.go

build: 
	$(GOBUILD) -o $(BINARY_OUTPUT_NAME) -v cmd/terminstall/main.go

test: 
	$(GOTEST) ./... -cover ./... -v