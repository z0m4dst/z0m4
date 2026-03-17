package main

import (
	"fmt"
	"os"

	"z0m4/internal/cli"
)

func help() {
	fmt.Println("z0m4-distro CLI")
	fmt.Println()
	fmt.Println("zf info")
	fmt.Println("zf doctor")
	fmt.Println("zf i <module>")
	fmt.Println("zf r <module>")
	fmt.Println("zf up")
	fmt.Println("zf sdx [script]")
	fmt.Println()
	fmt.Println("flags:")
	fmt.Println("-g   ghost mode")
}

func main() {

	if len(os.Args) < 2 {
		help()
		return
	}

	cmd := os.Args[1]

	switch cmd {

	case "info":
		cli.Info(false)

	case "asc":
                cli.Asc(false)
       case "doctor":
                cli.Asc(false)

	case "i":
		cli.Install(false)

	case "r":
		cli.Remove(false)

	case "up":
		cli.Update(false)

	case "sdx":
         cli.SDX(false, os.Args[2:])

	default:
		help()
	}
}
