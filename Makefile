SOURCES = main.go $(wildcard internal/*/*.go)
BINARY = bin/gor

GOPATH ?= $(HOME)/go
AIR ?= $(GOPATH)/bin/air

default: help

build: $(BINARY) ## build the application binary

$(BINARY): $(SOURCES)
	@go build -o $@

serve: $(AIR) ## run a live reloading development server
	@$< -c .air.toml
.PHONY: serve

$(AIR):
	@go install github.com/air-verse/air@latest

clean: ## clean up generated/temporary files
	@rm -rf tmp
.PHONY: clean

realclean: clean ## clean up All the Things
	@rm -rf $(BINARY)
.PHONY: realclean

help: ## show this help (default)
	@echo "\nSpecify a command. The choices are:\n"
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[0;36m%-18s\033[m %s\n", $$1, $$2}'
	@echo ""
.PHONY: help
