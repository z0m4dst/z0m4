package ui

const (
	Reset = "\033[0m"
	Bold  = "\033[1m"

	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Cyan   = "\033[36m"
)

func Title(text string) string {
	return Bold + Cyan + text + Reset
}

func Ok(text string) string {
	return Green + "✓ " + text + Reset
}

func Err(text string) string {
	return Red + "✗ " + text + Reset
}

func Warn(text string) string {
	return Yellow + "• " + text + Reset
}
