# WRK 2

## 1. Source

https://github.com/giltene/wrk2

<hr><br><br><br><br>

## 2. Install

```sh
sudo apt-get install build-essential libssl-dev git zlib1g-dev
git clone https://github.com/giltene/wrk2.git
cd wrk2
make
```

<hr><br><br><br><br>

## 3. Getting Started

### [기본 사용법]

```sh
-t{N}     : Thread 개수
-c{N}     : Connection 개수(최소값: Thread 개수)
-R{N}     : TPS 조절
-d{N}s    : 얼마간 진행할 것인가(duration)
-s {name} : 스크립트 파일
```

더 다양한 옵션 존재

### [요청 예시]

 - GET

```sh
./wrk -t20 -c400 -d30s -R10000 http://localhost:3002/health
```

 - POST

LUA 확장명을 가진 파일에 HTTP 요청 명시

##### - myPostBench.lua

```lua
-- example HTTP POST script which demonstrates setting the
-- HTTP method, body, and adding a header

wrk.method = "POST"
wrk.body   = '{"foo": "bar", "baz": 1000}'
wrk.headers["Content-Type"] = "application/json"
```

`요청`

```sh
./wrk -t20 -c400 -d30s -R10000 -s ./myPostBench.lua http://localhost:3002/posts
```