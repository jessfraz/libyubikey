# Setup name variables for the package/tool
NAME := libyubikey
PKG := github.com/jessfraz/$(NAME)

CGO_ENABLED := 1

# Set any default go build tags.
BUILDTAGS :=

include basic.mk

.PHONY: prebuild
prebuild:

.PHONY: generate
generate: image ## Generate the cgo files
	@echo "+ $@"
	@docker run --rm $(DOCKER_FLAGS) \
		--name $(NAME) \
		--disable-content-trust \
		$(REGISTRY)/$(NAME) hack/generate.sh

.PHONY: dtest
dtest: image ## Run the tests in a docker image
	@echo "+ $@"
	@docker run --rm $(DOCKER_FLAGS) \
		--name $(NAME) \
		--disable-content-trust \
		$(REGISTRY)/$(NAME) make build test
