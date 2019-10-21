package linux

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/keloran/go-dotfiles/console"
)

// CLI ...
func (s Settings) CLI() error {
	console.Start("Linux CLI Apps")

	err := s.docker()
	if err != nil {
		return err
	}

	// get go version
	goVersion := "1.13.1"
	console.Info("Go Version")
	s.GoVersion, err = console.Optional(fmt.Sprintf("Which go version [%s]", goVersion), false)
	if err != nil {
		return fmt.Errorf("go version %w", err)
	}
	if s.GoVersion == "" {
		s.GoVersion = goVersion
	}

	err = s.golang()
	if err != nil {
		return fmt.Errorf("cli: %w", err)
	}

	err = s.ripgrep()
	if err != nil {
		return fmt.Errorf("cli: %w", err)
	}

	err = s.installApps()
	if err != nil {
		return fmt.Errorf("cli: %w", err)
	}

	err = s.getFonts()
	if err != nil {
		return fmt.Errorf("cli: %w", err)
	}

	err = s.diffSoFancy()
	if err != nil {
		return fmt.Errorf("cli: %w", err)
	}

	err = s.kr()
	if err != nil {
		return fmt.Errorf("cli: %w", err)
	}

	err = s.npmPackages()
	if err != nil {
		return fmt.Errorf("cli: %w", err)
	}

	console.End("Linux CLI Apps")
	return nil
}

func (s Settings) docker() error {
	if !s.Files.FileExists("/etc/yum.repos.d/docker-ce.repo") || s.Force {
		console.Installing("Docker")
		err := exec.Command("yum", "install", "-y", "epel-release").Run()
		if err != nil {
			return fmt.Errorf("docker >> enable extras repo err: %w", err)
		}

		err = exec.Command("yum", "install", "-y", "yum-utils", "device-mapper-persistent-data", "lvm2").Run()
		if err != nil {
			return fmt.Errorf("docker >> install yum config-manager err: %w", err)
		}

		err = exec.Command("yum-config-manager", "--add-repo", "https://download.docker.com/linux/centos/docker-ce.repo").Run()
		if err != nil {
			return fmt.Errorf("docker >> add repo: %w", err)
		}

		err = exec.Command("yum", "install", "-y", "docker-ce", "docker-ce-cli", "containerd.io").Run()
		if err != nil {
			return fmt.Errorf("docker >> install docker: %w", err)
		}

		err = exec.Command("systemctl", "start", "docker").Run()
		if err != nil {
			return fmt.Errorf("docker >> start docker: %w", err)
		}

		err = exec.Command("systemctl", "enable", "docker").Run()
		if err != nil {
			return fmt.Errorf("docker >> enable docker at boot: %w", err)
		}

		console.Installed("Docker")
		return nil
	}

	console.Skipping("Docker")
	return nil
}

func (s Settings) ripgrep() error {
	if !s.Files.FileExists("/etc/yum.repos.d/carlwgeorge-ripgrep-epel-7.repo") || s.Force {
		console.Installing("ripgrep")

		err := exec.Command("yum-config-manager", "--add-repo", "https://copr.fedorainfracloud.org/coprs/carlwgeorge/ripgrep/repo/epel-7/carlwgeorge-ripgrep-epel-7.repo").Run()
		if err != nil {
			return fmt.Errorf("ripgrep >> repo: %w", err)
		}

		err = exec.Command("yum", "install", "-y", "ripgrep").Run()
		if err != nil {
			return fmt.Errorf("ripgrep >> install: %w", err)
		}

		console.Installed("ripgrep")
		return nil
	}

	console.Skipping("ripgrep")
	return nil
}

func (s Settings) golang() error {
	overRide := false

	if s.Force {
		overRide = true
	}

	if !s.Force {
		if s.Files.FileExists("/usr/local/go/bin/go") {
			cmd := exec.Command("/usr/local/go/bin/go", "version")
			out, err := cmd.CombinedOutput()
			if err != nil {
				return fmt.Errorf("golang >> version check failed: %w", err)
			}

			if strings.Contains(string(out), s.GoVersion) {
				in, err := console.Question(fmt.Sprintf("You already have Go %s want me to override [y,N]", s.GoVersion), false)
				if err != nil {
					return fmt.Errorf("golang >> override %w", err)
				}
				if strings.Contains(in, "y") || strings.Contains(in, "yes") {
					overRide = true
				}
			}
		} else {
			overRide = true
		}
	}

	if overRide {
		console.Installing("Go")
		console.Info(fmt.Sprintf("Downloading GoLang %s", s.GoVersion))
		err := s.Files.GetFile("https://dl.google.com/go", fmt.Sprintf("go%s.linux-amd64.tar.gz", s.GoVersion), fmt.Sprintf("go%s.tar.gz", s.GoVersion))
		if err != nil {
			return fmt.Errorf("golang >> get file: %w", err)
		}

		if !s.Files.FileExists(fmt.Sprintf("go%s.tar.gz", s.GoVersion)) {
			return errors.New("golang >> go tar didnt save")
		}

		err = exec.Command("tar", "-C", "/usr/local", "-xzf", fmt.Sprintf("%s/go%s.tar.gz", s.Prefix, s.GoVersion)).Run()
		if err != nil {
			return fmt.Errorf("golang >> untar file: %w", err)
		}
		console.Installed("Go")
		return nil
	}

	console.Skipping("Go")
	return nil
}

