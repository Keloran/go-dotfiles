package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/keloran/go-dotfiles/apps"
	"github.com/keloran/go-dotfiles/console"
	"github.com/keloran/go-dotfiles/dots"
	"github.com/keloran/go-dotfiles/osspecific"
)

// USERNAME constant to send between settings
var USERNAME = ""

func main() {
	if runtime.GOOS == "windows" {
		console.Error("Windows not yet supported")
		return
	}

	command := ""
	if len(os.Args) >= 2 {
		for _, arg := range os.Args {
			if strings.Contains(arg, "-") {
				continue
			}

			command = arg
		}
	}

	var dotfiles, appsGUI, appsCLI, OS, all bool
	help := true

	switch command {
	case "dots":
		dotfiles = true
	case "gui":
		appsGUI = true
	case "cli":
		appsCLI = true
	case "all":
		all = true
	case "help":
		help = true
	}

	// flags
	github := flag.Bool("github", false, "")
	force := flag.Bool("force", false, "")
	flag.StringVar(&USERNAME, "user", "", "")

	// set the flags to parse
	flag.Parse()

	// Linux specifics
	if runtime.GOOS == "linux" {
		if USERNAME == "" && command != "" {
			err := errors.New("")
			USERNAME, err = console.Question("Need your linux username (e.g. bobsmith)", false)
			if err != nil {
				console.Error(fmt.Sprintf("username err: %v", err))
			}
		}

		if os.Geteuid() != 0 {
			console.Error("You have to run this as root")
			return
		}
	}

	// actually do stuff
	Dots := dots.Dots{
		Username: USERNAME,
		Force:    *force,
		Github:   *github,
	}
	if dotfiles {
		err := Dots.Install()
		if err != nil {
			console.Error(fmt.Sprintf(".Files err: %v", err))
		}
		return
	}

	Apps := apps.Apps{
		Username: USERNAME,
		Prefix:   fmt.Sprintf("/Users/%s", USERNAME),
		Force:    *force,
	}
	if appsGUI {
		err := Apps.GUI()
		if err != nil {
			console.Error(fmt.Sprintf("GUI Apps err: %v", err))
		}
		return
	}

	if appsCLI {
		err := Apps.CLI()
		if err != nil {
			console.Error(fmt.Sprintf("CLI Apps err: %v", err))
		}
		return
	}

	OSS := osspecific.OSSpecific{
		Username: USERNAME,
		Force:    *force,
	}
	if OS {
		err := OSS.Install()
		if err != nil {
			console.Error(fmt.Sprintf("OS Specific err: %v", err))
		}
		return
	}

	if all {
		err := Dots.Install()
		if err != nil {
			console.Error(fmt.Sprintf(".Files err: %v", err))
		}
		err = OSS.Install()
		if err != nil {
			console.Error(fmt.Sprintf("OS Specific err: %v", err))
		}
		err = Apps.GUI()
		if err != nil {
			console.Error(fmt.Sprintf("GUI Apps err: %v", err))
		}
		err = Apps.CLI()
		if err != nil {
			console.Error(fmt.Sprintf("CLI Apps err: %v", err))
		}
		help = false
	}

	// help is last because been triggered
	if help {
		displayHelp()
		return
	}
}
