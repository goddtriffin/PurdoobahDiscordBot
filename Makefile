$(VERBOSE).SILENT:
.DEFAULT_GOAL := help

.PHONY: help
help: # Prints out help
	@IFS=$$'\n' ; \
	help_lines=(`fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##/:/'`); \
	printf "%-30s %s\n" "target" "help" ; \
	printf "%-30s %s\n" "------" "----" ; \
	for help_line in $${help_lines[@]}; do \
			IFS=$$':' ; \
			help_split=($$help_line) ; \
			help_command=`echo $${help_split[0]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
			help_info=`echo $${help_split[2]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
			printf '\033[36m'; \
			printf "%-30s %s" $$help_command ; \
			printf '\033[0m'; \
			printf "%s\n" $$help_info; \
	done
	@echo

.PHONY: build
build: ## builds the binary locally
	CGO_ENABLED=0 go build -o bin/bot ./...

.PHONY: dev
dev: build ## runs the binary locally
	./bin/bot -token=$(PURDOOBAH_DISCORD_BOT_TOKEN)

.PHONY: build_docker
build_docker: ## builds the binary and Docker container
	docker build --tag goddtriffin/purdoobahs-discord-bot:latest --file deployment/Dockerfile .

.PHONY: run_docker
run_docker: build_docker ## creates and runs a new Docker container
	docker run \
	--name "purdoobahs_discord_bot" \
	-d --restart unless-stopped \
	-p 8080:8080 \
	-e PURDOOBAH_DISCORD_BOT_TOKEN \
	goddtriffin/purdoobahs-discord-bot:latest

.PHONY: start_docker
start_docker: ## resumes a stopped Docker container
	docker start purdoobahs_discord_bot

.PHONY: stop_docker
stop_docker: ## stops the Docker container
	docker stop purdoobahs_discord_bot

.PHONY: remove_docker
remove_docker: ## removes the Docker container
	docker rm purdoobahs_discord_bot

.PHONY: push_docker
push_docker: ## pushes new Docker image to Docker Hub
	docker push goddtriffin/purdoobahs-discord-bot:latest

.PHONY: restart_deployment
restart_deployment: ## restarts all pods in the purdoobahs-website k8s deployment
	kubectl rollout restart deployment purdoobahs-discord-bot

.PHONY: deploy
deploy: build_docker push_docker restart_deployment # builds/pushes new docker image at :latest and restarts k8s deployment

.PHONY: mem_usage
mem_usage: ## displays the memory usage of the currently running Docker container
	docker stats purdoobahs_discord_bot --no-stream --format "{{.Container}}: {{.MemUsage}}"

.PHONY: logs
logs: ## displays logs
	docker logs purdoobahs_discord_bot
