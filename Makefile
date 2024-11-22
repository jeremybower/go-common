help: ## Show this help message.
	@echo
	@echo 'usage: make [target]'
	@echo
	@echo 'targets:'
	@echo
	@egrep '^(.+)\:\ ##\ (.+)' ${MAKEFILE_LIST} | column -t -c 2 -s ':#'
	@echo
.PHONY: help

init: ## Initialize the project.
	touch docker-compose.local.yml

test: ## Test the project.
	@$(eval include .env)
	@$(eval export)
	@mkdir -p coverage
	@go test \
		-cover \
		-covermode=atomic \
		-coverprofile coverage/coverage.out \
		-count=1 \
		-failfast \
		./...
	@go tool cover \
		-html=coverage/coverage.out \
		-o coverage/coverage.html

tidy: ## Tidy the go modules.
	@$(eval include .env)
	@$(eval export)
	@go mod tidy
.PHONY: tidy

update: ## Update the go modules.
	@$(eval include .env)
	@$(eval export)
	@go get -u ./...
	@go mod tidy
.PHONY: update
