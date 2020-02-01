.PHONY: all fmt lint test build clean help

all: fmt dep lint test build

fmt: ## gofmt all project
	@gofmt -l -s -w .

dep: ## make dependencies
	@go mod vendor

lint: ## Lint the files
	@golangci-lint run

test: ## Run unittests
	@go test -short ./... -coverprofile=coverage.txt

build: ## Build the binary file
	@go build

clean: ## Remove previous build
	@rm -f gopostgres-time

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
