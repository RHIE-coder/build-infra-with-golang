# 🍀Node.js 설치

## # Ubuntu 18.04

### -  PPA를 통해 최신버전 가져오기(16버전) LTS
```
curl -sL https://deb.nodesource.com/setup_16.x | sudo -E bash -
```

### - install
```
sudo apt-get install -y nodejs
```
### - 확인
```
node -v
npm -v
```

### - NPM의 제 기능을 위해 부가설치(npm install 에러 방지)

```bash
sudo apt-get install build-essential
```

<br><br><br>


## # Centos 7

### - NPM의 제 기능을 위해 부가설치(npm install 에러 방지)

```bash
yum install gcc gcc-c++ kernel-devel kernel-headers
```

<br><br><br>

## # `node-gyp` 설치(OPTIONAL)

```
sudo npm install -g node-gyp
```

<hr><br><br><br><br><br>