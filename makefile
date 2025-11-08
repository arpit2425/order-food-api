# Project Makefile

# Use bash for nicer command handling
SHELL := /bin/bash

.DEFAULT_GOAL := help



test: 
	@echo "🧪 Running all Go tests..."
	go test -v ./...

tidy: 
	go mod tidy


run:
	@echo "🚀 Starting app..."
	go run ./cmd/server/main.go

