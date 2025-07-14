// colors.go
package internal

import (
	"fmt"
	"strings"
)

// ANSI Color Codes
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
	ColorBold   = "\033[1m"
	ColorDim    = "\033[2m"
)

// Background Colors
const (
	BgRed    = "\033[41m"
	BgGreen  = "\033[42m"
	BgYellow = "\033[43m"
	BgBlue   = "\033[44m"
	BgPurple = "\033[45m"
	BgCyan   = "\033[46m"
	BgWhite  = "\033[47m"
)

// Helper function to repeat strings
func repeat(s string, count int) string {
	return strings.Repeat(s, count)
}

// Helper Functions
func Red(text string) string {
	return ColorRed + text + ColorReset
}

func Green(text string) string {
	return ColorGreen + text + ColorReset
}

func Yellow(text string) string {
	return ColorYellow + text + ColorReset
}

func Blue(text string) string {
	return ColorBlue + text + ColorReset
}

func Purple(text string) string {
	return ColorPurple + text + ColorReset
}

func Cyan(text string) string {
	return ColorCyan + text + ColorReset
}

func Bold(text string) string {
	return ColorBold + text + ColorReset
}

func Dim(text string) string {
	return ColorDim + text + ColorReset
}

// Success, Warning, Error functions
func SuccessText(text string) string {
	return ColorGreen + "✅ " + text + ColorReset
}

func WarningText(text string) string {
	return ColorYellow + "⚠️  " + text + ColorReset
}

func ErrorText(text string) string {
	return ColorRed + "❌ " + text + ColorReset
}

func InfoText(text string) string {
	return ColorBlue + "ℹ️  " + text + ColorReset
}

// Enhanced formatting
func Header(text string) string {
	return ColorBold + ColorCyan + text + ColorReset
}

func Subtitle(text string) string {
	return ColorBold + ColorYellow + text + ColorReset
}

func Code(text string) string {
	return BgBlue + ColorWhite + " " + text + " " + ColorReset
}

// Example usage function
func ColorExamples() {
	fmt.Println(Header("🎨 Color Examples"))
	fmt.Println(repeat("=", 50))

	fmt.Println(Red("This is red text"))
	fmt.Println(Green("This is green text"))
	fmt.Println(Yellow("This is yellow text"))
	fmt.Println(Blue("This is blue text"))
	fmt.Println(Purple("This is purple text"))
	fmt.Println(Cyan("This is cyan text"))
	fmt.Println(Bold("This is bold text"))
	fmt.Println(Dim("This is dim text"))

	fmt.Println("\n" + Subtitle("Status Messages:"))
	fmt.Println(SuccessText("Operation completed successfully!"))
	fmt.Println(WarningText("This is a warning message"))
	fmt.Println(ErrorText("This is an error message"))
	fmt.Println(InfoText("This is an info message"))

	fmt.Println("\n" + Subtitle("Code Examples:"))
	fmt.Println("Variable:", Code("myVariable"))
	fmt.Println("Function:", Code("func main()"))
}
