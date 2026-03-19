package sandbox

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"

	"z0m4/internal/ui"
)

func SDX(args []string) {

	tmpDir := os.TempDir()
	tmpFile, err := os.CreateTemp(tmpDir, "sdx_*.sh")
	if err != nil {
		fmt.Println(ui.Err(fmt.Sprintf("cannot create temp file: %v", err)))
		return
	}
	defer os.Remove(tmpFile.Name())

	var script []byte

	// =========================
	// INPUT
	// =========================

	if len(args) > 0 {

		// archivo
		if _, err := os.Stat(args[0]); err == nil {
			script, err = os.ReadFile(args[0])
			if err != nil {
				fmt.Println(ui.Err(fmt.Sprintf("cannot read file: %v", err)))
				return
			}
		} else {
			// string directa
			script = []byte(args[0])
		}

	} else {

		info, _ := os.Stdin.Stat()

		// pipe
		if (info.Mode() & os.ModeCharDevice) == 0 {
			script, err = io.ReadAll(os.Stdin)
			if err != nil {
				fmt.Println(ui.Err(fmt.Sprintf("cannot read stdin: %v", err)))
				return
			}
		} else {
			// interactivo
			fmt.Println(ui.Info("paste script then press CTRL+D"))
			script, err = io.ReadAll(os.Stdin)
			if err != nil {
				fmt.Println(ui.Err(fmt.Sprintf("cannot read input: %v", err)))
				return
			}
		}
	}

	if len(script) == 0 {
		fmt.Println(ui.Err("empty script"))
		return
	}

	// =========================
	// ANALYZE
	// =========================

	fmt.Println(ui.Info("analyzing"))

	content := string(script)

	if containsDanger(content) {
		fmt.Println(ui.Err("dangerous command detected"))
		return
	}

	fmt.Println(ui.Ok("safe"))

	// =========================
	// WRITE TMP
	// =========================

	_, err = tmpFile.Write(script)
	if err != nil {
		fmt.Println(ui.Err(fmt.Sprintf("cannot write temp file: %v", err)))
		return
	}

	err = tmpFile.Chmod(0755)
	if err != nil {
		fmt.Println(ui.Err(fmt.Sprintf("cannot set permissions: %v", err)))
		return
	}

	tmpFile.Close()

	// =========================
	// SYNTAX CHECK
	// =========================

	fmt.Println(ui.Info("checking syntax"))

	check := exec.Command("sh", "-n", tmpFile.Name())

	err = check.Run()
	if err != nil {
		fmt.Println(ui.Err("syntax error"))
		return
	}

	fmt.Println(ui.Ok("syntax ok"))

	// =========================
	// EXECUTE + TIMEOUT
	// =========================

	fmt.Println(ui.Info("executing"))

	cmd := exec.Command("sh", tmpFile.Name())

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// entorno mínimo
	cmd.Env = []string{
		"PATH=/usr/bin:/bin",
	}

	done := make(chan error, 1)

	go func() {
		done <- cmd.Run()
	}()

	select {
	case err := <-done:
		if err != nil {
			fmt.Println(ui.Err(fmt.Sprintf("execution failed: %v", err)))
			return
		}
	case <-time.After(3 * time.Second):
		cmd.Process.Kill()
		fmt.Println(ui.Err("execution timeout"))
		return
	}

	fmt.Println(ui.Ok("sandbox ok"))
}

// =========================
// ANALYZER
// =========================

func containsDanger(s string) bool {

	dangers := []string{
		"rm -rf /",
		"mkfs",
		"dd if=",
	}

	for _, d := range dangers {
		if contains(s, d) {
			return true
		}
	}

	return false
}

func contains(s, sub string) bool {
	return len(s) >= len(sub) && stringContains(s, sub)
}

func stringContains(s, sub string) bool {
	return len(sub) == 0 || (len(s) >= len(sub) && indexOf(s, sub) >= 0)
}

func indexOf(s, sub string) int {
	for i := 0; i <= len(s)-len(sub); i++ {
		if s[i:i+len(sub)] == sub {
			return i
		}
	}
	return -1
}
