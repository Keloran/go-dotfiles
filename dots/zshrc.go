package dots

import (
	"fmt"
	"os"

	"github.com/keloran/go-dotfiles/console"
)

func (d Dots) zshrc() error {
	if !d.Files.UserFileExists(".zshrc") || d.Force {
		if d.Force {
			console.Info("Forced .zshrc")
		}

		if d.Github {
			return d.Files.GetGithubFile(".zshrc")
		}

		console.Nice("Creating .zshrc")
		err := d.createZSHRC()
		if err != nil {
			return fmt.Errorf("ZSHRC err: %w", err)
		}

		if !d.Files.UserFileExists(".oh-my-zsh/themes/keloran.zsh-theme") || d.Force {
			return d.getZSHTheme()
		}
	}

	console.Info("Skipped .zshrc")
	return nil
}

func (d Dots) createZSHRC() error {
	f, err := os.Create(fmt.Sprintf("%/.zshrc", d.Prefix))
	if err != nil {
		return fmt.Errorf("create ZSHRC err: %w", err)
	}

	_, err = f.WriteString("# ZSH Install\n" +
		"export ZSH=~/.oh-my-zsh\n\n" +
		"# Z\n" +
		". /usr/local/etc/profile.d/z.sh\n\n" +
		"# Theme" +
		"ZSH_THEME=\"keloran\"\n\n" +
		"# Case-Sensitive Completion\n" +
		"CASE_SENSITIVE=true\n\n" +
		"# Plugins\n" +
		"plugins=(gitfast osx sudo common-aliases z brew docker aws zsh_reload)\n\n" +
		"# ZSH\n" +
		"source $ZSH/oh-my-zsh.sh\n\n" +
		"# Rust Cargo\n" +
		"if [[ -z $HOME/.cargo ]]; then\n" +
		"  source $HOME/.cargo/env\n" +
		"fi\n\n" +
		"# Language\n" +
		"export LANG=en_GB.UTF-8\n\n" +
		"# GPG\n" +
		"if [[ -e ~/.gnupg/S.gpg-agent ]]; then\n" +
		"  export GPG_AGENT_INFO\n" +
		"else\n" +
		"  eval $(gpg-agent --daemon)\n" +
		"fi\n\n" +
		"# Aliases\n" +
		"if [[ -f ~/.aliases ]]; then\n" +
		"  source ~/.aliases\n" +
		"fi\n\n" +
		"# Functions\n" +
		"if [[ -f ~/.functions ]]; then\n" +
		"  source ~/.functions\n" +
		"fi\n\n" +
		"# Go\n" +
		"export GO_ENV=~/.goenvs\n" +
		"export GOPATH=$(go env GOPATH)\n" +
		"GOBINS=$GOPATH/bin\n" +
		"export GPG_TTY=$(tty)\n\n" +
		"# Path" +
		"export PATH=\"/usr/local/bin:/usr/bin:/usr/sbin:/sbin:/opt/X11/bin:/usr/local/opt/go/libexec/bin:/usr/local/opt/python/libexec/bin:$GOBINS:$PATH\"\n\n" +
		"# ZPlug\n" +
		"source /usr/local/opt/zplug/init.zsh\n\n" +
		"zplug \"kingsj/atom_plugin.zsh\"\n" +
		"zplug \"bric3/nice-exit-code\"\n" +
		"zplug \"bbenne10/goenv\"\n" +
		"zplug \"chrissicool/zsh-256color\"\n" +
		"zplug \"desyncr/auto-ls\"\n" +
		"zplug \"mafredri/zsh-async\"\n" +
		"zplug \"supercrabtree/k\"\n" +
		"zplug \"eendroroy/zed-zsh\"\n" +
		"zplug \"zsh-users/zsh-apple-touchbar\"\n\n" +
		"zplug load\n\n" +
		"# Travis\n" +
		"if [[ -f ~/.travis/travis.sh ]]; then\n" +
		"  source ~/.travis/travis.sh\n" +
		"fi\n")
	if err != nil {
		return fmt.Errorf("ZSHRC: %w", err)
	}

	err = f.Close()
	if err != nil {
		return fmt.Errorf("zshrc file close: %w", err)
	}

	err = d.Files.SetUserPerm(".zshrc")
	if err != nil {
		return fmt.Errorf(".zshrc: %w", err)
	}

	return nil
}

func (d Dots) getZSHTheme() error {
	console.Info("Get ZSH Theme")
	err := d.Files.GetFile("https://raw.githubusercontent.com/Keloran/keloran.zsh-theme/master/", "keloran.zsh-theme")
	if err != nil {
		return fmt.Errorf("getZSHTheme file: %w", err)
	}

	err = os.Rename(fmt.Sprintf("%s/%s", d.Prefix, "keloran.zsh-theme"), fmt.Sprintf("%s/.oh-my-zsh/themes/%s", d.Prefix, "keloran.zsh-theme"))
	if err != nil {
		return fmt.Errorf("move zsh theme to themes folder: %w", err)
	}

	return nil
}
