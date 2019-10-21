package dots

import (
	"fmt"
	"os"

	"github.com/keloran/go-dotfiles/console"
)

func (d Dots) colorDiffRC() error {
	if !d.Files.UserFileExists(".colordiffrc") || d.Force {
		if d.Force {
			console.Warning("Forced .colordiffrc")
		}

		if d.Github {
			return d.Files.GetGithubFile(".colordiffrc")
		}

		console.Nice("Creating .colordiffrc")
		return d.createColorDiffRC()
	}

	console.Info("Skipped .colordiffrc")
	return nil
}

func (d Dots) createColorDiffRC() error {
	f, err := os.Create(fmt.Sprintf("%s/.colordiffrc", d.Prefix))
	if err != nil {
		return fmt.Errorf("ColorDiffRC create err: %w", err)
	}

	_, err = f.WriteString("# ColorDiffRC\n" +
		"newtext = green\n" +
		"oldtext = red\n" +
		"diffstuff = cyan\n")
	if err != nil {
		return fmt.Errorf("ColorDiffRC: %w", err)
	}

	err = f.Close()
	if err != nil {
		return fmt.Errorf("colordiffrc file close: %w", err)
	}

	err = d.Files.SetUserPerm(".colordiffrc")
	if err != nil {
		return fmt.Errorf("colordiffrc: %w", err)
	}

	return nil
}
