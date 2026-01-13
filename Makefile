service-run:
	sudo go run ./internal/cmd/main.go

wg-stop:
	sudo ip link delete dev wg0 2>/dev/null || true

wg-start:
	sudo ip link add dev wg0 type wireguard
	sudo ip address add dev wg0 10.0.0.1/24
	sudo ip link set up dev wg0

wg-restart: wg-stop wg-start

wg-status:
	sudo awg show wg0

