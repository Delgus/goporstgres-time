PKG := "github.com/delgus/gopostgres-time"

.PHONY: all fmt dep lint build clean help

all: fmt dep lint build

fmt: ## gofmt all project
	@gofmt -l -s -w .

dep: ## make dependencies
	@go mod vendor

lint: ## Lint the files
	@golangci-lint run

build: ## Build the binary file
	@go build -a -o ex1 -v $(PKG)/cmd/ex1

clean: ## Remove previous build
	@rm -f ex1

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
