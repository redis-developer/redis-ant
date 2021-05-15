.PHONY: help setup run-server run-client build install run clean

help: ## Display help screen
		@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

setup: ## Setup dev environment
		go get ./...

run-server: ## Run server
		go run ./cmd/redis-ant-server
run-client: ## Run client
		go run ./cmd/redis-ant-client

run-redis: ## Run RedisJSON
	  docker run -p 6379:6379 --name redis-redisjson redislabs/rejson:latest

install: ## install commands
		go install ./cmd/redis-ant-server
		go install ./cmd/redis-ant-client

build: ## Build commands
		go build -o ./bin/redis-ant-server ./cmd/redis-ant-server
		go build -o ./bin/redis-ant-client ./cmd/redis-ant-client

clean: ## Remove all the build files
		rm -rf ./bin/redis-ant-server
		rm -rf ./bin/redis-ant-client