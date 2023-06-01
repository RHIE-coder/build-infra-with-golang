# INSTALL(UBUNTU)

## [ Prerequisite ] 

### - 기반 패키지

```sh
apt-get install \
    build-essential zlib1g-dev libssl-dev libncurses-dev \
    libffi-dev libsqlite3-dev libreadline-dev libbz2-dev
```

### - 사전에 파이썬 설치 필요

 - virtualenv도 설치


## - [ EB CLI 설치 ]

```sh
git clone https://github.com/aws/aws-elastic-beanstalk-cli-setup.git
python ./aws-elastic-beanstalk-cli-setup/scripts/ebcli_installer.py
```

## [ ISSUE ]

### - `/usr/bin/env: ‘python’: No such file or directory`

아래 명령어로 해결

```sh
sudo apt install python-is-python3
```

<hr><br><br><br><br><br>

# USAGE

## [ basic flow ]

```
eb init
```