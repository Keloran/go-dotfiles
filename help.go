package main

import (
	"fmt"

	"github.com/keloran/go-dotfiles/console"
)

func displayHelp() {
	console.Start("Help")
	console.Info("Usage dotfiles <command> <flags>")
	console.BlankLine()
	console.Info("Commands:")
	helpCommand("help", "  This message")
	helpCommand("cli", "   Install CLI Apps")
	helpCommand("gui", "   Install GUI Apps")
	helpCommand("os", "    Install OS Specific settings")
	helpCommand("update", "Run the OS Specific updater (e.g. brew)")
	helpCommand("dots", "  Install only the .files")
	helpCommand("all", "   Do everything")
	console.BlankLine()
	console.Info("Flags:")
	helpFlag("github", "Get files from a github source")
	helpFlag("force", " Force override of already installed files / apps")
	helpFlag("user", "  Username if required")
}

func helpCommand(flag string, desc string) {
	fmt.Println(fmt.Sprintf("    %s        %s", console.Cyan(flag), console.Cyan(desc)))
}

func helpFlag(flag string, desc string) {
	fmt.Println(console.Cyan(fmt.Sprintf("    -%s        %s", flag, desc)))
}
