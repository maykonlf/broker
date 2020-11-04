SHELL := /bin/bash

SRC = $(shell find . -type f -name '*.go' -not -iname '*.pb.*' -not -path '**/mocks/*')

lint:
	@golangci-lint run ./...

test:
	@go test -cover ./...

cover:
	@go test -coverpkg=./pkg/... -coverprofile=cover.out ./pkg/... > /dev/null
	@sed -i '/mock.go/d' cover.out
	@sed -i '\#/mocks#d' cover.out
	@go tool cover -func cover.out

clean:
	@go clean

fmt:
	@gofmt -s -l -w $(SRC)

goimports:
	@goimports -w -local github.com/golangci/golangci-lint $(SRC)

format: fmt goimports

.PHONY: all test