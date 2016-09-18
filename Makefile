# Set an output prefix, which is the local directory if not specified
PREFIX?=$(shell pwd)
BUILDTAGS=

DOCKER_IMAGE := r.j3ss.co/libyubikey

# if this session isn't interactive, then we don't want to allocate a
# TTY, which would fail, but if it is interactive, we do want to attach
# so that the user can send e.g. ^C through.
DOCKER_FLAGS := docker run --rm -i -v $(CURDIR):/go/src/github.com/jfrazelle/libyubikey
INTERACTIVE := $(shell [ -t 0 ] && echo 1 || echo 0)
ifeq ($(INTERACTIVE), 1)
	DOCKER_FLAGS += -t
endif

.PHONY: clean all dbuild generate fmt vet lint build test install dtest

all: clean build fmt lint test vet install

dbuild:
	@echo "+ $@"
	@docker build --rm --force-rm -t $(DOCKER_IMAGE) .

generate: dbuild
	@echo "+ $@"
	@$(DOCKER_FLAGS) --name libyubikey $(DOCKER_IMAGE) hack/generate.sh

dtest: dbuild
	@echo "+ $@"
	@$(DOCKER_FLAGS) --name libyubikey $(DOCKER_IMAGE) make build fmt test

build:
	@echo "+ $@"
	@go build -tags "$(BUILDTAGS) cgo" .

fmt:
	@echo "+ $@"
	@gofmt -s -l . | grep -v vendor | tee /dev/stderr

lint:
	@echo "+ $@"
	@golint ./... | grep -v vendor | tee /dev/stderr

test:
	@echo "+ $@"
	@go test -v -tags "$(BUILDTAGS) cgo" $(shell go list ./... | grep -v vendor)

vet:
	@echo "+ $@"
	@go vet $(shell go list ./... | grep -v vendor)

install:
	@echo "+ $@"
	@go install .
