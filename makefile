SHELL := /bin/bash
PREFIX?=$(shell pwd)
GO := go

GO_LDFLAGS_STATIC=-ldflags "-extldflags -static"

build: ## Builds a static executable
	@echo "+ $@"
	$(GO) build -tags "static_build" ${GO_LDFLAGS_STATIC} -o server ./cmd

run: ## Run server
	@echo "+ $@"
	$(GO) run cmd/main.go
	
get-all: ## Run client to get all employees
	@echo "+ $@"
	$(GO) run client.go 

add: ## Run client to add new employee
	@echo "+ $@"
	$(GO) run client.go -add -g $(gender)

clean: ## Cleanup any build binaries, packages or artifacts
	@echo "+ $@"
	$(RM) server