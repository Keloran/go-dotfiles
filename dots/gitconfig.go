package dots

import (
	"errors"
	"fmt"
	"os"
	"runtime"

	"github.com/keloran/go-dotfiles/console"
)

func (d Dots) gitConfig() error {
	if !d.Files.UserFileExists(".gitconfig") || d.Force {
		if d.Force {
			console.Warning("Forced .gitconfig")
		}

		if d.Github {
			return d.Files.GetGithubFile(".gitconfig")
		}

		d, err := d.getGitConfigSettings()
		if err != nil {
			return fmt.Errorf("gitconfig settings err: %w", err)
		}

		console.Nice("Creating .gitconfig")
		return d.createGitConfig()
	}

	console.Info("Skipped .gitconfig")
	return nil
}

func (d Dots) getGitConfigSettings() (Dots, error) {
	// Inputs
	err := errors.New("")
	console.Info("GitConfig Setup")
	d.GitName, err = console.Question("Name to attach to commit message (e.g. Bob Smith)", false)
	if err != nil {
		return d, fmt.Errorf("gitconfig >> name: %w", err)
	}

	d.GitEmail, err = console.Question("Email for commit messages (e.g. bob.smith@email.com)", false)
	if err != nil {
		return d, fmt.Errorf("gitconfig >> email: %w", err)
	}

	d.GitImage, err = console.Optional("Image to attach to commit name", false)
	if err != nil {
		return d, fmt.Errorf("gitconfig >> image: %w", err)
	}

	d.GithubName, err = console.Optional("Github Username", false)
	if err != nil {
		return d, fmt.Errorf("gitconfig >> github username: %w", err)
	}

	return d, nil
}

func (d Dots) createGitConfig() error {
	f, err := os.Create(fmt.Sprintf("%s/.gitconfig", d.Prefix))
	if err != nil {
		return fmt.Errorf("GitConfig create err: %w", err)
	}

	_, err = f.WriteString("# GitConfig\n" +
		"[user]\n")
	if err != nil {
		return fmt.Errorf("failed to write gitconfig header: %w", err)
	}

	if d.GitName != "" {
		_, err = f.WriteString(fmt.Sprintf("  name = %s\n", d.GitName))
		if err != nil {
			return fmt.Errorf("failed to write gitconfig name: %w", err)
		}
	}

	if d.GitImage != "" {
		_, err = f.WriteString(fmt.Sprintf("  image = %s\n", d.GitImage))
		if err != nil {
			return fmt.Errorf("failed to write gitconfig image: %w", err)
		}
	}

	if d.GitEmail != "" {
		_, err = f.WriteString(fmt.Sprintf("  email = %s\n", d.GitEmail))
		if err != nil {
			return fmt.Errorf("failed to write gitconfig email: %w", err)
		}
	}

	_, err = f.WriteString("[push]\n" +
		"  default = simple\n" +
		"[core]\n" +
		"  editor = /usr/bin/vim\n" +
		"  excludesfile = ~/.gitignore_global\n")
	if err != nil {
		return fmt.Errorf("failed to write gitconfig core: %w", err)
	}

	if d.GithubName != "" {
		_, err = f.WriteString("[github]\n" +
			fmt.Sprintf("  user = %s\n", d.GithubName))
		if err != nil {
			return fmt.Errorf("failed to write gitconfig github: %w", err)
		}
	}

	_, err = f.WriteString("[alias]\n" +
		"  a = add --all\n" +
		"  b = branch\n" +
		"  c = commit\n" +
		"  co = checkout\n" +
		"  rh = reset --hard\n" +
		"  s = status -sb\n" +
		"  am = add --all -m\n" +
		"  logs = !\"git log --oneline --decorate --all --graph --stat\"\n" +
		"  standup = !\"git log --reverse --branches --since=$(if [[ \"Mon\" == \"$(date +%a)\" ]]; then echo \"last friday\"; else echo \"yesterday\"; fi) --author=$(git config --get user.email) --format=format:'%C(cyan) %ad %C(yellow)%h %Creset %s %Cgreen%d' --date=local\"\n" +
		"  pushit = !\"git push -u origin $(git rev-parse --abbrev-ref HEAD)\"\n" +
		"  ticket = !\"git checkout -b $2\"\n" +
		"  master = !\"git checkout master && git pull\"\n" +
		"  opensource = \"!f() {\\n" +
		"    git fetch upstream;\\n" +
		"    git merge upstream/master;\\n" +
		"    git rebase;\\n" +
		"    git push;\\n" +
		"  };f\"\n")
	if err != nil {
		return fmt.Errorf("GitConfig: %w", err)
	}

	if runtime.GOOS == "darwin" {
		_, err = f.WriteString("[credential]\n" +
			"  helper = osxkeychain\n")
		if err != nil {
			return fmt.Errorf("GitConfig: %w", err)
		}
	}

	_, err = f.WriteString("[console]\n" +
		"  branch = auto\n" +
		"  diff = auto\n" +
		"  status = auto\n" +
		"[push]\n" +
		"  default = upstream\n" +
		"[commit]\n" +
		"  gpgSign = true\n" +
		"[filter \"lfs\"]\n" +
		"  clean = git-lfs clean %f\n" +
		"  smudge = git-lfs smudge %f\n" +
		"  required = true\n" +
		"[gpg]\n" +
		"  program = /usr/local/bin/krgpg\n" +
		"[tag]\n" +
		"  forceSignAnnotated = true\n" +
		"[diff]\n" +
		"  tool = opendiff\n" +
		"[difftool]\n" +
		"  prompt = true\n" +
		"[difftool \"opendiff\"]\n" +
		"  cmd = /usr/bin/opendiff \\\"$LOCAL\\\" \\\"$REMOTE\\\" -merge \\\"$MERGED\\\" | cat\n" +
		"[difftool \"Kaleidoscope\"]\n" +
		"  cmd = ksdiff --partial-changeset --relative-path \\\"$MERGED\\\" -- \\\"$LOCAL\\\" \\\"$REMOTE\\\"\n" +
		"[mergetool \"Kaleidoscope\"]\n" +
		"  cmd kddiff --merge --output \\\"$MERGED\\\" --base \\\"$BASE\\\" -- \\\"$LOCAL\\\" --snapshot \\\"$REMOTE\\\" --snapshot\n" +
		"  trustexitcode = true\n" +
		"[pager]\n" +
		"  diff = diff-so-fancy | less --tabs=1,5 -RFX\n" +
		"  show = diff-so-fancy | less --tabs1,4 --RFX\n")
	if err != nil {
		return fmt.Errorf("GitConfig: %w", err)
	}

	err = f.Close()
	if err != nil {
		return fmt.Errorf("gitconfig file close: %w", err)
	}

	err = d.Files.SetUserPerm(".gitconfig")
	if err != nil {
		return fmt.Errorf("gitconfig: %w", err)
	}

	return nil
}
