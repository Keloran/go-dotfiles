package osx

import "github.com/keloran/go-dotfiles/files"

// Settings ...
type Settings struct {
	Username string
	Prefix   string

	GoVersion string
	Force     bool

	Files files.Files
}
