AWG_ENV := services/awg/.env
TELEGRAM_ENV := services/telegram/.env
DB_CONN := $$(grep DB_CONN $(TELEGRAM_ENV) | cut -d '=' -f2-)

.PHONY: awg-run telegram-run migrate-up migrate-down migrate-version
awg-run:
	@export $$(cat $(AWG_ENV) | xargs) && \
	cd services/awg && sudo -E go run cmd/main.go

telegram-run:
	@export $$(cat $(TELEGRAM_ENV) | xargs) && \
	cd services/telegram && go run cmd/main.go

migrate-up:
	@migrate -path=./services/telegram/migrations -database "$(DB_CONN)" up
migrate-down:
	@migrate -path=./services/telegram/migrations -database "$(DB_CONN)" down
migrate-version:
	@migrate -path=./services/telegram/migrations -database "$(DB_CONN)" version
