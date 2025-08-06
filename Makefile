SOURCES = main.go $(wildcard internal/*/*.go)
BINARY = bin/gor

GOPATH ?= $(HOME)/go
AIR ?= $(GOPATH)/bin/air

COVDIRS = cov/unit cov/int

default: help

$(COVDIRS):
	@mkdir -p $@

build: $(BINARY) ## build the application binary

$(BINARY): $(SOURCES)
	@go build -cover -o $@

serve: $(AIR) ## run a live reloading development server
	@$< -c .air.toml
.PHONY: serve

$(AIR):
	@go install github.com/air-verse/air@latest

test: ## run unit tests
	@go test ./...
.PHONY: test

coverage: $(BINARY) cov/unit cov/int ## generate a test coverage report
	go test -cover ./internal/... -args -test.gocoverdir="$(PWD)/cov/unit"
	GOCOVERDIR="$(PWD)/cov/int" go run test/integration.go $(BINARY)
	go tool covdata textfmt -i=./cov/unit,./cov/int -o cov/profile
	go tool cover -func cov/profile
.PHONY: coverage

container: ## build a container image
	@docker build -t gor .
.PHONY: container

container-serve: container ## run the containerized app
	@docker run --rm -p 8080:8080 gor
.PHONY: container-serve

clean: ## clean up generated/temporary files
	@rm -rf cov tmp
.PHONY: clean

realclean: clean ## clean up All the Things
	@rm -rf bin
.PHONY: realclean

help: ## show this help (default)
	@echo "\nSpecify a command. The choices are:\n"
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[0;36m%-18s\033[m %s\n", $$1, $$2}'
	@echo ""
.PHONY: help
