service-run:
	sudo go run ./internal/cmd/main.go

wg-stop:
	sudo ip link delete dev awg0 2>/dev/null || true

wg-start:
	sudo ip link add dev awg0 type wireguard
	sudo ip address add dev awg0 10.0.0.1/24
	sudo ip link set up dev awg0

wg-restart: wg-stop wg-start

wg-status:
	sudo awg show awg0

