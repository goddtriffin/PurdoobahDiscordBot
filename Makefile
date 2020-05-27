.DEFAULT_GOAL := usage
COVERPROFILE ?= /tmp/profile.out

.PHONY: usage
usage:
	@echo "Usage:"
	@echo "=========="
	@echo "make usage - display Makefile target info"
	@echo "make buildlocal - builds the binary locally"
	@echo "make runlocal - runs the binary locally"
	@echo "make builddocker - builds the binary and Docker container"
	@echo "make rundocker - creates and runs a new Docker container"
	@echo "make startdocker - resumes a stopped Docker container"
	@echo "make stopdocker - stops the Docker container"
	@echo "make removedocker - removes the Docker container"
	@echo "make memusage - displays the memory usage of the currently running Docker container"

.PHONY: buildlocal
buildlocal:
	CGO_ENABLED=0 go build -o bin/bot-local ./...

.PHONY: runlocal
runlocal: buildlocal
	./bin/bot-local -token=$(PURDOOBAH_DISCORD_BOT_TOKEN)

.PHONY: builddocker
builddocker:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/bot-docker ./...
	docker build --tag purdoobah-discord-bot --file build/Dockerfile .

.PHONY: rundocker
rundocker: builddocker
	docker run \
	--name "purdoobah_discord_bot" \
	-d --restart unless-stopped \
	-e PURDOOBAH_DISCORD_BOT_TOKEN \
	purdoobah-discord-bot

.PHONY: startdocker
startdocker:
	docker start purdoobah_discord_bot

.PHONY: stopdocker
stopdocker:
	docker stop purdoobah_discord_bot

.PHONY: removedocker
removedocker:
	docker rm purdoobah_discord_bot

.PHONY: memusage
memusage:
	docker stats purdoobah_discord_bot --no-stream --format "{{.Container}}: {{.MemUsage}}"
