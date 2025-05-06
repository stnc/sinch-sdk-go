
GOCMD=go

GOTEST=go test

BUILD_DIR=build
BINARY_DIR=$(BUILD_DIR)/bin
CODE_COVERAGE=code-coverage

all: build

${BINARY_DIR}:
	mkdir -p $(
		
	)

build: ${BINARY_DIR} ## Compile the code, build Executable File ## mkdir -p $(BINARY_DIR)
	$(GOCMD) build -o $(BINARY_DIR) -v ./cmd

run: ## Start application
	$(GOCMD) run ./cmd