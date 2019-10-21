package osx

import (
	"fmt"
	"os/exec"

	"github.com/keloran/go-dotfiles/console"
)

// CLI ...
func (s Settings) CLI() error {
	console.Start("OSX CLI Apps")

	err := s.brewInstall()
	if err != nil {
		return fmt.Errorf("cli >> %w", err)
	}

	err = s.kr()
	if err != nil {
		return fmt.Errorf("cli >> %w", err)
	}

	err = s.apps()
	if err != nil {
		return fmt.Errorf("cli >> %w", err)
	}

	err = s.fonts()
	if err != nil {
		return fmt.Errorf("cli >> %w", err)
	}

	console.End("OSX CLI Apps")

	return nil
}

func (s Settings) brewInstall() error {
	if s.Force || !s.Files.FileExists("/usr/local/bin/brew") {
		console.Installing("brew")

		err := exec.Command("ruby", "-e", "\"$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)\"").Run()
		if err != nil {
			return fmt.Errorf("brew >> install: %w", err)
		}

		console.Installed("brew")
		return nil
	}

	console.Skipping("brew")
	return nil
}

func (s Settings) apps() error {
	console.Installing("brew apps")

	apps := []string{
		"adns",
		"afflib",
		"antibody",
		"asciidoc",
		"asio",
		"augeas",
		"awscli",
		"azure-cli",
		"bash",
		"bdw-gc",
		"bit",
		"boost",
		"c-ares",
		"carthage",
		"cask",
		"certbot",
		"cmake",
		"colordiff",
		"conan",
		"coreutils",
		"crystal",
		"crystal-lang",
		"ctags",
		"dep",
		"dialog",
		"diff-so-fancy",
		"dnsmasq",
		"docbook",
		"docker-clean",
		"dockutil",
		"dungeon",
		"editorconfig",
		"elinks",
		"emacs",
		"engine_pkcs11",
		"fac",
		"freetype",
		"gcc",
		"gdbm",
		"gettext",
		"gflags",
		"git",
		"git-extras",
		"git-sizer",
		"glew",
		"glib",
		"glog",
		"gmp",
		"gnupgg",
		"gnutls",
		"go",
		"gource",
		"gpg-agent",
		"gradle",
		"htop",
		"icu4c",
		"isl",
		"jansoon",
		"jemalloc",
		"jfrog-cli-go",
		"jpeg",
		"jq",
		"jsoncpp",
		"kops",
		"kr",
		"kubernetes-cli",
		"kubespy",
		"lastpass-cli",
		"libarchive",
		"libassuan",
		"libev",
		"libevent",
		"libewf",
		"libffi",
		"libgcrypt",
		"libgpg-error",
		"libidn2",
		"libksba",
		"libmagic",
		"libmaxminddb",
		"libmpc",
		"libp11",
		"libpng",
		"libpq",
		"librdkafka",
		"libsass",
		"libsmi",
		"libssh",
		"libssh2",
		"libtasn1",
		"libtiff",
		"libtool",
		"libunistring",
		"libusb",
		"libxml2",
		"libyaml",
		"libzip",
		"lldpd",
		"llvm",
		"llvm@5",
		"llvm@6",
		"logstalgia",
		"lua",
		"lua@5.1",
		"luajit",
		"luarocks",
		"lynx",
		"lz4",
		"lzlib",
		"mas",
		"mpfr",
		"multitail",
		"ncurses",
		"nethacked",
		"nettle",
		"nghttp2",
		"nmap",
		"node",
		"npth",
		"nmv",
		"oniguruma",
		"openal-soft",
		"openssl",
		"openssl@1.1",
		"osquery",
		"p11-kit",
		"pcre",
		"pcre2",
		"perl",
		"pinentry",
		"pkgconfig",
		"prettyping",
		"pth",
		"python",
		"python3",
		"python@2",
		"ranger",
		"rapidjson",
		"readline",
		"ripgrep",
		"rocksdb",
		"ruby",
		"ruby@2.0",
		"rust",
		"rustup-init",
		"s3cmd",
		"saml2aws",
		"sassc",
		"sdl2",
		"sdl2_image",
		"siege",
		"sleuthkit",
		"snappy",
		"source-highlight",
		"sqlite",
		"ssdeep",
		"ssh-copy-id",
		"swig",
		"telnet",
		"the_silver_searcher",
		"thrift",
		"tig",
		"tldr",
		"tmux",
		"trash",
		"tree",
		"typescript",
		"unbound",
		"vegeta",
		"vim",
		"watch",
		"webp",
		"wireshark",
		"wtf",
		"xz",
		"yara",
		"yarn",
		"z",
		"zork",
		"zplug",
		"zsh",
		"zstd",
	}

	for _, app := range apps {
		console.Info(fmt.Sprintf("%s installing", app))
		err := exec.Command("brew", "install", app).Run()
		if err != nil {
			return fmt.Errorf("brew >> %s: %w", app, err)
		}
	}

	console.Installed("brew apps")

	return nil
}

func (s Settings) kr() error {
	if s.Force || !s.Files.FileExists("/usr/bin/kr") {
		console.Installing("kryptonite")

		err := exec.Command("brew", "install", "kryptco/tap/kr").Run()
		if err != nil {
			return fmt.Errorf("kryptonite >> install: %w", err)
		}

		err = exec.Command("echo", "\"n\"", "|", "kr", "codesign").Run()
		if err != nil {
			return fmt.Errorf("kryptonite >> codesign: %w", err)
		}

		console.Installed("kryptonite")
		return nil
	}

	console.Skipping("kryptonite")
	return nil
}

func (s Settings) fonts() error {
	console.Installing("fonts")

	fonts := []string{
		"anonymouspro-nerd-font",
		"consolas-for-powerline",
		"dejavusansmono-nerd-font",
		"firacode-nerd-font",
		"hack-nerd-font",
		"menlo-for-powerline",
		"meslo-nerd-font",
		"monofur-nerd-font",
		"monoid-nerd-font",
		"mononoki-nerd-font",
		"profont-nerd-font",
		"robotomono-nerd-font",
		"source-code-pro-for-powerline",
		"sourcecodepro-nerd-font",
		"spacemono-nerd-font",
		"terminus-nerd-font",
		"ubuntu-nerd-font",
	}

	for _, font := range fonts {
		err := exec.Command("brew", "cask", "install", fmt.Sprint("font-%s", font)).Run()
		if err != nil {
			return fmt.Errorf("font >> %s: %w", font, err)
		}
	}

	console.Installed("fonts")

	return nil
}
