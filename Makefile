# Make variables
BIN_OUTPUT_FOLDER := bin
BIN_MAIN_FILE := minst

# Commands
default: 
	@go run cmd/minst/*.go

build: 
	@go build -o $(BIN_OUTPUT_FOLDER)/$(BIN_MAIN_FILE) -v cmd/minst/*.go

tests: 
	@go test ./... -cover ./... -v