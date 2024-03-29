package color

import "fmt"

// Yellow color
func Yellow(text string) string {
	return fmt.Sprintf("\033[38;5;3m%s\033[0m", text)
}

// Light Grey 242
func Grey(text string) string {
	return fmt.Sprintf("\033[38;5;242m%s\033[0m", text)
}

// Blue color
func Blue(text string) string {
	return fmt.Sprintf("\033[38;5;4m%s\033[0m", text)
}

// Red color
func Red(text string) string {
	return fmt.Sprintf("\033[38;5;1m%s\033[0m", text)
}
