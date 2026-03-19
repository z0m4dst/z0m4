package main

import (
	"fmt"
	"os"

	"z0m4/internal/cli"
	"z0m4/internal/hef"
)

func main() {

	if len(os.Args) < 2 {
		help()
		return
	}

	cmd := os.Args[1]

	switch cmd {

	// =========================
	// SYSTEM
	// =========================

	case "info":
		cli.Info(false)

	case "asc":
		cli.Asc(false)

	case "doctor":
		cli.Asc(false)

	case "up":
		cli.Update(false)

	case "list":
		fmt.Println("system modules (todo)")

	// =========================
	// SANDBOX
	// =========================

	case "sdx":
		cli.SDX(false, os.Args[2:])

	// =========================
	// ENV (hef)
	// =========================

	case "i":
		if len(os.Args) < 3 {
			fmt.Println("✗ missing distro")
			return
		}
		hef.Run([]string{"init", os.Args[2]})

	case "r":
		if len(os.Args) < 3 {
			fmt.Println("✗ missing distro")
			return
		}
		hef.Run([]string{"run", os.Args[2]})

	case "l":
		hef.Run([]string{"list"})

	case "rm":
		handleRemove(os.Args[2:])

	case "hef":
		hef.Run(os.Args[2:])

	// =========================
	// DEFAULT
	// =========================

	default:
		help()
	}
}

//
// =========================
// REMOVE INTELIGENTE
// =========================
//

func handleRemove(args []string) {

	if len(args) == 0 {
		fmt.Println("✗ missing target")
		return
	}

	target := args[0]

	// -------------------------
	// AUTODETECT DISTRO
	// -------------------------

	if isDistroInstalled(target) {
		fmt.Println("→ removing distro:", target)
		hef.Run([]string{"remove", target})
		return
	}

	// -------------------------
	// FALLBACK MODULE
	// -------------------------

	fmt.Println("→ removing module:", target)
}

//
// =========================
// CHECK DISTRO INSTALADA
// =========================
//

func isDistroInstalled(name string) bool {

	base := os.Getenv("PREFIX") + "/var/lib/proot-distro/installed-rootfs"

	_, err := os.Stat(base + "/" + name)
	return err == nil
}

//
// =========================
// HELP
// =========================
//

func help() {

	fmt.Println("z0m4 CLI:\n")

	fmt.Println("system:")
	fmt.Println("  zf info")
	fmt.Println("  zf doctor")
	fmt.Println("  zf list")
	fmt.Println("  zf up")

	fmt.Println()
	fmt.Println("env:")
	fmt.Println("  zf i <distro>")
	fmt.Println("  zf r <distro>")
	fmt.Println("  zf l")
	fmt.Println("  zf rm <target>")
	fmt.Println("  zf hef")

	fmt.Println()
	fmt.Println("tools:")
	fmt.Println("  zf sdx [script]")
}
