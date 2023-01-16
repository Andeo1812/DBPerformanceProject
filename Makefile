.PHONY: all clear build run-linter run-tests

all: run-linter build

PKG = ./...
LINTER_CFG = ./configs/linter.yml

run-linter:
	$(GOPATH)/bin/golangci-lint run $(PKG) --config=$(LINTER_CFG)
	go fmt $(PKG)

build:
	go build -o main.out cmd/main/main.go

# easyjson
generate:
	go generate ${PKG}
