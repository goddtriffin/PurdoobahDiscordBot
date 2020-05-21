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
	CGO_ENABLED=0 $(GOBUILD) -o bin/bot-local ./...

.PHONY: runlocal
runlocal: buildlocal
	./bin/bot-local -token=$(PURDOOBAH_DISCORD_BOT_TOKEN)

.PHONY: builddocker
builddocker:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o bin/bot-docker ./...
	docker build --tag purdoobah-discord-bot --file build/Dockerfile .

.PHONY: rundocker
rundocker: builddocker
	docker rm purdoobah_discord_bot
	docker run \
	--name "purdoobah_discord_bot" \
	-d --restart unless-stopped \
	-e PURDOOBAH_DISCORD_BOT_TOKEN \
	purdoobah-discord-bot

.PHONY: stopdocker
stopdocker:
	docker stop purdoobah_discord_bot

.PHONY: memusage
memusage:
	docker stats purdoobah_discord_bot --no-stream --format "{{.Container}}: {{.MemUsage}}"
