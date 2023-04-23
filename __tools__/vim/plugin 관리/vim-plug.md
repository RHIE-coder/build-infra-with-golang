# [vim-plug](https://github.com/junegunn/vim-plug)

## [Download]

### - `plug.vim`

```sh
curl -fLo ~/.vim/autoload/plug.vim --create-dirs \
    https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim
```

### - Getting Started

 - add content in `~/.vimrc`

```sh
" Plugins will be downloaded under the specified directory.
call plug#begin(has('nvim') ? stdpath('data') . '/plugged' : '~/.vim/plugged')

" Declare the list of plugins.
Plug 'tpope/vim-sensible'
Plug 'junegunn/seoul256.vim'

" List ends here. Plugins become visible to Vim after this call.
call plug#end()
```

 - installation

```sh
:source ~/.vimrc
:PlugInstall
```

### - programming language syntax highligting : [vim-polyglot](https://github.com/sheerun/vim-polyglot)

```sh
set nocompatible

call plug#begin()

Plug 'sheerun/vim-polyglot'

call plug#end()
```

## [Example]

```sh
" Plugins will be downloaded under the specified directory.
call plug#begin(has('nvim') ? stdpath('data') . '/plugged' : '~/.vim/plugged')

" Declare the list of plugins.
Plug 'junegunn/seoul256.vim'
Plug 'sheerun/vim-polyglot'
Plug 'vim-airline/vim-airline'
Plug 'scrooloose/nerdtree'
Plug 'airblade/vim-gitgutter'
Plug 'scrooloose/syntastic'
Plug 'ctrlpvim/ctrlp.vim'
Plug 'nanotech/jellybeans.vim'

" List ends here. Plugins become visible to Vim after this call.
call plug#end()
```
