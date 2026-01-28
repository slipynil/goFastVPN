awg_env := services/awg/.env
pay_env := services/pay/.env

.phony: awg-run pay-run
awg-run:
	@export $$(cat $(awg_env) | xargs) && \
	cd services/awg && sudo -e go run cmd/main.go

pay-run:
	@export $$(cat $(pay_env) | xargs) && \
	cd services/pay && go run cmd/main.go

compose-up:
	@cd services/telegram && docker compose up -d
compose-down:
	@cd services/telegram && docker compose down
compose-logs:
	@cd services/telegram && docker compose logs
