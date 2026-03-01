.DEFAULT_GOAL := build

GO = go
LDFLAGS = -X 'main.GitTag=$(shell git describe --tags)'
GCFLAGS =

# CMD = $(shell basename $(CURDIR))
CMD = storeinit
BIN ?= $(CURDIR)/../../bin/$(CMD)
SOURCES = $(shell find "$(CURDIR)" -type f \
	-name "*.go" \
	-or -name "*.sql" \
	-or -name "*.json" \
	-or -name "Makefile")

$(BIN): $(SOURCES)
	@CGO_ENABLED=0 $(GO) build -v \
	-ldflags "$(LDFLAGS)" \
	-gcflags "$(GCFLAGS)" \
	-o $(BIN) main.go

.PHONY: build
build: $(BIN)

.PHONY: clean
clean:
	@rm -f $(BIN)

.PHONY: test
test:
	$(GO) test ./...

.PHONY: localdev
localdev:
	docker-compose --project-directory docker up -d

.PHONY: localdev-down
localdev-down:
	docker-compose --project-directory docker down --volumes --remove-orphans

.PHONY: localdev-reset
localdev-reset: localdev-down localdev

.PHONY: status
status: $(BIN)
	$(BIN) --status

.PHONY: up
up: $(BIN)
	$(BIN) --up

.PHONY: down
down: $(BIN)
	$(BIN) --down

.PHONY: force
force: $(BIN)
	$(BIN) --force
