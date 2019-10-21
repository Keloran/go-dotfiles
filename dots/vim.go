package dots

import (
	"fmt"
	"os"

	"github.com/keloran/go-dotfiles/console"
)

func (d Dots) vim() error {
	if !d.Files.UserFileExists(".vimrc") || d.Force {
		if d.Force {
			console.Warning("Forced .vimrc")
		}

		if d.Github {
			return d.Files.GetGithubFile(".vimrc")
		}

		console.Nice("Creating .vimrc")
		err := d.createVimRC()
		if err != nil {
			return fmt.Errorf("vimrc err: %w", err)
		}
		return nil
	}

	console.Info("Skipped .vimrc")
	return nil
}

func (d Dots) createVimRC() error {
	f, err := os.Create(fmt.Sprintf("%/.vimrc", d.Prefix))
	if err != nil {
		return fmt.Errorf("create vimrc err: %w", err)
	}

	_, err = f.WriteString("\" Vundle\n" +
		"set nocompatiblen\n" +
		"filetype off\n" +
		"set rtp+=~/.vim/bundle/Vundle.vim\n\n" +
		"\" Start Vundle\n" +
		"call vundle#begin()\n\n" +
		"\" Vundle Plugin\n" +
		"Plugin 'gmarik/Vundle.vim'\n\n" +
		"\" Language\n" +
		"Plugin 'instant-markdown,vim'\n" +
		"Plugin 'majutsushi/tagbar'\n" +
		"Plugin 'sheerun/vim-polyglot'\n\n" +
		"\" Git\n" +
		"Plugin 'tpope/vim-fugitive'\n" +
		"Plugin 'esneider/YUNOcommit.vim'\n" +
		"Plugin 'tpope/vim-rhubarb'\n" +
		"Plugin 'airblade/vim-gitgutter'\n\n" +
		"\" UI\n" +
		"Plugin 'vim-airline/vim-airline'\n" +
		"Plugin 'vim-airline/vim-airline-themes'\n" +
		"Plugin 'scrooloose/syntastic'\n" +
		"Plugin 'scrooloose/nerdtree'\n\n" +
		"\" Name conflicts\n" +
		"Plugin 'L9'\n\n" +
		"\" Commands\n" +
		"Plugin 'wincent/command-t'\n" +
		"Plugin 'tpope/vim-eunuch'\n" +
		"Plugin 'tpope/vim-sensible'\n" +
		"Plugin 'jremmen/vim-ripgrep'\n" +
		"Plugin 'mileszs/ack.vim'\n\n" +
		"\" SparkLine\n" +
		"Plugin 'rstacruz/sparkup',{'rtp':'vim/'}\n\n" +
		"\" Colors\n" +
		"Plugin 'flazz/vim-colorschemes'\n\n" +
		"\" AutoComplete\n" +
		"Plugin 'ervandew/supertab'\n" +
		"Plugin 'tpope/vim-sleuth'\n\n" +
		"\" EditorConfig\n" +
		"Plugin 'editorconfig/editorconfig-vim'\n\n" +
		"\" Quotes\n" +
		"Plugin 'tpope/vim-surround'\n\n" +
		"\" Stop Vundle\n" +
		"call vundle#end()\n\n" +
		"\" Indent and Highlight\n" +
		"filetype plugin indent on\n" +
		"syntax on\n\n" +
		"\" Colors\n" +
		"set laststatus=2\n" +
		"let g:airline_theme = 'molokai'\n" +
		"let g:airline_powerline_fonts = 1\n" +
		"colorscheme molokai\n\n" +
		"\" Tabs\n" +
		"set smarttab\n" +
		"set expandtab\n" +
		"set autoindent\n" +
		"set smartindent\n\n" +
		"\" Paste\n" +
		"set paste\n\n" +
		"\" MouseScrolling\n" +
		"set mouse=a\n\n" +
		"\" Languages\n" +
		"if has(\"autocmd\")\n" +
		"  filetype on\n" +
		"  autocmd bufwritepost .vimrc source $MYVIMRC\n" +
		"  autocmd BufRead,BufNewFile *.json set autoindent filetype=javascript\n" +
		"  autocmd BufRead,BufNewFile *.md set filetype=markdown\n" +
		"endif\n\n" +
		"\" Keys\n" +
		"set switchbuf=usetab\n" +
		"nnoremap <F7> :sbnext<CR>\n" +
		"nnoremap <S-F7> :sbprevious<CR>\n" +
		"nmap <F8> :TagbarToggle<CR>\n" +
		"nmap <S-F8> :TagbarOpenAutoClose<CR>\n" +
		"nmap <F4> <Plug>CommentaryLine\n" +
		"nmap md :InstantMarkdownPreview<CR>\n" +
		"nmap <C-n> :NERDTreeToggle<CR>\n" +
		"nmap <S-n> :NERDTreeFocus<CR>\n\n" +
		"\" Markdown\n" +
		"l g:instant_markdown_autostart = 0\n")
	if err != nil {
		return fmt.Errorf("VimRC: %w", err)
	}

	err = f.Close()
	if err != nil {
		return fmt.Errorf("vimrc close file: %w", err)
	}

	err = d.Files.SetUserPerm(".vimrc")
	if err != nil {
		return fmt.Errorf("vimrc: %w", err)
	}

	return nil
}
