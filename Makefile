.DEFAULT_GOAL := all
MAKEFLAGS += --no-print-directory

# DO NOT EDIT BY HAND
# To set a different name, use:
# $ export ENVIRONMENT="new_env"
# from your shell.

# Where to put all builded binaries
BINDIR ?= $(CURDIR)/bin

TRIPPY_BIN = $(BINDIR)/trippy
TRIPPY_CONFIG_ROOT_DIR ?= $(CURDIR)
TRIPPY_SOURCES = $(shell find $(CURDIR) -type f -name "*.go" -or -name "*.sql")

SWAGGER_OUTPUT = $(CURDIR)/docs

STOREINIT_BIN = $(BINDIR)/storeinit

GO = go
GOMODULE = $(shell head -1 go.mod | cut -d" " -f2)
LDFLAGS = -X 'main.GitTag=$(shell git describe --tags)'
GCFLAGS =

.PHONY: all
all: trippy storeinit

.PHONY: storeinit
storeinit:
	@echo "Building $@"
	@cd cmd/_$@ && $(MAKE)

.PHONY: trippy
trippy: $(TRIPPY_BIN)


$(TRIPPY_BIN): $(TRIPPY_SOURCES)

	@echo "Building Trippy"

	@echo "    Generating swagger files"
	@swag init --quiet --dir ./cmd/trippy,. --parseDependency --parseInternal --parseDepth 1 --output $(SWAGGER_OUTPUT) --outputTypes yaml,go

	@echo "    Compiling"
	@echo
	@CGO_ENABLED=0 $(GO) build \
	-ldflags "$(LDFLAGS)" \
	-gcflags "$(GCFLAGS)" \
	-o $(TRIPPY_BIN) cmd/trippy/*.go

.PHONY: run
run: $(TRIPPY_BIN)
	@echo "Loading environment variables from .env"
	@export $$(grep -v '^#' .env | xargs) && \
	$(TRIPPY_BIN)

.PHONY: tidy
tidy:
	$(GO) mod tidy

.PHONY: clean
clean:
	@rm -f $(BINDIR)/*


.PHONY: test
test:
	$(GO) test ./...

.PHONY: run-serve
run-serve:
	swag init --parseDependency --parseDepth 1 -g main.go
	$(GO) run main.go


.PHONY: deploy
deploy:
	@docker-compose --env-file ./.env -f ./docker/local/docker-compose.yml up -d

.PHONY: local
local:
	@cd docker/local && $(MAKE) up

.PHONY: local-down
local-down:
	@cd docker/local && $(MAKE) down

.PHONY: local-stop
local-stop:
	@cd docker/local && $(MAKE) stop

.PHONY: local-reset
local-reset:
	@cd docker/local && $(MAKE) reset

