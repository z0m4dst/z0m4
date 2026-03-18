package forge

import (
	"fmt"
	"os"
	"os/exec"
)

func Run(args []string) {

	if len(args) < 1 {
		fmt.Println("usage: zf forge <init|run|list|remove|reset>")
		return
	}

	switch args[0] {

	case "init":
		initForge()

	case "run":
		runForge(args[1:])

	case "list":
		listForge()

	case "remove":
		removeForge(args[1:])

	case "reset":
		resetForge()

	default:
		fmt.Println("unknown command:", args[0])
	}
}

// ---------------- INIT ----------------

func initForge() {

	fmt.Println("→ initializing forge")
	fmt.Println("→ available distros:\n")

	cmd := exec.Command("proot-distro", "list")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	fmt.Print("\nselect distro (write name): ")

	var distro string
	fmt.Scanln(&distro)

	if distro == "" {
		fmt.Println("✗ no distro selected")
		return
	}

	fmt.Println("→ installing", distro, "...")

	install := exec.Command("proot-distro", "install", distro)
	install.Stdout = os.Stdout
	install.Stderr = os.Stderr
	install.Stdin = os.Stdin

	err := install.Run()
	if err != nil {
		fmt.Println("✗ install failed:", err)
		return
	}

	fmt.Println("✓ forge ready:", distro)
}

// ---------------- RUN ----------------

func runForge(args []string) {

	if len(args) < 1 {
		fmt.Println("usage: zf forge run <distro>")
		return
	}

	distro := args[0]

	fmt.Println("→ entering", distro)

	cmd := exec.Command("proot-distro", "login", distro)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		fmt.Println("✗ failed:", err)
	}
}

// ---------------- LIST ----------------

func listForge() {

	fmt.Println("→ installed distros:\n")

	cmd := exec.Command("proot-distro", "list", "--installed")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

// ---------------- REMOVE ----------------

func removeForge(args []string) {

	if len(args) < 1 {
		fmt.Println("usage: zf forge remove <distro>")
		return
	}

	distro := args[0]

	fmt.Println("→ removing", distro)

	cmd := exec.Command("proot-distro", "remove", distro)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("✗ failed:", err)
		return
	}

	fmt.Println("✓ removed:", distro)
}

// ---------------- RESET ----------------

func resetForge() {

	fmt.Println("⚠ this will remove ALL distros")
	fmt.Print("confirm (y/N): ")

	var confirm string
	fmt.Scanln(&confirm)

	if confirm != "y" {
		fmt.Println("aborted")
		return
	}

	cmd := exec.Command("proot-distro", "list", "--installed")
	out, _ := cmd.Output()

	lines := string(out)
	var distros []string

	for _, line := range splitLines(lines) {
		if line != "" {
			distros = append(distros, line)
		}
	}

	for _, d := range distros {
		fmt.Println("→ removing", d)

		exec.Command("proot-distro", "remove", d).Run()
	}

	fmt.Println("✓ forge reset complete")
}

// helper simple (sin strings pkg para mantenerlo básico)
func splitLines(s string) []string {
	var result []string
	current := ""

	for _, c := range s {
		if c == '\n' {
			result = append(result, current)
			current = ""
		} else {
			current += string(c)
		}
	}

	if current != "" {
		result = append(result, current)
	}

	return result
}
