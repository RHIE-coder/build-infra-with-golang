# [Vundle](https://github.com/VundleVim/Vundle.vim)

## [Installation]

```
git clone https://github.com/VundleVim/Vundle.vim.git ~/.vim/bundle/Vundle.vim
```
## [Configuraion]

> `~/.vimrc`

```vim
set nocompatible              " be iMproved, required
filetype off                  " required

" set the runtime path to include Vundle and initialize
set rtp+=~/.vim/bundle/Vundle.vim
call vundle#begin()
" alternatively, pass a path where Vundle should install plugins
"call vundle#begin('~/some/path/here')

" let Vundle manage Vundle, required
Plugin 'VundleVim/Vundle.vim'

Plugin ...
Plugin ...
Plugin ...

" All of your Plugins must be added before the following line
call vundle#end()            " required
filetype plugin indent on    " required
" To ignore plugin indent changes, instead use:
"filetype plugin on
"
" Brief help
" :PluginList       - lists configured plugins
" :PluginInstall    - installs plugins; append `!` to update or just :PluginUpdate
" :PluginSearch foo - searches for foo; append `!` to refresh local cache
" :PluginClean      - confirms removal of unused plugins; append `!` to auto-approve removal
"
" see :h vundle for more details or wiki for FAQ
" Put your non-Plugin stuff after this line
```

## [Useful]

```vim
Plugin 'preservim/nerdtree'        " 폴더 구조 Netrw보다 편함
Plugin 'vim-airline/vim-airline'   " 상태바
Plugin 'ctrlpvim/ctrlp.vim'        " 파일 찾기
Plugin 'Syntastic'                 " 코드 문법 체크

" 각 언어에 맞는 개발툴 
Plugin 'fatih/vim-go'              " Example(golang)
```

 - 더 많은 플러그인 [https://vimawesome.com/]