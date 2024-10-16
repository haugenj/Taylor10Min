HELM_OPTS ?=

build: ## Build the Tay binary
	go build -o ./bin/tay ./cmd/tay

clean: ## Delete the binary
	rm ./bin/tay

help: ## Display help
	@awk 'BEGIN {FS = ":.*##"; printf "Usage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

run: build ## Build and run
	./bin/tay

verify: ## Verify code. Includes dependencies, linting, formatting, etc
	go mod tidy
	go mod download
	go vet ./...
	go fmt ./...
	golangci-lint run --timeout 5m


.PHONY: build clean help run verify
