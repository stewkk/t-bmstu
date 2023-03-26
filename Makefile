PYTHON     = $(firstword $(shell which python3.9 python3.8 python3.7 python3))
PYTEST      ?= $(PYTHON) -m pytest
PYTEST_ARGS ?= -vv --disable-pytest-warnings

OPENAPI ?= api/openapi.yaml

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
.PHONY: $(TARGETS) $(DOCKER_TARGETS)

codegen:
	@mkdir -p internal/api/
	@oapi-codegen -package=models -generate=types,skip-prune $(OPENAPI) > pkg/models/types.gen.go
	@oapi-codegen -package=api -generate=types,server -include-tags=api $(OPENAPI) > internal/api/handlers.gen.go
	@oapi-codegen -package=views -generate=types,server -include-tags=ui $(OPENAPI) > internal/views/handlers.gen.go

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

$(DOCKER_TARGETS): docker-%:
	@docker compose run --rm app $(MAKE) $*

.PHONY: docker-clean-data
docker-clean-data:
	@docker compose down -v
