.PHONY: goa_gen
goa_gen:
	go mod tidy
	@go run goa.design/goa/v3/cmd/goa version
	go run goa.design/goa/v3/cmd/goa gen calc/design

.PHONY: test
test: goa_gen
	go test ./...

