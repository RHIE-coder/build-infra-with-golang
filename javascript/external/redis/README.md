# Redis

## # Intro

캐시 메모리 서버 (`redis-server`)

클라이언트인 `redis-cli`를 통해 접근 가능(`redis-tools`)

 - 공식사이트 [CLICK](https://redis.io/docs/getting-started/installation/install-redis-on-linux/)

<hr><br><br><br><br>

## # Installation

### - Ubuntu 18.04 Installation

If you're running a very minimal distribution (such as a Docker container) you may need to install lsb-release first:

```sh
sudo apt install lsb-release
```

Add the repository to the apt index, update it, and then install:

```sh
curl -fsSL https://packages.redis.io/gpg | sudo gpg --dearmor -o /usr/share/keyrings/redis-archive-keyring.gpg

echo "deb [signed-by=/usr/share/keyrings/redis-archive-keyring.gpg] https://packages.redis.io/deb $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/redis.list

sudo apt-get update
sudo apt-get install redis
```

<br><br><br>

### - Centos 7 Installation 

```sh
sudo yum install epel-release yum-utils
sudo yum install http://rpms.remirepo.net/enterprise/remi-release-7.rpm
sudo yum-config-manager --enable remi
sudo yum install redis
```
<br><br><br>

### - Use `tar.gz`

```sh
wget http://download.redis.io/releases/redis-7.0.4.tar.gz
tar -xzf redis-7.0.4.tar.gz
cd redis-7.0.4
make
```

 - run the redis

```sh
./redis-7.0.4/src/redis-server ./redis-7.0.4/redis.conf
./redis-7.0.4/src/redis-cli -h 10.0.2.15 -p 6379
```


<hr><br><br><br><br>
