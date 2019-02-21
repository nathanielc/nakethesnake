install:
	go install github.com/nathanielc/nakethesnake
.PHONY: install

run: install
	PORT=9000 nakethesnake server
.PHONY: run

test:
	go test ./...
.PHONY: test

fmt:
	@echo ">> Running Gofmt.."
	gofmt -l -s -w .
.PHONY: fmt
