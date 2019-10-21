package dots

import (
	"fmt"
	"os"

	"github.com/keloran/go-dotfiles/console"
)

func (d Dots) multitailRC() error {
	if !d.Files.UserFileExists(".multitailrc") || d.Force {
		if d.Force {
			console.Warning("Forced .multitailrc")
		}

		if d.Github {
			return d.Files.GetGithubFile(".multitailrc")
		}

		console.Nice("Creating .multitailrc")
		return d.createMultitailRC()
	}

	console.Info("Skipped .multitailrc")
	return nil
}

func (d Dots) createMultitailRC() error {
	f, err := os.Create(fmt.Sprintf("%s/.multitailrc", d.Prefix))
	if err != nil {
		return fmt.Errorf("MultitailRC create err: %w", err)
	}

	_, err = f.WriteString("# MultitailRC\n" +
		"check_main:0\n")
	if err != nil {
		return fmt.Errorf("MultitailRC: %w", err)
	}

	err = f.Close()
	if err != nil {
		return fmt.Errorf("multitailrc file close: %w", err)
	}

	err = d.Files.SetUserPerm(".multitailrc")
	if err != nil {
		return fmt.Errorf("multitailrc: %w", err)
	}

	return nil
}
