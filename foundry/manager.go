package foundry

import (
	"fmt"
	"os/exec"
)

func Status() {
	check("go")
	check("git")
	check("hx")
}

func check(cmd string) {

	_, err := exec.LookPath(cmd)

	if err != nil {
		fmt.Printf("%-10s missing\n", cmd)
		return
	}

	fmt.Printf("%-10s ok\n", cmd)
}
