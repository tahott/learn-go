.PHONY: tool help

default: help

tool: ## runs specified go tool
	go vet ./...; true
	gofmt -w .

help:
	@echo 'USAGE:'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[1;36m\%-10s\033[0m %s\n", $$1, $$2}'
	@echo
