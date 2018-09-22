PKGS ?= $(shell go list ./...)
TESTS ?= ".*"
COVERS ?= "c.out"
FMT_FILES ?= $$(find . -name '*.go' |grep -v vendor)

fmt:
	@go fmt $(GOFMT_FILES)

build:
	@go build $(PKGS)

install: build
	@go install

test:
	@go test $(PKGS)

testacc:
	@TF_ACC=1 go test $(PKGS) -v -coverprofile=$(COVERS) -run ^$(TESTS)$
