package cli

import (
	"fmt"

	"z0m4/foundry"
	"z0m4/internal/sandbox"
)

func Info(ghost bool) {
	fmt.Println("z0m4-distro")
	fmt.Println("core: z0m4F0rg3")
	fmt.Println("manager: Foundry")
}

func Doctor(ghost bool) {
	fmt.Println("system check")
	fmt.Println("core ........ ok")
	foundry.Status()
}

func Install(ghost bool) {
	fmt.Println("install module (todo)")
}

func Remove(ghost bool) {
	fmt.Println("remove module (todo)")
}

func Update(ghost bool) {
	fmt.Println("update system (todo)")
}

func Sdx(args []string) {
	fmt.Println("SDX module loaded")
	sandbox.SDX(args)
}
