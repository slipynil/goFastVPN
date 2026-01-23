AWG_ENV := services/awg/.env

.PHONY: awg-run
awg-run:
	@export $$(cat $(AWG_ENV) | xargs) && \
	cd services/awg && sudo -E go run cmd/main.go

compose-up:
	@cd services/telegram && docker compose up -d
compose-down:
	@cd services/telegram && docker compose down
compose-logs:
	@cd services/telegram && docker compose logs
