AWG_ENV := services/awg/.env

.PHONY: awg-run
awg-run:
	@export $$(cat $(AWG_ENV) | xargs) && \
	cd services/awg && sudo -E go run cmd/main.go

container-up:
	@docker container up -d
container-down:
	@docker container down
container-logs:
	@docker container logs
