package linux

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/keloran/go-dotfiles/console"
)

// GUI ...
func (s Settings) GUI() error {
	console.Start("Linux GUI Apps")

	err := s.createLocalBinFolder()
	if err != nil {
		return fmt.Errorf("gui: %w", err)
	}

	err = s.jetBrains()
	if err != nil {
		return fmt.Errorf("gui: %w", err)
	}

	//err = s.vscode()
	//if err != nil {
	//	return fmt.Errorf("gui: %w", err)
	//}

	console.End("Linux GUI Apps")

	return nil
}

func (s Settings) createLocalBinFolder() error {
	if !s.Files.UserFileExists(".local/bin") {
		err := s.Files.MakeUserFolder(".local")
		if err != nil {
			return fmt.Errorf("local dir: %w", err)
		}

		err = s.Files.MakeUserFolder(".local/bin")
		if err != nil {
			return fmt.Errorf("local bin dir: %w", err)
		}
	}

	return nil
}

func (s Settings) jetBrains() error {
	if !s.Files.UserFileExists(".local/bin/jetbrains-toolbox") || s.Force {
		console.Installing("jetbrains-toolbox")

		version := "jetbrains-toolbox-1.15.5796"

		err := s.Files.GetFile("https://download.jetbrains.com/toolbox", fmt.Sprintf("%s.tar.gz", version), "jetbrains.tar.gz")
		if err != nil {
			return fmt.Errorf("jetbrains-toolbox >> getfile: %w", err)
		}

		err = exec.Command("tar", "-C", s.Prefix, "-zxf", fmt.Sprintf("%s/jetbrains.tar.gz", s.Prefix)).Run()
		if err != nil {
			return fmt.Errorf("jetbrains-toolbox >> unzip: %w", err)
		}

		err = exec.Command("rm", "-f", fmt.Sprintf("%s/jetbrains.tar.gz", s.Prefix)).Run()
		if err != nil {
			return fmt.Errorf("jetbrains-toolbox >> delete zip: %w", err)
		}

		err = s.Files.MoveFile(fmt.Sprintf("%s/%s/jetbrains-toolbox", s.Prefix, version), fmt.Sprintf("%s/.local/bin/jetbrains-toolbox", s.Prefix))
		if err != nil {
			return fmt.Errorf("jetbrains-toolbox >> move toolbox: %w", err)
		}

		err = exec.Command("rm", "-rf", fmt.Sprintf("%s/%s", s.Prefix, version)).Run()
		if err != nil {
			return fmt.Errorf("jetbrains-toolbox >> delete unzipped folder: %w", err)
		}

		err = s.Files.SetPerms(".local/bin")
		if err != nil {
			return fmt.Errorf("jetbrains-toolbox >> setperms: %w", err)
		}

		console.Installed("jetbrains-toolbox")
		return nil
	}

	console.Skipping("jetbrains-toolbox")
	return nil
}

func (s Settings) atom() error {
	if !s.Files.FileExists("/usr/sbin/apm") || s.Force {
		console.Installing("atom")

		err := exec.Command("yum", "-y", "install", "atom").Run()
		if err != nil {
			return fmt.Errorf("atom >> install: %w", err)
		}

		err = s.apm()
		if err != nil {
			return fmt.Errorf("atom >> apm: %w", err)
		}

		console.Installed("atom")
		return nil
	}

	console.Skipping("atom")
	return nil
}

func (s Settings) apm() error {
	console.Installing("apm packages")

	packages := []string{
		"atom-beautify",
		"atom-clock",
		"busy-signal",
		"docker",
		"editorconfig",
		"fonts",
		"formatter-gofmt",
		"hyperclick",
		"ide-golang",
		"intentions",
		"language-docker",
		"language-gradle",
		"language-groovy",
		"linter",
		"linter-golinter",
		"linter-swagger",
		"linter-ui-default",
		"platformio-ide-terminal",
		"ssh-config",
	}

	for _, pack := range packages {
		err := exec.Command("apm", "install", pack).Run()
		if err != nil {
			return fmt.Errorf("apm >> package %s: %w", pack, err)
		}
	}

	console.Installed("apm packages")
	return nil
}

func (s Settings) vscode() error {
	if !s.Files.FileExists("/usr/bin/code") || s.Force {
		console.Installing("vscode")

		f, err := os.Create("/etc/yum.repos.d/vscode.repo")
		if err != nil {
			return fmt.Errorf("vscode repo create err: %w", err)
		}

		_, err = f.WriteString("[code]\n" +
			"name=Visual Studio Code\n" +
			"baseurl=https://packages.microsoft.com/yumrepos/vscode\n" +
			"enabled=1\n" +
			"gpgcheck=1\n" +
			"gpgkey=https://packages.microsoft.com/keys/microsoft.asc\n")
		if err != nil {
			return fmt.Errorf("vscode write repo err: %w", err)
		}

		err = f.Close()
		if err != nil {
			return fmt.Errorf("vscode repo file close: %w", err)
		}

		err = exec.Command("rpm", "--import", "https://packages.microsoft.com/keys/microsoft.asc").Run()
		if err != nil {
			return fmt.Errorf("vscode >> import key: %w", err)
		}

		err = exec.Command("dnf", "-y", "install", "code").Run()
		if err != nil {
			return fmt.Errorf("vscode >> install code: %w", err)
		}

		err = s.codeExt()
		if err != nil {
			return fmt.Errorf("vscode >> extension: %w", err)
		}

		console.Installed("vscode")
		return nil
	}

	console.Skipping("vscode")
	return nil
}

func (s Settings) codeExt() error {
	console.Installing("vscode extensions")

	exts := []string{
		"42crunch.vscode-openapi",
		"azemoh.one-monokai",
		"compulim.vscode-clock",
		"editorconfig.editorconfig",
		"formulahendry.terminal",
		"hookyqr.beautify",
		"maxmedia.go-prof",
		"ms-azuretools.vscode-docker",
		"ms-kubernetes-tools.vscode-kubernetes-tools",
		"ms-vscode.cpptools",
		"ms-vscode.csharp",
		"ms-vscode.go",
		"qinjina.seti-icons",
	}

	for _, ext := range exts {
		err := exec.Command("code", "--install-extension", ext, "--user-data-dir", s.Prefix).Run()
		if err != nil {
			return fmt.Errorf("vscode extension >> ext: %s, %w", ext, err)
		}
	}

	console.Installed("vscode extensions")
	return nil
}
