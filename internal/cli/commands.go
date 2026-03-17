package cli

import (
        "fmt"
        "os"
        "os/exec"
        "strings"

        "z0m4/internal/sandbox"
        "z0m4/internal/ui"
)

func Info(ghost bool) {
        fmt.Println(ui.Title("z0m4-distro"))
        fmt.Println("core:    z0m4F0rg3")
        fmt.Println("manager: Foundry")
}

func Asc(ghost bool) {
        fmt.Println(ui.Title("Asclepius"))
        fmt.Println(ui.Info("system diagnostics"))

        check("shell", detectShell(), "")
        check("environment", detectEnv(), "")

        check("git", checkCmd("git"), installCmd("git"))
        check("go", checkCmd("go"), installCmd("golang"))
        check("zf in PATH", checkCmd("zf"), "export PATH=$HOME/go/bin:$PATH")
       // fmt.Println(ui.Ok("system ready"))
        // opcionales útiles
        check("curl", checkCmd("curl"), installCmd("curl"))
        check("tree", checkCmd("tree"), installCmd("tree"))
        fmt.Println(ui.Ok("system ready"))
}

func check(name string, ok bool, fix string) {
        if ok {
                fmt.Println(ui.Ok(name + " ok"))
        } else {
                fmt.Println(ui.Err(name + " missing"))
                if fix != "" {
                        fmt.Println(ui.Info("install with: " + fix))
                }
        }
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

func SDX(ghost bool, args []string) {
        fmt.Println(ui.Title("SDX module"))
        sandbox.SDX(args)
}

// -------------------------
// ENV DETECTION
// -------------------------

func detectEnv() bool {
        if os.Getenv("TERMUX_VERSION") != "" {
                fmt.Println("env: termux")
                return true
        }

        data, err := os.ReadFile("/etc/os-release")
        if err == nil {
                txt := string(data)

                if strings.Contains(txt, "Alpine") {
                        fmt.Println("env: alpine")
                        return true
                }
                if strings.Contains(txt, "Debian") {
                        fmt.Println("env: debian")
                        return true
                }
        }

        fmt.Println("env: unknown")
        return false
}

func detectShell() bool {
        shell := os.Getenv("SHELL")
        if shell == "" {
                fmt.Println("shell: unknown")
                return false
        }

        fmt.Println("shell:", shell)
        return true
}

// -------------------------
// COMMAND CHECK
// -------------------------

func checkCmd(name string) bool {
        _, err := exec.LookPath(name)
        return err == nil
}

// -------------------------
// PACKAGE MANAGER DETECTION
// -------------------------

func getPkgManager() string {
        if os.Getenv("TERMUX_VERSION") != "" {
                return "pkg"
        }

        data, err := os.ReadFile("/etc/os-release")
        if err == nil {
                txt := string(data)

                if strings.Contains(txt, "Alpine") {
                        return "apk"
                }
                if strings.Contains(txt, "Debian") {
                        return "apt"
                }
        }

        return "unknown"
}

// -------------------------
// INSTALL SUGGESTIONS
// -------------------------

func installCmd(pkg string) string {
        pm := getPkgManager()

        switch pm {
        case "pkg":
                return "pkg install " + pkg
        case "apt":
                return "sudo apt install " + pkg
        case "apk":
                return "apk add " + pkg
        default:
                return "install " + pkg
        }
}
