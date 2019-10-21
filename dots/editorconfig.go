package dots

import (
	"fmt"
	"os"

	"github.com/keloran/go-dotfiles/console"
)

func (d Dots) editorConfig() error {
	if !d.Files.UserFileExists(".editorconfig") || d.Force {
		if d.Force {
			console.Warning("Forced .editorconfig")
		}

		if d.Github {
			return d.Files.GetGithubFile(".editorconfig")
		}

		fmt.Println(console.Yellow(fmt.Sprintf("Creating editorconfig in %s", d.Prefix)))
		return d.createEditorConfig()
	}

	fmt.Println(console.Cyan("Skipped editorconfig"))
	return nil
}

func (d Dots) createEditorConfig() error {
	f, err := os.Create(fmt.Sprintf("%s/.editorconfig", d.Prefix))
	if err != nil {
		return fmt.Errorf("editorconfig create err: %w", err)
	}

	_, err = f.WriteString("# EditorConfig\n" +
		"root = true\n\n" +
		"[*]\n" +
		"end_of_line = LF\n" +
		"insert_final_newline = true\n" +
		"charset = utf-8\n" +
		"indent_style = space\n" +
		"indent_size = 4\n" +
		"trim_trailing_whitespace = true\n" +
		"tab_width = 4\n\n" +
		"[*.go]\n" +
		"indent_style = tab\n\n" +
		"[{*.yml,*.yaml}]\n" +
		"indent_size = 2\n\n" +
		"[{.babelrc,yest.config,.stylelintrc,.eslintrc,*.json,*.jsb3,*.jsb2,*.bowerrc}]\n" +
		"indent_size = 2\n")
	if err != nil {
		return fmt.Errorf("editorconfig: %w", err)
	}

	err = f.Close()
	if err != nil {
		return fmt.Errorf("editorconfig close file: %w", err)
	}

	err = d.Files.SetUserPerm(".editorconfig")
	if err != nil {
		return fmt.Errorf("editorconfig: %w", err)
	}

	return nil
}
