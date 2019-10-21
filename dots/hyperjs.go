package dots

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/keloran/go-dotfiles/console"
)

func (d Dots) hyperJS() error {
	if !d.Files.UserFileExists(".hyper.js") || d.Force {
		if d.Force {
			console.Warning("Forced .hyper.js")
		}

		if d.Github {
			return d.Files.GetGithubFile(".hyper.js")
		}

		console.Nice("Creating .hyper.js")
		err := d.createHyperJS()
		if err != nil {
			return fmt.Errorf("createHyperJS: %w", err)
		}
		err = d.getHyperJSFiles()
		if err != nil {
			return fmt.Errorf("getHyperJSFiles: %w", err)
		}

		if runtime.GOOS == "linux" {
			return d.Files.SetPerms(fmt.Sprintf("%s/.hyper_plugins/icons", d.Prefix))
		}

		return nil
	}

	console.Info("Skipped .hyper.js")
	return nil
}

func (d Dots) createHyperJS() error {
	f, err := os.Create(fmt.Sprintf("%s/.hyper.js", d.Prefix))
	if err != nil {
		return fmt.Errorf("HyperJS create err: %w", err)
	}
	_, err = f.WriteString("// HyperJS\n" +
		"module.exports = {\n" +
		"  config: {\n" +
		"    hyperline: {\n" +
		"      plugins: [\n" +
		"        'battery',\n" +
		"        'network',\n" +
		"        'memory',\n" +
		"        'cpu'\n" +
		"      ]\n" +
		"    },\n" +
		"    hyperTabs: {\n" +
		"      tabIcons: false,\n" +
		"      tabIconsColored: true\n" +
		"    },\n" +
		"    hyperLinks: {\n" +
		"      defaultBrowser: true\n" +
		"    },\n" +
		"    hyperCustomTouchbar: [{\n" +
		"      label: 'general',\n" +
		"      backgroundColor: '#000000',\n" +
		"      options: [{\n" +
		"        label: 'clear',\n" +
		"        command: 'clear'\n" +
		"      },{\n" +
		"        label: 'update',\n" +
		"        command: 'updateSys'\n" +
		"      }]\n" +
		"    },{\n" +
		"      icon: '~/.hyper_plugins/icons/localstack.png',\n" +
		"      backgroundColor: '#000000',\n" +
		"      options: [{\n" +
		"        label: 'dynamodb',\n" +
		"        command: 'SERVICES=dynamodb TMPDIR=private$TMPDIR localstack start --docker'\n" +
		"      }]\n" +
		"    },{\n" +
		"      icon: '~/.hyper_plugins/icons/docker.png',\n" +
		"      options: [{\n" +
		"        icon: '~/.hyper_plugins/icons/info.png',\n" +
		"        command: 'docker ps -a',\n" +
		"        backgroundColor: '#6767FF'\n" +
		"      },{\n" +
		"        icon: '~/.hyper_plugins/icons/start.png',\n" +
		"        command: 'dockerStart',\n" +
		"        backgroundColor: '#6767FF'\n" +
		"      },{\n" +
		"        icon: '~/.hyper_plugins/icons/stop.png',\n" +
		"        command: 'dockerStop',\n" +
		"        backgroundColor: '#6767FF'\n" +
		"      }]\n" +
		"    },{\n" +
		"      icon: '~/.hyper_plugins/icons/github.png',\n" +
		"      options: [{\n" +
		"        command: 'git diff',\n" +
		"        icon: '~/.hyper_plugins/icons/diff.png',\n" +
		"        backgroundColor: '#CFCFCF'\n" +
		"      },{\n" +
		"        command: 'git status',\n" +
		"        icon: '~/.hyper_plugins/icons/info.png',\n" +
		"        backgroundColor: '#CFCFCF'\n" +
		"      },{\n" +
		"        command: 'git log',\n" +
		"        icon: '~/.hyper_plugins/icons/log.png',\n" +
		"        backgroundColor: '#CFCFCF'\n" +
		"      },{\n" +
		"        command: 'git add .',\n" +
		"        icon: '~/.hyper_plugins/icons/add.png',\n" +
		"        backgroundColor: '#CFCFCF'\n" +
		"      },{\n" +
		"        command: 'git clone',\n" +
		"        icon: '~/.hyper_plugins/icons/download.png',\n" +
		"        backgroundColor: '#CFCFCF',\n" +
		"        prompt: true\n" +
		"      }]\n" +
		"    },{\n" +
		"      icon: '~/.hyper_plugins/icons/vim.png',\n" +
		"      backgroundColor: '#B2D8B2',\n" +
		"      options: [{\n" +
		"        icon: '~/.hyper_plugins/icons/quit.png',\n" +
		"        command: ':q!',\n" +
		"        esc: true\n" +
		"      },{\n" +
		"        icon: '~/.hyper_plugins/icons/save.png',\n" +
		"        command: ':w',\n" +
		"        esc: true\n" +
		"      }]\n" +
		"    }],\n" +
		"    fontSize: 11,\n" +
		"    fontFamily: '\"Monoid Nerd Font\", \"MesloLGM Nerd Font\", \"DejaVuSansMono Nerd Font\", \"TerminessTTF Nerd Font\", \"SauceCodePro Nerd Font\"',\n" +
		"    cursorColor: 'rgba(248, 28, 229, 0.8)',\n" +
		"    cursorShape: 'BLOCK',\n" +
		"    foregroundColor: '#FFFFFF',\n" +
		"    backgroundColor: '#000000',\n" +
		"    borderColor: '#333333',\n" +
		"    css: '',\n" +
		"    termCSS: '.unicode-node {position: relative}',\n" +
		"    showHamburgerMenu: '',\n" +
		"    showWindowControls: '',\n" +
		"    padding: '.5rem .5rem .8rem .5rem',\n" +
		"    colors: {\n" +
		"      black: '#000000',\n" +
		"      red: '#FF0000',\n" +
		"      green: '#33FF00',\n" +
		"      yellow: '#FFFF00',\n" +
		"      blue: '#0066FF',\n" +
		"      magenta: '$CC00FF',\n" +
		"      cyan: '#00FFFF',\n" +
		"      white: '#D0D0D0',\n" +
		"      lightBlack: '#808080',\n" +
		"      lightRed: '#FF0000',\n" +
		"      lightGreen: '#33FF00',\n" +
		"      lightYellow: '#FFFF00',\n" +
		"      lightBlue: '#0066FF',\n" +
		"      lightMagenta: '#CC00FF',\n" +
		"      lightCyan: '#00FFFF',\n" +
		"      lightWhite: '#FFFFFF'\n" +
		"    },\n" +
		"    shell: '/bin/zsh',\n" +
		"    shellArgs: ['--login'],\n" +
		"    env: {},\n" +
		"    bell: 'SOUND',\n" +
		"    copyOnSelect: true,\n" +
		"  },\n" +
		"  plugins: [\n" +
		"    'hyperline',\n" +
		"    'hyperlinks',\n" +
		"    'hyper-tabs-enhanced',\n" +
		"    'hyper-history',\n" +
		"    'hyper-afterglow',\n" +
		"    'hyper-broadcast',\n" +
		"    'hyper-hide-title',\n" +
		"    'hypercwd',\n" +
		"    'hyper-custom-touchbar',\n" +
		"    'hyper-quit'\n" +
		"  ],\n" +
		"  localPlugins: []\n" +
		"}\n")
	if err != nil {
		return fmt.Errorf("HyperJS: %w", err)
	}

	err = f.Close()
	if err != nil {
		return fmt.Errorf("hyperjs close file err: %w", err)
	}

	err = d.Files.SetUserPerm(".hyper.js")
	if err != nil {
		return fmt.Errorf("HyperJS: %w", err)
	}

	return nil
}

