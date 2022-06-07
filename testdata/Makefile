BINPATH:=$(abspath .)/bin

.PHONY: goa_install
goa_install:
	mkdir -p ${BINPATH}
	go get goa.design/goa/v3/cmd/goa
	go build -o ${BINPATH}/goa goa.design/goa/v3/cmd/goa

.PHONY: goa_gen
goa_gen: goa_install
	go mod tidy
	@goa version
	${BINPATH}/goa gen calc/design

.PHONY: test
test: goa_gen
	go test ./...

