PYTHON     = $(firstword $(shell which python3.9 python3.8 python3.7 python3))
PYTEST      ?= $(PYTHON) -m pytest
PYTEST_ARGS ?= -vv --disable-pytest-warnings

OPENAPI ?= docs/api.yaml

APP_ARGS ?=

# NOTE: use Makefile.local for customization
-include Makefile.local

.PHONY: all
all: tests

TARGETS = \
	codegen \
	test \
	utest \
	build \
	run \
	clean
DOCKER_TARGETS = $(foreach target,$(TARGETS),docker-$(target))
.PHONY: $(TARGETS) $(DOCKER_TARGETS) docker docker-update

codegen:
	@oapi-codegen -package=rest -generate=types,server -include-tags=api $(OPENAPI) > internal/adapters/rest/handlers.gen.go
	@oapi-codegen -package=html -generate=types,server,skip-prune -include-tags=ui $(OPENAPI) > internal/adapters/html/handlers.gen.go
	@oapi-codegen -package=errors -generate=types,skip-prune docs/errors.yaml > pkg/errors/errors.gen.go
	@go generate ./...

test: utest build
	@PYTHONPATH=../.. TESTSUITE_ALLOW_ROOT=1 $(PYTEST) $(PYTEST_ARGS) tests

utest:
	@go test ./...

run: build
	@./main $(APP_ARGS)

build:
	@go build cmd/server/main.go

clean:
	@rm main

docker:
	@docker compose run --rm -it app bash

docker-update:
	@docker build --tag stewkk/testsuite-golang .
	@docker push stewkk/testsuite-golang
	@docker compose pull

$(DOCKER_TARGETS): docker-%:
	@docker compose run --service-ports --rm app $(MAKE) $*

.PHONY: docker-clean-data
docker-clean-data:
	@docker compose down -v
