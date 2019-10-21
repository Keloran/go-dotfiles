package osspecific

import (
	"fmt"
	"os/exec"

	"github.com/keloran/go-dotfiles/console"
)

func (o OSSpecific) linuxInstall() error {
	console.Start("Linux Specific")

	err := o.ssh()
	if err != nil {
		return fmt.Errorf("linux: %w", err)
	}

	console.End("Linux Specific")

	return nil
}

func (o OSSpecific) ssh() error {
	if o.Force || !o.Files.UserFileExists(".ssh/id_rsa") {
		console.Installing("ssh key")

		err := exec.Command("ssh-keygen", "-b", "2048", "-t", "rsa", "-f", fmt.Sprintf("%s/.ssh/id_rsa", o.Prefix), "-q", "-N", `""`).Run()
		if err != nil {
			return fmt.Errorf("ssh >> keygen: %w", err)
		}

		err = o.Files.SetPerms(".ssh")
		if err != nil {
			return fmt.Errorf("ssh >> initial folder perms: %w", err)
		}

		err = o.Files.SetUserPerm(".ssh/id_rsa.pub")
		if err != nil {
			return fmt.Errorf("ssh >> initial pub perm: %w", err)
		}

		err = o.Files.SetUserPerm(".ssh/id_rsa")
		if err != nil {
			return fmt.Errorf("ssh >> initial key perm: %w", err)
		}

		err = exec.Command("chmod", "700", fmt.Sprintf("%s/.ssh", o.Prefix)).Run()
		if err != nil {
			return fmt.Errorf("ssh >> folder perm: %w", err)
		}

		err = exec.Command("chmod", "644", fmt.Sprintf("%s/.ssh/id_rsa.pub", o.Prefix)).Run()
		if err != nil {
			return fmt.Errorf("ssh >> pub perm: %w", err)
		}

		err = exec.Command("chmod", "600", fmt.Sprintf("%s/.ssh/id_rsa", o.Prefix)).Run()
		if err != nil {
			return fmt.Errorf("ssh >> key perm: %w", err)
		}

		console.Installed("ssh key")
		return nil
	}

	console.Skipping("ssh key")
	return nil
}
