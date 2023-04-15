# VIM 설정

```
" vim config file
set nocompatible        " 오리지날 VI와 호환하지 않음
set autoindent          " 자동 들여쓰기
set cindent             " c 스타일 들여쓰기 사용
set number              " 행 번호 표시하기
set tabstop=4           " 탭의 공백 수치 설정
set shiftwidth=4        " { 입력 후의 자동 들여쓰기 수치
syntax on               " 문법 강조 표시
set visualbell          " 경고음 대신 화면을 깜빡임
set showcmd             " 현재 입력하고 있는 명령 표시
set colorcolumn=80      " 색상으로 열을 제한
set textwidth=79        " 열 너비를 79까지 허용
set cursorline          " 커서가 있는 행에 언더라인 표시
set hlsearch            " 검색어 하이라이팅
set sw=1                " 스크롤바 너비
set autoread            " 작업 중인 파일 외부에서 변경됐을 경우 자동으로 불러옴
set showmatch           " 일치하는 괄호 하이라이팅
set smartcase           " 검색시 대소문자 구별
set smartindent         " 스마트한 들여쓰기
set scrolloff=5         " 5개 행을 미리 스크롤 
set paste               " 복사 붙여넣기 시 자동 개행 꺼짐
set encoding=utf-8      " 인코딩 UTF-8로 설정
set incsearch           " 키워드 입력시 점진적 검색
filetype indent on      " 파일 종류에 따른 구문강조
set expandtab           " 탭대신 스페이스
set ruler               " 화면 우측 하단에 현재 커서의 위치(줄,칸) 표시
set tenc=utf-8          " 터미널 인코딩
set wrap                " window 크기가 부족하면 행 다보이게 하기(<> nowrap)
set nowrapscan          " 검색할 때 문서의 끝에서 처음으로 안돌아감
set backspace=eol,start,indent 
    "  줄의 끝, 시작, 들여쓰기에서 백스페이스시 이전줄로
colorscheme morning     " color theme 설정
```

# Color Theme List

 - `/usr/share/vim/vim**/colors` : **은 버전

# Stopped 된 Vim으로 돌아가기

 - 예기치 않게 종료되면 .swp가 남게됨
 - `[CTRL] + z`를 누르면 suspend

```
[bash]# jobs       // suspend된 vim 리스트
[bash]# fg <index> // 해당 vim으로 돌아가기
```

# 유용한 단축키

 - `@:` 이전 커맨드 모드 실행
 - `q:` 명령어 히스토리 보기 `[ENTER]`