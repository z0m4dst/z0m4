package ui

import "fmt"

func Ok(msg string) string {
	return fmt.Sprintf("\033[1;32m✓ %s\033[0m", msg)
}

func Err(msg string) string {
	return fmt.Sprintf("\033[1;31m✗ %s\033[0m", msg)
}

func Info(msg string) string {
	return fmt.Sprintf("\033[1;36m→ %s\033[0m", msg)
}

func Title(msg string) string {
	return fmt.Sprintf("\033[1;95m%s\033[0m", msg)
}
