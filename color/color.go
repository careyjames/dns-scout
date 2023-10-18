package color

import "fmt"

// Yellow color
func Yellow(text string) string {
	return fmt.Sprintf("\033[38;5;222m%s\033[0m", text)
}

// Green color
func Green(text string) string {
	return fmt.Sprintf("\033[38;5;78m%s\033[0m", text)
}

// Blue color
func Blue(text string) string {
	return fmt.Sprintf("\033[38;5;39m%s\033[0m", text)
}

// Red color
func Red(text string) string {
	return fmt.Sprintf("\033[38;5;88m%s\033[0m", text)
}
