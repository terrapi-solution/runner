SHELL := bash
NAME := terrapi-worker
BIN := bin

GOBUILD ?= CGO_ENABLED=0 go build
PACKAGES ?= $(shell go list ./...)
SOURCES ?= $(shell find . -name "*.go" -type f -not -path ./.devenv/\* -not -path ./.direnv/\*)
GENERATE ?= $(PACKAGES)
TAGS ?= netgo

.PHONY: generate
generate:
	go generate $(GENERATE)

.PHONY: vet
vet:
	go vet $(PACKAGES)

.PHONY: staticcheck
staticcheck: $(STATICCHECK)
	$(STATICCHECK) -tags '$(TAGS)' $(PACKAGES)

.PHONY: build
build: $(BIN)/$(NAME)

$(BIN)/$(NAME): $(SOURCES)
	$(GOBUILD) -v -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $@ 

$(BIN)/$(NAME)-debug: $(SOURCES)
	$(GOBUILD) -v -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -gcflags '$(GCFLAGS)' -o $@

.PHONY: test
test:
	go test -coverprofile coverage.out $(PACKAGES)