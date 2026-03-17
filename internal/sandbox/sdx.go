package sandbox

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"z0m4/internal/ui"
)

func SDX(args []string) {

	var script []byte
	var err error

	// Caso 1: argumento (archivo)
	if len(args) > 0 {

		script, err = os.ReadFile(args[0])
		if err != nil {
			fmt.Println("✗ sandbox error — cannot read file")
			return
		}

	} else {

		// Caso 2: stdin (pipe o paste)
		info, _ := os.Stdin.Stat()

		if (info.Mode() & os.ModeCharDevice) == 0 {
			// pipe
			script, err = io.ReadAll(os.Stdin)
		} else {
			// paste manual
			fmt.Println("paste script then press CTRL+D")
			script, err = io.ReadAll(os.Stdin)
		}

		if err != nil {
			fmt.Println("✗ sandbox error — see log")
			return
		}
	}

	tmp := "/tmp/z0m4_sdx.sh"

	err = os.WriteFile(tmp, script, 0755)
	if err != nil {
		fmt.Println("✗ sandbox error — cannot create temp script")
		return
	}

	cmd := exec.Command("sh", tmp)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()

	if err != nil {
		fmt.Println("✗ sandbox error — see log")
		return
	}

	fmt.Println(ui.Ok("sandbox ok"))

}
