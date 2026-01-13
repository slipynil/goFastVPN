package core

import (
	"os/exec"
)

func CreateInterface() {
	exec.Command("amneziawg-go wg0")
}
