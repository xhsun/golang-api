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

clean: ## Cleanup any build binaries, packages or artifacts
	@echo "+ $@"
	$(RM) server