.PHONY: all clear build run-linter run-tests

all: run-linter build

PKG_LINTERS = ./...
LINTER_CFG = ./configs/linter.yml

run-linter:
	$(GOPATH)/bin/golangci-lint run $(PKG_LINTERS) --config=$(LINTER_CFG)
	go fmt $(PKG_LINTERS)

build:
	go build -o main.out cmd/main/main.go

