package hef

import (
	"fmt"
	"os"
	"os/exec"
)

func Run(args []string) {

	if len(args) < 1 {
		fmt.Println("usage: zf hef <init|run|list|remove|reset>")
		return
	}

	switch args[0] {

	case "init":
		initHef(args[1:])

	case "run":
		runHef(args[1:])

	case "list":
		listHef()

	case "remove":
		removeHef(args[1:])

	case "reset":
		resetHef()

	default:
		fmt.Println("unknown command:", args[0])
	}
}

//
// =========================
// INIT (install)
// =========================
//

func initHef(args []string) {

	if len(args) < 1 {
		fmt.Println("usage: zf i <distro>")
		return
	}

	distro := args[0]

	fmt.Println("→ installing", distro, "...")

	cmd := exec.Command("proot-distro", "install", distro)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		fmt.Println("✗ install failed:", err)
		return
	}

	fmt.Println("✓ ready:", distro)
}

//
// =========================
// RUN (login)
// =========================
//

func runHef(args []string) {

	if len(args) < 1 {
		fmt.Println("usage: zf r <distro>")
		return
	}

	distro := args[0]

	if !isInstalled(distro) {
		fmt.Println("✗ distro not installed:", distro)
		return
	}

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

//
// =========================
// LIST (installed)
// =========================
//

func listHef() {

	base := getBasePath()

	entries, err := os.ReadDir(base)
	if err != nil {
		fmt.Println("✗ cannot read distros:", err)
		return
	}

	if len(entries) == 0 {
		fmt.Println("no distros installed")
		return
	}

	fmt.Println("installed distros:\n")

	for _, e := range entries {
		if e.IsDir() {
			fmt.Println(" -", e.Name())
		}
	}
}

//
// =========================
// REMOVE
// =========================
//

func removeHef(args []string) {

	if len(args) < 1 {
		fmt.Println("usage: zf rm <distro>")
		return
	}

	distro := args[0]

	if !isInstalled(distro) {
		fmt.Println("✗ distro not installed:", distro)
		return
	}

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

//
// =========================
// RESET (all distros)
// =========================
//

func resetHef() {

	fmt.Println("⚠ this will remove ALL distros")
	fmt.Print("confirm (y/N): ")

	var confirm string
	fmt.Scanln(&confirm)

	if confirm != "y" {
		fmt.Println("aborted")
		return
	}

	base := getBasePath()

	entries, err := os.ReadDir(base)
	if err != nil {
		fmt.Println("✗ failed:", err)
		return
	}

	for _, e := range entries {
		if e.IsDir() {
			distro := e.Name()
			fmt.Println("→ removing", distro)
			exec.Command("proot-distro", "remove", distro).Run()
		}
	}

	fmt.Println("✓ reset complete")
}

//
// =========================
// HELPERS
// =========================
//

func getBasePath() string {
	return os.Getenv("PREFIX") + "/var/lib/proot-distro/installed-rootfs"
}

func isInstalled(distro string) bool {

	base := getBasePath()

	_, err := os.Stat(base + "/" + distro)
	return err == nil
}
