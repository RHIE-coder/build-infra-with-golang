# GO111MODULE

 - `auto` : 기본값으로, 현재 디렉토리 내에 go.mod 파일이 있으면 모듈 모드를 사용하고, 그렇지 않으면 GOPATH 모드를 사용합니다.
 - `on` : 모듈 모드를 사용합니다. 현재 디렉토리 내에 go.mod 파일이 없더라도 모듈 모드를 강제로 사용
 - `off` : GOPATH 모드를 사용합니다.

<br><br><br><br><br>

---


# Install `Golang`

[origin install guide](https://go.dev/doc/install)

##  [ Linux ]

 - GET `tar.gz`

```bash
wget https://go.dev/dl/go1.20.2.linux-amd64.tar.gz
```

1. Remove any previous Go installation by deleting the /usr/local/go folder (if it exists), then extract the archive you just downloaded into /usr/local, creating a fresh Go tree in /usr/local/go:
```bash
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.20.2.linux-amd64.tar.gz
```
(You may need to run the command as root or through sudo).

**Do not** untar the archive into an existing /usr/local/go tree. This is known to produce broken Go installations.

2. Add /usr/local/go/bin to the PATH environment variable.

You can do this by adding the following line to your $HOME/.profile or /etc/profile (for a system-wide installation):
```bash
export PATH=$PATH:/usr/local/go/bin
```
**Note:** Changes made to a profile file may not apply until the next time you log into your computer. To apply the changes immediately, just run the shell commands directly or execute them from the profile using a command such as source $HOME/.profile.

3. Verify that you've installed Go by opening a command prompt and typing the following command:
```bash
go version
```
Confirm that the command prints the installed version of Go.

4. bashrc

```sh
GOROOT=/usr/local/go
GOPATH=$HOME/go
PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```


<br><br><br><br><br>

---


# VSCODE 개발 전용 실행 파일 목록

## [ List ]
 - dlv
 - dlv-dap
 - goimports
 - go-outline
 - gomodifytags
 - gopkgs
 - goplay
 - gopls
 - gotests
 - impl
 - staticcheck

## [ Install Command ]

```sh
# dlv
# - dlv-dap ???
go get github.com/go-delve/delve/cmd/dlv

# goimports
go install golang.org/x/tools/cmd/goimports@latest

# go-outline
go get -u github.com/ramya-rao-a/go-outline

# gomodifytags
go get -u github.com/fatih/gomodifytags
go install github.com/fatih/gomodifytags@latest

# gopkgs
go get -v
go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest

# goplay
go get -u github.com/haya14busa/goplay/cmd/goplay

# gopls
go get -v golang.org/x/tools/gopls@latest
go install golang.org/x/tools/gopls@latest

# gotests
go get -v
go install github.com/cweill/gotests/gotests@latest

# impl
go get -v
go install github.com/josharian/impl@latest

# staticcheck
go get honnef.co/go/tools/cmd/staticcheck
go install honnef.co/go/tools/cmd/staticcheck@2022.1
```