func (s Settings) installApps() error {
	// apps
	apps := []string{
		"multitail",
		"htop",
		"git-extras",
		"colordiff",
		"dnsmasq",
		"jq",
		"nmap",
		"prettyping",
		"npm",
	}
	for _, app := range apps {
		err := s.appInstall(app)
		if err != nil {
			return fmt.Errorf("cli: %w", err)
		}
	}

	return nil
}

func (s Settings) appInstall(name string) error {
	exists := false
	if s.Files.FileExists(fmt.Sprintf("/usr/bin/%s", name)) {
		exists = true
	}
	if s.Files.FileExists(fmt.Sprintf("/usr/sbin/%s", name)) {
		exists = true
	}

	if exists == false {
		console.Installing(name)
		err := exec.Command("yum", "install", "-y", name).Run()
		if err != nil {
			return fmt.Errorf("%s err: %w", name, err)
		}
		console.Installed(name)
		return nil
	}
	console.Skipping(name)
	return nil
}

func (s Settings) getFonts() error {
	if !s.Files.UserFileExists(".local/share/fonts") {
		err := s.Files.MakeUserFolder(".local/share")
		if err != nil {
			return fmt.Errorf("getFonts >> share folder: %w", err)
		}

		err = s.Files.MakeUserFolder(".local/share/fonts")
		if err != nil {
			return fmt.Errorf("getFonts >> font folder: %w", err)
		}
	}

	console.Installing("fonts")
	fmt.Println("")
	fonts := []struct {
		Category string
		Name     string
		Format   string
	}{
		{
			Category: "DroidSansMono",
			Name:     "Droid Sans Mono Nerd Font Complete",
			Format:   "otf",
		},
		{
			Category: "Mononoki/Regular",
			Name:     "mononoki-Regular Nerd Font Complete",
		},
		{
			Category: "Mononoki/Bold",
			Name:     "mononoki Bold Nerd Font Complete",
		},
		{
			Category: "Mononoki/Bold-Italic",
			Name:     "mononoki Bold Italic Nerd Font Complete",
		},
		{
			Category: "Mononoki/Italic",
			Name:     "mononoki Italic Nerd Font Complete",
		},
		{
			Category: "DejaVuSansMono/Regular",
			Name:     "DejaVu Sans Mono Nerd Font Complete Mono",
		},
		{
			Category: "DejaVuSansMono/Bold-Italic",
			Name:     "DejaVu Sans Mono Bold Oblique Nerd Font Complete Mono",
		},
		{
			Category: "DejaVuSansMono/Italic",
			Name:     "DejaVu Sans Mono Oblique Nerd Font Complete Mono",
		},
		{
			Category: "DejaVuSansMono/Bold",
			Name:     "DejaVu Sans Mono Bold Nerd Font Complete Mono",
		},
		{
			Category: "Hack/Bold",
			Name:     "Hack Bold Nerd Font Complete",
		},
		{
			Category: "Hack/BoldItalic",
			Name:     "Hack Bold Italic Nerd Font Complete",
		},
		{
			Category: "Hack/Italic",
			Name:     "Hack Italic Nerd Font Complete",
		},
		{
			Category: "Hack/Regular",
			Name:     "Hack Regular Nerd Font Complete",
		},
		{
			Category: "Meslo/L-DZ/Regular",
			Name:     "Meslo LG L DZ Regular Nerd Font Complete",
			Format:   "otf",
		},
		{
			Category: "Meslo/L/Regular",
			Name:     "Meslo LG L Regular Nerd Font Complete",
			Format:   "otf",
		},
		{
			Category: "Meslo/M-DZ/Regular",
			Name:     "Meslo LG M DZ Regular Nerd Font Complete",
			Format:   "otf",
		},
		{
			Category: "Meslo/M/Regular",
			Name:     "Meslo LG M Regular Nerd Font Complete",
			Format:   "otf",
		},
		{
			Category: "Meslo/S-DZ/Regular",
			Name:     "Meslo LG S DZ Regular Nerd Font Complete",
			Format:   "otf",
		},
		{
			Category: "Meslo/S/Regular",
			Name:     "Meslo LG S Regular Nerd Font Complete",
			Format:   "otf",
		},
		{
			Category: "Monoid/Bold",
			Name:     "Monoid Bold Nerd Font Complete",
		},
		{
			Category: "Monoid/Italic",
			Name:     "Monoid Italic Nerd Font Complete",
		},
		{
			Category: "Monoid/Regular",
			Name:     "Monoid Regular Nerd Font Complete",
		},
		{
			Category: "Monoid/Retina",
			Name:     "Monoid Retina Nerd Font Complete",
		},
		{
			Category: "ProFont/profontiix",
			Name:     "ProFont IIx Nerd Font Complete",
		},
		{
			Category: "ProFont/ProFontWinTweaked",
			Name:     "ProFontWindows Nerd Font Complete",
		},
		{
			Category: "RobotoMono/Bold-Italic",
			Name:     "Roboto Mono Bold Italic Nerd Font Complete",
		},
		{
			Category: "RobotoMono/Bold",
			Name:     "Roboto Mono Bold Nerd Font Complete",
		},
		{
			Category: "RobotoMono/Italic",
			Name:     "Roboto Mono Italic Nerd Font Complete",
		},
		{
			Category: "RobotoMono/Light-Italic",
			Name:     "Roboto Mono Light Italic Nerd Font Complete",
		},
		{
			Category: "RobotoMono/Light",
			Name:     "Roboto Mono Light Nerd Font Complete",
		},
		{
			Category: "RobotoMono/Medium-Italic",
			Name:     "Roboto Mono Medium Italic Nerd Font Complete",
		},
		{
			Category: "RobotoMono/Medium",
			Name:     "Roboto Mono Medium Nerd Font Complete",
		},
		{
			Category: "RobotoMono/Regular",
			Name:     "Roboto Mono Regular Nerd Font Complete",
		},
		{
			Category: "RobotoMono/Thin-Italic",
			Name:     "Roboto Mono Thin Italic Nerd Font Complete",
		},
		{
			Category: "RobotoMono/Thin",
			Name:     "Roboto Mono Thin Nerd Font Complete",
		},
		{
			Category: "SourceCodePro/Black-Italic",
			Name:     "Sauce Code Pro Black Italic Nerd Font Complete",
		},
		{
			Category: "SourceCodePro/Black",
			Name:     "Sauce Code Pro Black Nerd Font Complete",
		},
		{
			Category: "SourceCodePro/Bold-Italic",
			Name:     "Sauce Code Pro Bold Italic Nerd Font Complete",
		},
		{
			Category: "SourceCodePro/Bold",
			Name:     "Sauce Code Pro Bold Nerd Font Complete",
		},
		{
			Category: "SourceCodePro/Extra-Light",
			Name:     "Sauce Code Pro ExtraLight Nerd Font Complete",
		},
		{
			Category: "SourceCodePro/ExtraLight-Italic",
			Name:     "Sauce Code Pro ExtraLight Italic Nerd Font Complete",
		},
		{
			Category: "SourceCodePro/Italic",
			Name:     "Sauce Code Pro Italic Nerd Font Complete",
		},
		{
			Category: "SourceCodePro/Light-Italic",
			Name:     "Sauce Code Pro Light Italic Nerd Font Complete",
		},
		{
			Category: "SourceCodePro/Medium-Italic",
			Name:     "Sauce Code Pro Medium Italic Nerd Font Complete",
		},
		{
			Category: "SourceCodePro/Medium",
			Name:     "Sauce Code Pro Medium Nerd Font Complete",
		},
		{
			Category: "SourceCodePro/Regular",
			Name:     "Sauce Code Pro Regular Nerd Font Complete",
		},
		{
			Category: "SourceCodePro/Semibold-Italic",
			Name:     "Sauce Code Pro Semibold Italic Nerd Font Complete",
			Format:   "ttf",
		},
		{
			Category: "SourceCodePro/Semibold",
			Name:     "Sauce Code Pro Semibold Nerd Font Complete",
		},
		{
			Category: "SourceCodePro/Black-Italic",
			Name:     "Sauce Code Pro Black Italic Nerd Font Complete",
		},
		{
			Category: "Terminus/terminus-ttf-4.40.1/Bold",
			Name:     "Terminus (TTF) Bold Nerd Font Complete",
		},
		{
			Category: "Terminus/terminus-ttf-4.40.1/BoldItalic",
			Name:     "Terminus (TTF) Bold Italic Nerd Font Complete",
		},
		{
			Category: "Terminus/terminus-ttf-4.40.1/Italic",
			Name:     "Terminus (TTF) Italic Nerd Font Complete",
		},
		{
			Category: "Terminus/terminus-ttf-4.40.1/Regular",
			Name:     "Terminus (TTF) Nerd Font Complete",
		},
		{
			Category: "Ubuntu/Bold-Italic",
			Name:     "Ubuntu Bold Italic Nerd Font Complete",
		},
		{
			Category: "Ubuntu/Bold",
			Name:     "Ubuntu Bold Nerd Font Complete",
		},
		{
			Category: "Ubuntu/Condensed",
			Name:     "Ubuntu Condensed Nerd Font Complete",
		},
		{
			Category: "Ubuntu/Light-Italic",
			Name:     "Ubuntu Light Italic Nerd Font Complete",
		},
		{
			Category: "Ubuntu/Light",
			Name:     "Ubuntu Light Nerd Font Complete",
		},
		{
			Category: "Ubuntu/Medium-Italic",
			Name:     "Ubuntu Medium Italic Nerd Font Complete",
		},
		{
			Category: "Ubuntu/Medium",
			Name:     "Ubuntu Medium Nerd Font Complete",
		},
		{
			Category: "Ubuntu/Regular-Italic",
			Name:     "Ubuntu Regular Italic Nerd Font Complete",
		},
		{
			Category: "Ubuntu/Regular",
			Name:     "Ubuntu Regular Nerd Font Complete",
		},
	}

	for _, font := range fonts {
		fontName := fmt.Sprintf("%s.ttf", font.Name)
		if font.Format != "" {
			fontName = fmt.Sprintf("%s.%s", font.Name, font.Format)
		}

		if !s.Files.UserFileExists(fmt.Sprintf(".local/share/fonts/%s", fontName)) {
			console.Info(fmt.Sprintf("font: %s", font.Name))

			err := s.Files.GetFile(fmt.Sprintf("https://github.com/ryanoasis/nerd-fonts/raw/master/patched-fonts/%s/complete", font.Category), fontName, fmt.Sprintf(".local/share/fonts/%s", fontName))
			if err != nil {
				return fmt.Errorf("getFonts >> getfont: %w", err)
			}
		}
	}

	err := s.Files.SetPerms(".local/share/fonts")
	if err != nil {
		return fmt.Errorf("getFonts >> perms: %w", err)
	}
	console.Installed("fonts")

	return nil
}

