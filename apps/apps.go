package apps

import (
	"fmt"
	"runtime"

	"github.com/keloran/go-dotfiles/apps/linux"
	"github.com/keloran/go-dotfiles/apps/osx"
	"github.com/keloran/go-dotfiles/files"
)

// Apps ...
type Apps struct {
	Username string
	Prefix   string
	Force    bool
}

// GUI ...
func (a Apps) GUI() error {
	switch runtime.GOOS {
	case "linux":
		prefix := fmt.Sprintf("/home/%s", a.Username)

		return linux.Settings{
			Username: a.Username,
			Prefix:   prefix,
			Files: files.Files{
				Username: a.Username,
				Prefix:   prefix,
			},
			Force: a.Force,
		}.GUI()
	case "darwin":
		return osx.Settings{
			Username: a.Username,
			Prefix:   a.Prefix,
		}.GUI()
	}

	return nil
}

// CLI ...
func (a Apps) CLI() error {
	switch runtime.GOOS {
	case "linux":
		prefix := fmt.Sprintf("/home/%s", a.Username)

		return linux.Settings{
			Username: a.Username,
			Prefix:   prefix,
			Files: files.Files{
				Username: a.Username,
				Prefix:   prefix,
			},
			Force: a.Force,
		}.CLI()
	case "darwin":
		return osx.Settings{
			Username: a.Username,
			Prefix:   a.Prefix,
		}.CLI()
	}

	return nil
}
