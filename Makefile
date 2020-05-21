# go commands
GOCMD := go
GOTEST := $(GOCMD) test
GOTOOL := $(GOCMD) tool
GOBUILD := $(GOCMD) build

COVERPROFILE ?= /tmp/profile.out

.PHONY: test
test:
	$(GOTEST) ./...

.PHONY: coverage
coverage:
	$(GOTEST) -covermode=atomic -coverprofile=$(COVERPROFILE) ./...

.PHONY: showcoverage
showcoverage: coverage
	$(GOTOOL) cover -html=$(COVERPROFILE)

.PHONY: buildlocal
buildlocal:
	$(GOBUILD) -o bin/bot-local ./...

.PHONY: runlocal
runlocal: buildlocal
	./bin/bot-local -token=$(PURDOOBAH_DISCORD_BOT_TOKEN)
