.PHONY: lint tool test help

export GOBIN ?= $(PWD)/bin

default: help

lint: ## runs the golint, checks for lint violations
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.49.0
	@echo Running golangci-lint
	$(GOBIN)/golangci-lint run ./...

tool: ## runs specified go tool
	go vet ./...; true
	gofmt -w .

test: ## runs the tests of this project
	go test ./...

help:
	@echo 'USAGE:'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[1;36m\%-10s\033[0m %s\n", $$1, $$2}'
	@echo
