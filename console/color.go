package console

import "fmt"

// Blue ...
func Blue(text string) string {
	return console(text, 34)
}

// Green ...
func Green(text string) string {
	return console(text, 32)
}

// Red ...
func Red(text string) string {
	return console(text, 31)
}

// Black ...
func Black(text string) string {
	return console(text, 30)
}

// Yellow ...
func Yellow(text string) string {
	return console(text, 33)
}

// Magenta ...
func Magenta(text string) string {
	return console(text, 35)
}

// Cyan ...
func Cyan(text string) string {
	return console(text, 36)
}

// White ...
func White(text string) string {
	return console(text, 37)
}

func console(text string, color int) string {
	return fmt.Sprintf("\u001b[%dm%s%s", color, text, reset())
}

func reset() string {
	return fmt.Sprintf("\u001b[0m")
}
