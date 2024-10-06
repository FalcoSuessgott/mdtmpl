default: help

.PHONY: help
help: ## list makefile targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

PHONY: fmt
fmt: ## format go files
	gofumpt -w .
	gci write .

PHONY: test
test: ## display test coverage
	gotestsum -- -v -race -coverprofile="coverage.out" -covermode=atomic ./...

.PHONY: lint
lint: ## lint go files
	golangci-lint run -c .golang-ci.yml
