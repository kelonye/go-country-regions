test: data.go
	@go test

example: data.go
	@go run cmd/example.go

data.go: regions.json
	@go-bindata -o $@ -pkg regions regions.json

.PHONY: example test
