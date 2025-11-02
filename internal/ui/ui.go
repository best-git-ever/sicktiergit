package ui

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

var (
	colorGreen  = color.New(color.FgGreen).SprintFunc()
	colorYellow = color.New(color.FgYellow).SprintFunc()
	colorRed    = color.New(color.FgRed).SprintFunc()
	colorCyan   = color.New(color.FgCyan).SprintFunc()
	colorDim    = color.New(color.Faint).SprintFunc()
)

// PrintSuccess prints a success message in green
func PrintSuccess(msg string) {
	fmt.Println(colorGreen(msg))
}

// PrintWarning prints a warning message in yellow
func PrintWarning(msg string) {
	fmt.Println(colorYellow(msg))
}

// PrintError prints an error message in red
func PrintError(msg string) {
	fmt.Println(colorRed(msg))
}

// PrintInfo prints an info message in cyan
func PrintInfo(msg string) {
	fmt.Println(colorCyan(msg))
}

// PrintDim prints a dimmed message
func PrintDim(msg string) {
	fmt.Println(colorDim(msg))
}

// ColorGreen returns a green colored string
func ColorGreen(s string) string {
	return colorGreen(s)
}

// ColorYellow returns a yellow colored string
func ColorYellow(s string) string {
	return colorYellow(s)
}

// ShowProgress displays a progress indicator
func ShowProgress(seconds int) {
	chars := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	iterations := seconds * 10

	for i := 0; i < iterations; i++ {
		fmt.Printf("\r%s ", chars[i%len(chars)])
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Print("\r  \r") // Clear the spinner
}

// GenerateHash generates a random git-like hash
func GenerateHash(length int) string {
	const charset = "0123456789abcdef"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// BoolToStatus converts a boolean to an enabled/disabled string
func BoolToStatus(b bool) string {
	if b {
		return colorGreen("enabled")
	}
	return colorRed("disabled")
}