func (d Dots) getHyperJSFiles() error {
	console.Info("Get HyperJS image files")
	if !d.Files.UserFileExists(".hyper_plugins/icons") {
		err := os.Mkdir(fmt.Sprintf("%s/.hyper_plugins", d.Prefix), 755)
		if err != nil {
			return fmt.Errorf("create hyper_plugins folder: %w", err)
		}

		err = os.Mkdir(fmt.Sprintf("%s/.hyper_plugins/icons", d.Prefix), 755)
		if err != nil {
			return fmt.Errorf("create hyperjs icons folder: %w", err)
		}
	}

	// add icon
	err := d.getIconFile("add.png")
	if err != nil {
		return fmt.Errorf("add icon err: %w", err)
	}

	// localstack icon
	err = d.getIconFile("form.png", "localstack.png")
	if err != nil {
		return fmt.Errorf("localstack icon err: %w", err)
	}

	// download icon
	err = d.getIconFile("downloading-updates.png", "download.png")
	if err != nil {
		return fmt.Errorf("download icon err: %w", err)
	}

	// diff icon
	err = d.getIconFile("change-theme.png", "diff.png")
	if err != nil {
		return fmt.Errorf("diff icon err: %w", err)
	}

	// save icon
	err = d.getIconFile("save-all.png", "save.png")
	if err != nil {
		return fmt.Errorf("save icon err: %w", err)
	}

	// info icon
	err = d.getIconFile("info.png")
	if err != nil {
		return fmt.Errorf("info icon err: %w", err)
	}

	// quit icon
	err = d.getIconFile("close-window.png", "quit.png")
	if err != nil {
		return fmt.Errorf("quit icon err: %w", err)
	}

	// log icon
	err = d.getIconFile("view-details.png", "log.png")
	if err != nil {
		return fmt.Errorf("log icon err: %w", err)
	}

	// start icon
	err = d.getIconFile("play.png", "start.png")
	if err != nil {
		return fmt.Errorf("start icon err: %w", err)
	}

	// stop icon
	err = d.getIconFile("stop.png")
	if err != nil {
		return fmt.Errorf("stop icon err: %w", err)
	}

	// editor
	err = d.getIconFile("multi-edit.png", "editor.png")
	if err != nil {
		return fmt.Errorf("editor icon err: %w", err)
	}

	// docker
	err = d.getDockerImageFile()
	if err != nil {
		return fmt.Errorf("docker icon err: %w", err)
	}

	// github
	err = d.getGithubImageFile()
	if err != nil {
		return fmt.Errorf("github icon err: %w", err)
	}

	return nil
}

