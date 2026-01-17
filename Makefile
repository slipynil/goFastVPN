include .env
export

su-run:
	@sudo JS=$(JS) JMIN=$(JMIN) JMAX=$(JMAX) S1=$(S1) S2=$(S2) \
	     H1=$(H1) H2=$(H2) H3=$(H3) H4=$(H4) \
	     DEVICE=$(DEVICE) ENDPOINT=$(ENDPOINT) \
		 TELEGRAM_KEY=$(TELEGRAM_KEY) \
	     go run internal/cmd/main.go

run:
	@go run internal/cmd/main.go

wg-stop:
	sudo ip link delete dev awg0 2>/dev/null || true

wg-start:
	sudo ip link add dev awg0 type wireguard
	sudo ip address add dev awg0 10.0.0.1/24
	sudo ip link set up dev awg0

wg-restart: wg-stop wg-start

wg-status:
	sudo awg show awg0
