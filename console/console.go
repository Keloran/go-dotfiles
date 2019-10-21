package console

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Start ...
func Start(text string) {
	fmt.Println(Green(fmt.Sprintf("--- Starting %s ---", text)))
}

// End ...
func End(text string) {
	fmt.Println(Green(fmt.Sprintf("*** Finished %s ***", text)))
}

// Warning ...
func Warning(text string) {
	fmt.Println(Blue(fmt.Sprintf("+++ %s +++", text)))
}

// Info ...
func Info(text string) {
	fmt.Println(Cyan(fmt.Sprintf("%s", text)))
}

// Error .
func Error(text string) {
	fmt.Println(Red(fmt.Sprintf("### %s ###", text)))
}

// Log ...
func Log(text string) {
	fmt.Println(White(fmt.Sprintf("%s", text)))
}

// Nice ...
func Nice(text string) {
	fmt.Println(Yellow(fmt.Sprintf("%s", text)))
}

// Question ...
func Question(text string, lower bool) (string, error) {
	fmt.Println(Magenta(fmt.Sprintf("%s ?", text)))
	fmt.Print(White("~> "))

	return getResponse(lower)
}

// Optional ...
func Optional(text string, lower bool) (string, error) {
	fmt.Println(fmt.Sprintf("%s ? (optional press enter to leave blank)", Magenta(text)))
	fmt.Print(White("~> "))

	return getResponse(lower)
}

func getResponse(lower bool) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	in, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("getResponse %s err : %w", err)
	}
	in = strings.Replace(in, "\n", "", -1)

	if lower {
		in = strings.ToLower(in)
	}

	return in, nil
}

// BlankLine ...
func BlankLine() {
	fmt.Println("")
}

// Skipping ...
func Skipping(text string) {
	Nice(fmt.Sprintf("    Skipping %s", text))
}

// Installing ...
func Installing(text string) {
	fmt.Print(Cyan(fmt.Sprintf(">>> Installing %s ---", text)))
}

// Installed ...
func Installed(text string) {
	Info(fmt.Sprintf(" %s Installed <<<", text))
}
