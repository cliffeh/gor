SOURCES = main.go $(wildcard internal/*/*.go)
BINARY = bin/gor

help: ## show this help (default)
	@echo "\nSpecify a command. The choices are:\n"
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[0;36m%-18s\033[m %s\n", $$1, $$2}'
	@echo ""
.PHONY: help

build: $(BINARY) ## build the application binary

$(BINARY): $(SOURCES)
	@go build -o $@

serve: $(BINARY) ## run a development server
	@$< -bind 127.0.0.1:8080
.PHONY: serve
