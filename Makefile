# Makefile for ya-training9 (Windows / Git Bash friendly)

TAGS      = must,sugar
GO        = go
GOWORK    = go.work
LOCAL_LIB = ./lib/contestio

# Find all directories containing main.go (skip root and lib)
MAIN_DIRS = $(shell find ./less* ./contest* -type f -name main.go -exec dirname {} \; | sort -u)

.PHONY: all info build test use-local use-remote clean

all: info build test

info:
	@echo "=== Project Information ==="
	@echo "Go version: $(shell $(GO) version)"
	@if [ -f "$(GOWORK)" ]; then \
		echo "go.work: active (LOCAL library in use)"; \
		echo "  Library location: $$($(GO) list -m -f '{{.Dir}}' github.com/aaa2ppp/contestio 2>/dev/null)"; \
	else \
		echo "go.work: absent (REMOTE library from go.mod)"; \
		$(GO) list -m -f '  Using: github.com/aaa2ppp/contestio@{{.Version}}' github.com/aaa2ppp/contestio 2>/dev/null || echo "  contestio not found in go.mod"; \
	fi
	@echo "Number of subprojects: $(words $(MAIN_DIRS))"
	@echo ""
	
build:
	@echo "=== Building all subprojects (tags=$(TAGS)) ==="
	@failed=0; \
	for dir in $(MAIN_DIRS); do \
		echo -n "$$dir ... "; \
		if $(GO) build -o /dev/null -tags=$(TAGS) ./$$dir; then \
			echo "ok"; \
		else \
			echo "failed"; \
			failed=1; \
		fi; \
	done; \
	if [ $$failed -eq 1 ]; then \
		echo ""; \
		echo "Some builds failed."; \
		exit 1; \
	else \
		echo ""; \
		echo "All builds succeeded."; \
	fi

test:
	@echo "=== Testing all subprojects (tags=$(TAGS)) ==="
	@failed=0; \
	for dir in $(MAIN_DIRS); do \
		if ! $(GO) test -tags=$(TAGS) ./$$dir; then \
			failed=1; \
		fi; \
	done; \
	if [ $$failed -eq 1 ]; then \
		echo ""; \
		echo "Some tests failed."; \
		exit 1; \
	else \
		echo ""; \
		echo "All tests passed."; \
	fi

use-local:
	@echo "Switching to LOCAL contestio library..."
	@if [ -f "$(GOWORK)" ]; then \
		echo "go.work already exists."; \
	else \
		echo "go 1.24.2" > $(GOWORK); \
		echo "" >> $(GOWORK); \
		echo "use (" >> $(GOWORK); \
		echo "	." >> $(GOWORK); \
		echo "	$(LOCAL_LIB)" >> $(GOWORK); \
		echo ")" >> $(GOWORK); \
		echo "Created $(GOWORK)."; \
	fi
	@echo "Now using LOCAL contestio from $(LOCAL_LIB)."
	@echo "Run 'make info' to verify."
	@echo ""

use-remote:
	@echo "Switching to REMOTE contestio library..."
	@if [ -f "$(GOWORK)" ]; then \
		rm -f "$(GOWORK)"; \
		echo "Removed $(GOWORK)."; \
	else \
		echo "go.work not found, already using remote version."; \
	fi
	@echo "Now using the version specified in go.mod."
	@echo "Run 'make info' to verify."
	@echo ""

clean: use-remote
	@echo "Cleaned up."
