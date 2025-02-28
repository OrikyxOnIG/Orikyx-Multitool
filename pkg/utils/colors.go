package utils

import "fmt"

// SetRed sets the color to red for CLI output
func SetRed(message string) string {
	return fmt.Sprintf("\033[31m%s\033[0m", message)
}

// SetGreen sets the color to green for CLI output
func SetGreen(message string) string {
	return fmt.Sprintf("\033[32m%s\033[0m", message)
}

// SetYellow sets the color to yellow for CLI output
func SetYellow(message string) string {
	return fmt.Sprintf("\033[33m%s\033[0m", message)
}
