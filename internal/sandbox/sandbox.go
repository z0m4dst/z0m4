package sandbox

import (
	"fmt"
	"os/exec"
)

func Run(script string) {

	cmd := exec.Command("/bin/sh", script)

	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("✗ sandbox error — see log")
		fmt.Println(string(output))
		return
	}

	fmt.Println("✓ sandbox ok")
}
