package dots

import (
	"fmt"
	"os"

	"github.com/keloran/go-dotfiles/console"
	"github.com/keloran/go-dotfiles/files"
)

func (d Dots) aliases() error {
	if !d.Files.UserFileExists(".aliases") || d.Force {
		if d.Force {
			console.Warning("Forced .aliases")
		}

		if d.Github {
			return files.Files{
				Username: d.Username,
				Prefix:   d.Prefix,
			}.GetGithubFile(".aliases")
		}

		console.Nice("Creating .aliases")
		return d.createAliases()
	}

	console.Info("Skipped .aliases")
	return nil
}

func (d Dots) createAliases() error {
	f, err := os.Create(fmt.Sprintf("%s/.aliases", d.Prefix))
	if err != nil {
		return fmt.Errorf("aliases create err: %w", err)
	}

	_, err = f.WriteString("## Aliases\n" +
		"\n" +
		"alias projects=\"cd ~/Documents/Projects\"\n" +
		"alias dpsa=\"docker ps -a\"\n" +
		"alias gsub=\"git submodule update --init --recursive\"\n")
	if err != nil {
		return fmt.Errorf("aliases write string err: %w", err)
	}

	err = f.Close()
	if err != nil {
		return fmt.Errorf("aliases file close: %w", err)
	}

	err = d.Files.SetUserPerm(".aliases")
	if err != nil {
		return fmt.Errorf("aliases: %w", err)
	}

	return nil
}
