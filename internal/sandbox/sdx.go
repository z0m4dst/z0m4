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

	// 📄 Caso 1: archivo
	if len(args) > 0 {
		script, err = os.ReadFile(args[0])
		if err != nil {
			fmt.Println(ui.Err(fmt.Sprintf("cannot read file: %v", err)))
			return
		}
	} else {
		// 📥 stdin (pipe o interactivo)
		info, _ := os.Stdin.Stat()

		if (info.Mode() & os.ModeCharDevice) == 0 {
			// pipe
			script, err = io.ReadAll(os.Stdin)
		} else {
			// interactivo
			fmt.Println(ui.Info("paste script then press CTRL+D"))
			script, err = io.ReadAll(os.Stdin)
		}

		if err != nil {
			fmt.Println(ui.Err(fmt.Sprintf("cannot read stdin: %v", err)))
			return
		}
	}

	// 🔒 validación básica
	if len(script) == 0 {
		fmt.Println(ui.Err("empty script"))
		return
	}

	// 📁 temp file (portable)
	tmpFile, err := os.CreateTemp("", "z0m4_sdx_*.sh")
	if err != nil {
		fmt.Println(ui.Err(fmt.Sprintf("cannot create temp file: %v", err)))
		return
	}
	defer os.Remove(tmpFile.Name())

	// ✍️ escribir script
	_, err = tmpFile.Write(script)
	if err != nil {
		fmt.Println(ui.Err(fmt.Sprintf("cannot write temp script: %v", err)))
		return
	}

	// permisos
	err = tmpFile.Chmod(0755)
	if err != nil {
		fmt.Println(ui.Err(fmt.Sprintf("cannot set permissions: %v", err)))
		return
	}

	tmpFile.Close()

	// 🚀 ejecución
	fmt.Println(ui.Info("executing script"))

	cmd := exec.Command("sh", tmpFile.Name())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println(ui.Err(fmt.Sprintf("execution failed: %v", err)))
		return
	}

	// ✅ éxito
	fmt.Println(ui.Ok("sandbox ok"))
}
