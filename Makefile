.PHONY: help run

HELPER_CMD = $(filter-out $(firstword $(MAKECMDGOALS)), $(MAKECMDGOALS))

# App Variables
PORT = 8000
CLOUDFLAR
# Build Variables
REPOSITORY = localhost:5000
BUILD_NAME = prometheus-domain-expiry-exporter
RELEASE = stable

help: ## Help. 
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# Development

build: ## builds container
	@docker rmi ${REPOSITORY}/${BUILD_NAME}:${RELEASE} || true
	@docker build . -t ${REPOSITORY}/${BUILD_NAME}:${RELEASE}

run: ## runs example
	@PORT=${PORT} go run main.go

test: ## runs tests
	@go mod tidy
	@go test