func (d Dots) getIconFile(name ...string) error {
	finalName := name[0]

	err := d.Files.GetFile("https://img.icons8.com/dotty/344", name[0])
	if err != nil {
		return fmt.Errorf("get icon file: %w", err)
	}

	if len(name) == 2 {
		err = os.Rename(fmt.Sprintf("%s/%s", d.Prefix, name[0]), fmt.Sprintf("%s/%s", d.Prefix, name[1]))
		if err != nil {
			return fmt.Errorf("rename %s to %s err: %w", name[0], name[1], err)
		}
		finalName = name[1]
	}

	err = os.Rename(fmt.Sprintf("%s/%s", d.Prefix, finalName), fmt.Sprintf("%s/.hyper_plugins/icons/%s", d.Prefix, finalName))
	if err != nil {
		return fmt.Errorf("move file to hyperjs icons folder: %w", err)
	}

	return nil
}

func (d Dots) getDockerImageFile() error {
	err := d.Files.GetFile("https://www.shareicon.net/download/128x128//2017/02/15", "878789_media_512x512.png")
	if err != nil {
		return fmt.Errorf("get docker icon image err: %w", err)
	}

	err = os.Rename(fmt.Sprintf("%s/878789_media_512x512.png", d.Prefix), fmt.Sprintf("%s/.hyper_plugins/icons/docker.png", d.Prefix))
	if err != nil {
		return fmt.Errorf("rename docker icon image: %w", err)
	}

	return nil
}

func (d Dots) getGithubImageFile() error {
	err := d.Files.GetFile("https://www.shareicon.net/download/128x128//2016/06/20", "606578_black_256x256.png")
	if err != nil {
		return fmt.Errorf("get github icon image: %w", err)
	}

	err = os.Rename(fmt.Sprintf("%s/606578_black_256x256.png", d.Prefix), fmt.Sprintf("%s/.hyper_plugins/icons/github.png", d.Prefix))
	if err != nil {
		return fmt.Errorf("rename github icon image: %w", err)
	}

	return nil
}

func (d Dots) hyperJSPermissions() error {
	fmt.Println("Set HyperJS permissions")

	// chown
	cmd := exec.Command("chown", "-R", d.Username, fmt.Sprintf("%s/.hyper_plugins/icons", d.Prefix))
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("hyperjs permissions chown: %w", err)
	}

	// chgrp
	cmd = exec.Command("chgrp", "-R", "adusers", fmt.Sprintf("%s/.hyper_plugins/icons", d.Prefix))
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("hyperjs permissions chgrp: %w", err)
	}

	// chmod
	cmd = exec.Command("chmod", "755", "-R", fmt.Sprintf("%s/.hyper_plugins/icons", d.Prefix))
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("hyperjs permissions chmod: %w", err)
	}

	return nil
}
