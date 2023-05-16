# MongoDB 

## # Intro

Document 기반 NoSQL 서버

 - 공식사이트 [CLICK](https://www.mongodb.com/)

<hr><br><br><br><br>

## # Installation
###  - [Ubuntu 18.04 Installation](https://www.mongodb.com/docs/manual/tutorial/install-mongodb-on-debian/)

#### 설치 전 잠깐!

```
 - `/proc/cpuinfo`에서 `avx` 및 `avx2`가 `grep`되어야 함.
 - 만일 없다면 CPU가 version 5를 지원하지 않음 -> 4.4.16 버전을 받자
 - https://docs.mongodb.com/v4.4/tutorial/install-mongodb-on-ubuntu/
```

1. Import the public key used by the package management system.

```sh
wget -qO - https://www.mongodb.org/static/pgp/server-6.0.asc | sudo apt-key add -
```

 - However, if you receive an error indicating that gnupg is not installed, you can:

> ```sh
> sudo apt-get install gnupg
> ```
>
> Once installed, retry importing the key

2. Create a `/etc/apt/sources.list.d/mongodb-org-6.0.list` file for MongoDB.

```sh
echo "deb http://repo.mongodb.org/apt/debian buster/mongodb-org/6.0 main" | sudo tee /etc/apt/sources.list.d/mongodb-org-6.0.list
```

3. Reload local package database.

```sh
sudo apt-get update
```

4. Install the MongoDB packages.

```sh
sudo apt-get install -y mongodb-org
```

<br><br><br>

###  - [Centos 7 Installation](https://www.mongodb.com/docs/manual/tutorial/install-mongodb-on-red-hat/)

1. Create a `/etc/yum.repos.d/mongodb-org-6.0.repo` file so that you can install MongoDB directly using yum:

```
[mongodb-org-6.0]
name=MongoDB Repository
baseurl=https://repo.mongodb.org/yum/redhat/$releasever/mongodb-org/6.0/x86_64/
gpgcheck=1
enabled=1
gpgkey=https://www.mongodb.org/static/pgp/server-6.0.asc
```

2. install

```
sudo yum install -y mongodb-org
```

### - Check Running

1. Start MongoDB

```sh
sudo systemctl start mongod
```

2. Verify that MongoDB has started successfully

```sh
sudo systemctl status mongod
```

3. Stop MongoDB

```sh
sudo systemctl stop mongod
```

4. Restart MongoDB

```sh
sudo systemctl restart mongod
```


<br><br><br>

### - Docker

```sh
sudo docker pull mongo
docker run --name mongodb-container -v ~/mongodata:/data/db -d -p 27017:27017 mongo:4.4.16

# 확인
docker ps -a

# 컨테이너 중지
docker stop mongodb-container

# 컨테이너 시작
docker start mongodb-container

# 컨테이너 재시작
docker restart mongodb-container
```

<br><br><br>

### - Directories

By default, a MongoDB instance stores:
 - its data files in `/var/lib/mongodb`
 - its log files in `/var/log/mongodb`

<br><br><br>

### - if "Failed to start mongod.service: Unit mongod.service not found." 

```sh
sudo systemctl daemon-reload
```

<hr><br><br><br><br>