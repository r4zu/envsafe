package internal

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorCyan   = "\033[36m"
	ColorBold   = "\033[1m"
)

// Green wraps string in Green ANSI
func Green(s string) string {
	return ColorGreen + s + ColorReset
}

// Red wraps string in Red ANSI
func Red(s string) string {
	return ColorRed + s + ColorReset
}

// Yellow wraps string in Yellow ANSI
func Yellow(s string) string {
	return ColorYellow + s + ColorReset
}

// Bold wraps string in Bold ANSI
func Bold(s string) string {
	return ColorBold + s + ColorReset
}

// Cyan wraps string in Cyan ANSI
func Cyan(s string) string {
	return ColorCyan + s + ColorReset
}
