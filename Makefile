test:
	@go test

example:
	@go run cmd/example.go

.PHONY: example test
