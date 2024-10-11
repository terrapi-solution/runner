package helper

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"golang.org/x/term"
)

// Prints a stylized header to the console using different colors.
func PrintHeader() {
	c1 := color.New(color.FgCyan).Add(color.Bold)
	c2 := color.New(color.FgYellow).Add(color.Bold)
	c3 := color.New(color.FgGreen).Add(color.Bold)
	c4 := color.New(color.FgRed).Add(color.Bold)

	c1.Println("88888888888                                      d8888 8888888b. 8888888")
	c2.Println("    888                                         d88888 888   Y88b  888  ")
	c3.Println("    888                                        d88P888 888    888  888  ")
	c4.Println("    888   .d88b.  888d888 888d888 8888b.      d88P 888 888   d88P  888  ")
	c1.Println("    888  d8P  Y8b 888P\"   888P\"      \"88b    d88P  888 8888888P\"   888  ")
	c2.Println("    888  88888888 888     888    .d888888   d88P   888 888         888  ")
	c3.Println("    888  Y8b.     888     888    888  888  d8888888888 888         888  ")
	c4.Println("    888   \"Y8888  888     888    \"Y888888 d88P     888 888       8888888")
	fmt.Println()
}

// Prints a section header with the specified value.
func PrintSection(value string) {
	PrintLine(color.FgYellow)
	PrintBoldMessage(value, color.FgYellow)
	PrintLine(color.FgYellow)
}

// Prints a message to the console with optional color attributes.
func PrintMessage(message string, value ...color.Attribute) {
	c := color.New(value...)
	c.Println(message)
}

// Prints the given message in bold with optional color attributes.
func PrintBoldMessage(message string, value ...color.Attribute) {
	c := color.New(value...).Add(color.Bold)
	c.Println(message)
}

// Prints a bold line across the terminal width.
func PrintLine(value ...color.Attribute) {
	c := color.New(value...).Add(color.Bold)

	line := make([]rune, getTerminalSize())
	for i := range line {
		line[i] = '-'
	}
	c.Println(string(line))
}

// Returns the width of the terminal in characters.
// If it fails to get the terminal size, it defaults to 80 characters.
func getTerminalSize() int {
	screenWidth, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		// Fallback to a standard terminal width of 80 characters
		screenWidth = 80
	}
	return screenWidth
}
