# build tags
TAGS      ?= must,sugar
# patterns for subprojects
TARGET    ?= less* contest*
# exact dirs to exclude
EXCLUDE   ?=
# path to local contestio
LOCAL_LIB ?= ./lib/contestio

GO        := go
GOWORK    := go.work

MAIN_DIRS := $(shell find $(TARGET) -type f -name main.go -exec dirname {} \; | sort -u)
MAIN_DIRS_FILTERED := $(filter-out $(EXCLUDE),$(MAIN_DIRS))

SHOW_LOCAL  := $(GO) list -m -f 'LOCAL: {{.Dir}}' github.com/aaa2ppp/contestio
SHOW_REMOTE := $(GO) list -m -f 'REMOTE: github.com/aaa2ppp/contestio@{{.Version}}' github.com/aaa2ppp/contestio 

.PHONY: all state local remote build test list clean-cache help

all: state build test ## show state, then build and test

state: ## display current contestio mode (local/remote)
	@if [ -f $(GOWORK) ]; then \
		$(SHOW_LOCAL); \
	else \
		$(SHOW_REMOTE); \
	fi

local: ## switch to local contestio library (creates go.work)
	@set -e; \
	test -d "$(LOCAL_LIB)"; \
	$(GO) work init . $(LOCAL_LIB) || $(GO) work use $(LOCAL_LIB); \
	$(SHOW_LOCAL)

remote: ## switch to remote contestio (removes go.work)
	@rm -f $(GOWORK) $(GOWORK).sum 2>/dev/null; \
	$(SHOW_REMOTE)

list: ## list solutions
	@for dir in $(MAIN_DIRS_FILTERED); do echo $$dir; done

build: ## test build solutions
	@echo "build -tags=$(TAGS)" >&2; \
	failed=0; \
	for dir in $(MAIN_DIRS_FILTERED); do \
		echo -n "$$dir ... " >&2; \
		$(GO) build -o /dev/null -tags=$(TAGS) ./$$dir; \
		exit_code=$$?; \
		case $$exit_code in \
			0) echo "ok" >&2;; \
			127|130|143|2) >&2 echo "interrupted"; exit 1;; \
			*) echo "failed ($$exit_code)" >&2; failed=1;; \
		esac; \
	done; \
	if [ $$failed -eq 1 ]; then \
		echo "Some builds failed." >&2; exit 1; \
	else \
		echo "All builds succeeded." >&2; \
	fi

test: ## test solutions
	@echo "test -tags=$(TAGS)" >&2; \
	failed=0; \
	for dir in $(MAIN_DIRS_FILTERED); do \
		$(GO) test -tags=$(TAGS) ./$$dir; \
		exit_code=$$?; \
		case $$exit_code in \
			0) ;; \
			127|130|143|2) >&2 echo "interrupted"; exit 1;; \
			*) echo "failed ($$exit_code)" >&2; failed=1;; \
		esac; \
	done; \
	if [ $$failed -eq 1 ]; then \
		echo "Some tests failed." >&2; exit 1; \
	else \
		echo "All tests passed." >&2; \
	fi

clean-cache: ## clean go test cache
	$(GO) clean -testcache ./...

help: ## show this help
	@printf "Usage: make [target] [VARIABLE=value]\n\n"
	@printf "Variables:\n"
	@awk 'BEGIN {comment=""} \
		/^[a-zA-Z0-9_-]+[[:space:]]*\?=/ { \
			split($$0, a, "?="); \
			if ( prev ~ /^#/ ) { \
				printf "  %-14s = %-20s %s\n", a[1], a[2], prev; \
			} else { \
				printf "  %-14s = %-20s\n", a[1], a[2]; \
			} \
		} \
		{ prev=$$0 }' $(MAKEFILE_LIST)
	@printf "\nTargets:\n"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {printf "  %-14s - %s\n", $$1, $$2}' $(MAKEFILE_LIST)
