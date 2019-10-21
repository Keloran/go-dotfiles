package dots

import (
	"fmt"
	"runtime"

	"github.com/keloran/go-dotfiles/console"
	"github.com/keloran/go-dotfiles/files"
)

// Dots ...
type Dots struct {
	Username string
	Force    bool
	Github   bool
	Prefix   string

	GitName    string
	GitEmail   string
	GitImage   string
	GithubName string

	Files files.Files
}

// Install ...
func (d Dots) Install() error {
	console.Start("DotFiles")

	d.Prefix = "~"
	if d.Username != "" {
		switch runtime.GOOS {
		case "darwin":
			d.Prefix = fmt.Sprintf("/Users/%s", d.Username)
		case "linux":
			d.Prefix = fmt.Sprintf("/home/%s", d.Username)
		}
	}

	d.Files = files.Files{
		Username: d.Username,
		Prefix:   d.Prefix,
	}

	err := d.gitConfig()
	if err != nil {
		return fmt.Errorf("gitconfig err: %w", err)
	}
	err = d.aliases()
	if err != nil {
		return fmt.Errorf("aliases err: %w", err)
	}
	err = d.functions()
	if err != nil {
		return fmt.Errorf("functions err: %w", err)
	}
	err = d.editorConfig()
	if err != nil {
		return fmt.Errorf("editorconfig err: %w", err)
	}
	err = d.gitIgnore()
	if err != nil {
		return fmt.Errorf("getignore err: %w", err)
	}
	err = d.multitailRC()
	if err != nil {
		return fmt.Errorf("multitailrc err: %w", err)
	}
	err = d.hyperJS()
	if err != nil {
		return fmt.Errorf("hyperjs err: %w", err)
	}

	console.End("DotFiles")

	return nil
}

func (d Dots) gitIgnore() error {
	return files.Files{
		Username: d.Username,
		Prefix:   d.Prefix,
	}.GetGithubFile(".gitignore_global")
}