func (s Settings) diffSoFancy() error {
	if !s.Files.FileExists("/usr/local/bin/diff-so-fancy") || s.Force {
		console.Installing("diff-so-fancy")

		if !s.Files.UserFileExists("bin") {
			err := s.Files.MakeUserFolder("bin")
			if err != nil {
				return fmt.Errorf("diff-so-fancy >> bin dir: %w", err)
			}
		}

		err := s.Files.GetFile("https://raw.githubusercontent.com/so-fancy/diff-so-fancy/master/third_party/build_fatpack", "diff-so-fancy")
		if err != nil {
			return fmt.Errorf("diff-so-fancy >> getfile %w", err)
		}
		err = s.Files.MoveFile(fmt.Sprintf("%s/%s", s.Prefix, "diff-so-fancy"), fmt.Sprintf("%s/bin/diff-so-fancy", s.Prefix))
		if err != nil {
			return fmt.Errorf("diff-so-fancy >> move: %w", err)
		}

		console.Installed("diff-so-fancy")
		return nil
	}

	console.Skipping("diff-so-fancy")
	return nil
}

func (s Settings) kr() error {
	if !s.Files.FileExists("/etc/yum.repos.d/kryptco.repo") || s.Force {
		console.Installing("kryptonite")

		err := exec.Command("yum-config-manager", "--add-repo", "https://krypt.co/repo/kryptco.repo").Run()
		if err != nil {
			return fmt.Errorf("kryptonite >> add repo: %w", err)
		}

		err = exec.Command("yum", "-y", "install", "kr").Run()
		if err != nil {
			return fmt.Errorf("kryptoite >> install kryptonite: %w", err)
		}

		console.Installed("kryptonite")
		return nil
	}

	console.Skipping("kryptonite")
	return nil
}

func (s Settings) npmPackages() error {
	packages := []string{
		"yarn",
		"localtunnel",
	}

	for _, pack := range packages {
		if !s.Files.FileExists(fmt.Sprintf("/usr/bin/%s", pack)) || s.Force {
			err := exec.Command("npm", "install", "-g", pack).Run()
			if err != nil {
				return fmt.Errorf("npmpackage >> install %s: %w", pack, err)
			}
		}
	}

	return nil
}
