package osspecific

import (
	"errors"
	"fmt"
	"runtime"

	"github.com/keloran/go-dotfiles/files"
)

// OSSpecific ...
type OSSpecific struct {
	Username string
	Prefix   string
	Force    bool
	Files    files.Files
}

// Install ...
func (o OSSpecific) Install() error {
	o.Prefix = "~"
	if o.Username != "" {
		switch runtime.GOOS {
		case "darwin":
			o.Prefix = fmt.Sprintf("/Users/%s", o.Username)
		case "linux":
			o.Prefix = fmt.Sprintf("/home/%s", o.Username)
		}
	}

	o.Files = files.Files{
		Username: o.Username,
		Prefix:   o.Prefix,
	}

	switch runtime.GOOS {
	case "linux":
		return o.linuxInstall()
	case "darwin":
		return o.macInstall()
	}

	return errors.New("only mac and linux supported")
}
