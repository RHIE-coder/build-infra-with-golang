# Docker 명령어들

## # 이미지

### - 이미지 형식

`이미지명[:태그명]`

```
debian:7
debian:latest
```

<br><br><br>

### - 이미지 다운로드

`docker image pull [옵션] 이미지명[:태그명]`

<br><br><br>

### - 이미지 목록 표시

`docker image ls [option] [repo name]`

`docker images`

#### 옵션
 - `--all, -a` 모든 이미지 표시
 - `--digests` 다이제스트 표시 여부
 - `--no-trunc` 결과모두 표시
 - `--quiet, -q` 이미지 ID만 표시

#### DCT

Docker Content Trust로 이미지 서명 검증을 하여 위변조된 이미지 활용 방지

```
export DOCKER_CONTENT_TRUST=1 # 1 유효화 / 0 무효화
```

<br><br><br>

### - 이미지 상세 정보 확인

`docker image inspect`

<br><br><br>

### - 이미지 태그 설정

`<Docker Hub 사용자명>/이미지명:[태그명]`

```
docker image tag nginx rhie/webserver:1.0
```

이미지에 별명을 붙일 뿐. 이미지ID가 똑같음.

<br><br><br>

### - 이미지 검색

`docker search [옵션] [키워드]`

#### 옵션
 - `--no-trunc` 결과를 모두 표시
 - `--limit` n건의 검색 결과를 표시
 - `filter=stars=n` 즐겨찾기의 수(n 이상)를 지정

#### 주의

Docker Hub에 공개되어 있는 이미지가 모두 안전하지 않으므로 공식이미지 및 Dockerfile이 제대로 공개되어 있는 것을 활용하기

<br><br><br>

### - 이미지 삭제

`docker image rm [옵션] 이미지명 [...이미지명]`

#### 옵션
 - `--force, -f` 이미지를 강제로 삭제
 - `--no-prune` 중간 이미지를 삭제하지 않음

<br><br><br>

### - 사용하지 않은 이미지 삭제

`docker image prune [옵션]`

#### 옵션
 - `--all, -a` 사용하지 않는 이미지 모두 삭제
 - `--force, -f` 이미지를 강제로 삭제

<br><br><br>

### - Docker Hub 로그인

`docker login [옵션] [서버]`

#### 옵션
 - `--password, -p`
 - `--username, -u`

서버명을 지정하지 않으면 Docker Hub에 엑세스됨

 <br><br><br>

### - 이미지 업로드

`docker image push <사용자명>/이미지명[:태그명]`

<br><br><br>

### - Docker Hub 로그아웃

`docker logout [서버]`



<hr><br><br><br><br><br>

## # 컨테이너

### - 컨테이너 생성

`docker container create`

컨테이너를 작성할 뿐 시작하지는 않음

<br><br><br>

### - 컨테이너 실행

`docker container run [옵션] 이미지명[:태그명] [인수]`

#### 옵션
 - `--attach, -a`: STDIN, STDOUT, STDERR에 Attach한다.
 - `--cidfile`: 컨테이너 id를 파일로 출력한다.
 - `--detach, -d`: 컨테이너를 생성하고 백그라운드에서 실행한다.
 - `--interactive, -i`: 컨테이너의 표준 입력을 연다
 - `--tty, -t`: 단말기 디바이스를 사용한다.
 - `--user, -u`: 사용자명 지정
 - `--restart=[no | on-failure | on-failure:n(횟수) | always | unless-stopped]`: 명령의 실행 결과에 따라 재시작을 하는 옵션
 - `--rm`: 명령 실행 완료 후에 컨테이너를 자동으로 삭제
 - `--add-host=[호스트명:IP주소]`: 컨테이너의 `/etc/hosts`에 호스트명과 ip주소를 정의
 - `--dns=[IP주소]`: 컨테이너용 DNS 서버의 IP주소 지정
 - `--expose`: 지정한 범위의 포트 번호를 할당
 - `--mac-address=[MAC주소]`: 컨테이너의 MAC 주소를 지정
 - `--net=[bridge | none | container:<name | id> | host | NETWORK(사용자 정의 네트워크)]`: 컨테이너의 네트워크를 지정
 - `--hostname, -h` 컨테이너 자신의 호스트명을 가짐
 - `--publish, -p[호스트의 포트번호]:[컨테이너의 포트번호]` 호스트와 컨테이너의 포트 매핑
 - `--publish-all, -P` 호스트의 임의의 포트를 컨테이너에 할당
 - `--cpu-shares, -c` CPU 사용 배분
 - `--memory, -m` 사용할 메모리를 제한하여 실행(unit: b k m g)
 - `--volume=[호스트의 디렉토리]:[컨테이너의 디렉토리], -v` 호스트와 컨테이너의 디렉토리를 공유
 - `--env=[환경변수], -e` 환경변수를 설정
 - `--env-file=[파일명]` 환경변수를 파일로부터 설정
 - `--read-only=[true | false]` 컨테이너의 파일 시스템을 읽기 전용으로 만듬
 - `--workdir=[경로], -w` 컨테이너의 작업 디렉토리 지정. 컨테이너에 들어갈 때 기본으로 잡히는 경로.
 - `--user=[사용자명], -u` 사용자명 또는 UID 지정

```sh
docker container run -it --name "my-os" ubuntu /bin/bash
docker container run -d -p 8080:80 nginx
```

#### tty, pty

 - `tty`: teletypewriter 항목은 리눅스 디바이스 드라이브 중에서 콘솔이나 터미널을 의미함
 - `pty`: 본체에 LAN으로 연결된 모드(원격 접속, 가상터미널pseudo-terminal)
 

#### bash prompt

 - `#`: root
 - `$`: user

#### --restart

 - `no`: 재시작안함
 - `on-failure`: 종료 status가 0이 아닐 때 재시작
 - `on-failure:n`: `on-failure` 옵션에 n번 재시작
 - `always`: 항상 재시작
 - `unless-stopped`: 정지상태가 아니라면 항상 재시작

<br><br><br>

### - 컨테이너 목록

`docker container ls [옵션]`

#### 옵션

 - `--all, -a`: 실행 중/정지 중인 것도 포함. 모든 컨테이너 표시
 - `--filter, -f`: 표시할 컨테이너의 필터링
 - `--format`: 표시 포맷 지정
 - `--last, -n`: 마지막으로 실행된 n건의 컨테이너만 표시
 - `--latest, -l`: 마지막으로 실행된 컨테이너만 표시
 - `--no-trunc`: 정보를 생략하지 않고 표시
 - `--quiet, -q`: 컨테이너 ID만 표시
 - `--size, -s`: 파일 크기 표시

<br><br><br>

### - 컨테이너 가동 확인

`docker container stats [컨테이너 식별자]`

`CTRL+c`로 종료


<br><br><br>

### - 컨테이너 시작

`docker container start [옵션] <컨테이너 식별자> [컨테이너 식별자]`

#### 옵션

 - `--attach, -a` 표준 출력, 표준 오류 출력을 연다
 - `--interactive, -i` 컨테이너의 표준 입력을 연다

<br><br><br>

### - 컨테이너 정지

`docker container stop [옵션] <컨테이너 식별자> [컨테이너 식별자]`

#### 옵션
 - `--time, -t` 컨테이너의 정지 시간을 지정(기본값은 10초)

<br><br><br>

### - 컨테이너 재시작

`docker container restart [옵션] <컨테이너 식별자> [컨테이너 식별자]`

#### 옵션
 - `--time, -t` 컨테이너 재시작 시작을 지정(기본값은 10초)

<br><br><br>

### - 컨테이너 삭제

`docker container rm [옵션] <컨테이너 식별자> [컨테이너 식별자]`

#### 옵션
 - `--force, -f` 실행 중인 컨테이너를 강제로 삭제
 - `--volumes, -v` 할당한 볼륨을 삭제

<br><br><br>

### - 컨테이너 중단/재개

`docker container pause <컨테이너 식별자>`

`docker container unpause <컨테이너 식별자>`

<hr><br><br><br><br><br>

## # 도커 컨테이너 네트워크

Docker 컨테이너끼리 통신은 Docker 네트워크를 통해 수행함

### - 네트워크 목록 표시

`docker network ls [옵션]`

#### 옵션
 - `--filter=[], -f` 출력을 필터링한다.
 - `--no-trunc` 상세 정보를 출력한다
 - `--quiet, -q` 네트워크 ID만 표시한다.

Docker는 기본값으로 bridge, host, none 이 3개의 네트워크를 만듬


<br><br><br>

### - 네트워크 작성

`docker network create [옵션] 네트워크`

#### 옵션

 - `--driver, -d` 네트워크 브리지 또는 오버레이(기본값 bridge)
 - `--ip-range` 컨테이너에 할당하는 IP 주소 범위를 지정
 - `--subnet` 서브넷을 CIDR 형식으로 지정
 - `--ipv6` IPv6 네트워크를 유효화할 지말지(true/false)
 - `-label` 네트워크에 설정하는 라벨

```
docker network create --driver=bridge web-network
```

#### CIDR

Classless Inter-Domain Routing

`/24` : 24비트 이후에 오는 4번째 옥텟을 전부 사용할 수 있다는 표현

 - `192.168.0.0/24` : 192.168.0.0 ~ 192.168.0.255

<br><br><br>

### - 네트워크 연결

`docker network connect [옵션] 네트워크 컨테이너`

#### 옵션

 - `--ip` IPv4 주소
 - `--ip6` IPv6 주소
 - `--alias` 앨리어스명
 - `--link` 다른 컨테이너에 대한 링크


#### `my-server`라는 컨테이너를 `web-network`라는 도커 네트워크에 연결시키기

```
docker network connect web-network my-server
```

#### 컨테이너 실행하면서 네트워크에 연결

```
docker container run -itd --name=webapp --net=web-network my-container
```

#### 네트워크 연결 해제

`docker network disconnect`

```
docker network disconnect web-network webapp
```

<br><br><br>

### - 네트워크 상세 정보 확인

`docker network inspect [옵션] 네트워크`

```
docker network inspect web-network
```

#### 컨테이너 네트워크 확인

```js
docker container inspect sample
```

네트워크를 명시적으로 지정하지 않으면 기본값으로 bridge를 잡음

<br><br><br>

### - 네트워크 삭제(docker network rm)

`docker network rm [옵션] 네트워크`

```
docker network rm web-server
```

<hr><br><br><br><br><br>

## # 가동 중인 컨테이너 조작

### - 가동 컨테이너 연결 `docker container attach`

```
docker container attach sample
```

 - 연결한 컨테이너를 종료: `[CTRL] + C`
 - 컨테이너에서 분리: `[CTRL] + P, [CTRL] + Q`

<br><br><br>

### - 가동 컨테이너에서 프로세스 실행 `docker container exec`

`docker container exec [옵션] <컨테이너 식별자> <실행할 명령> [인수]`

#### 옵션

 - `--detach, -d` 명령을 백그라운드에서 실행한다
 - `--interactive, -i` 컨테이너 표준 입력을 연다
 - `--tty, -t` tty 사용
 - `--user, -u` 사용자명 지정

```
docker container exec -it webserver /bin/echo "Hello World"
```

<br><br><br>

### - 가동 컨테이너의 프로세스 확인 `docker container top`

```
docker container top sebserver
```

<br><br><br>

### - 가동 컨테이너의 포트 전송 확인 `docker container port`

```
docker container port webserver
```

<br><br><br>

### - 컨테이너 이름 변경 `docker container rename`

```
docker container rename old new
```


<br><br><br>

### - 컨테이너 안의 파일을 복사 `docker container cp`

`docker container cp <컨테이너 식별자>:<컨테이너 안의 파일 경로> <호스트의 디렉토리 경로>`

`docker container cp <호스트의 디렉토리 경로> <컨테이너 식별자>:<컨테이너 안의 파일 경로>`

```
docker container cp ./test.txt webserver:/tmp/test.txt
```

<br><br><br>

### - 컨테이너 조작의 차분 확인 `docker container diff`

`docker container diff <컨테이너 식별자>`

컨테이너 안에서 어떤 조작을 하여 컨테이너가 이미지로부터 생성되었을 때와 달라진 점을 확인

#### 변경의 구분
 - `A` 파일 추가
 - `D` 파일 삭제
 - `C` 파일 수정

<hr><br><br><br><br><br>

## # 도커 이미지 생성

도커 컨테이너는 도커 이미지를 바탕으로 작성되지만

반대로 컨테이너를 바탕으로 이미지를 작성할 수도 있음

### - 컨테이너로부터 이미지 작성

`docker container commit [옵션] <컨테이너 식별자> [이미지명[:태그명]]`

#### 옵션
 - `--author, -a` 작성자를 지정함
 - `--message, -m` 메시지를 지정함
 - `--change, -c` commit시 Dockerfile 명령을 지정
 - `--pause, -p` 컨테이너를 일시정지하고 commit

```
docker container commit -a "rhie-coder" webserver rhie/webapp:1.0
```

<br><br><br>

### - 컨테이너를 tar 파일로 출력

`docker container export <컨테이너 식별자>`

```sh
docker container export webserver > latest.tar
tar -tf latest.tar | more #확인하기
```

<br><br><br>

### - 이미지 저장(from tar파일)

`docker image save [옵션] <저장 파일명> [이미지명]`

#### 옵션
 - `-o` 저장할 파일명

```
docker image save -o export.tar tensorflow
```

<br><br><br>


### - 이미지 읽어 들이기

`docker image load [옵션]`

#### 옵션
 - `-i` 읽어 들일 파일명

```
docker image load -i export.tar
```

<br><br><br>

### - export/import & save/load 차이

#### docker container export

```
# docker container export my-web-server > export.tar

# tar xvf export.tar

total 141072
drwxr-xr-x 22 rhiemh rhiemh      4096 Dec 15 21:37 ./
drwxr-xr-x  4 rhiemh rhiemh      4096 Dec 15 21:37 ../
-rwxr-xr-x  1 rhiemh rhiemh         0 Dec 15 21:34 .dockerenv*
drwxr-xr-x  2 rhiemh rhiemh      4096 Dec  5 09:00 bin/
drwxr-xr-x  2 rhiemh rhiemh      4096 Sep  3 21:10 boot/
drwxr-xr-x  4 rhiemh rhiemh      4096 Dec 15 21:34 dev/
drwxr-xr-x  2 rhiemh rhiemh      4096 Dec 14 10:20 docker-entrypoint.d/
-rwxr-xr-x  1 rhiemh rhiemh      1616 Dec 14 10:19 docker-entrypoint.sh*
drwxr-xr-x 34 rhiemh rhiemh      4096 Dec 15 21:34 etc/
-rw-r--r--  1 rhiemh rhiemh 144361472 Dec 15 21:37 export.tar
drwxr-xr-x  2 rhiemh rhiemh      4096 Sep  3 21:10 home/
drwxr-xr-x  8 rhiemh rhiemh      4096 Dec  5 09:00 lib/
drwxr-xr-x  2 rhiemh rhiemh      4096 Dec  5 09:00 lib64/
drwxr-xr-x  2 rhiemh rhiemh      4096 Dec  5 09:00 media/
drwxr-xr-x  2 rhiemh rhiemh      4096 Dec  5 09:00 mnt/
drwxr-xr-x  2 rhiemh rhiemh      4096 Dec  5 09:00 opt/
drwxr-xr-x  2 rhiemh rhiemh      4096 Sep  3 21:10 proc/
drwx------  2 rhiemh rhiemh      4096 Dec  5 09:00 root/
drwxr-xr-x  3 rhiemh rhiemh      4096 Dec 15 21:34 run/
drwxr-xr-x  2 rhiemh rhiemh      4096 Dec  5 09:00 sbin/
drwxr-xr-x  2 rhiemh rhiemh      4096 Dec  5 09:00 srv/
drwxr-xr-x  2 rhiemh rhiemh      4096 Sep  3 21:10 sys/
drwxr-xr-x  2 rhiemh rhiemh      4096 Dec 14 10:20 tmp/
drwxr-xr-x 11 rhiemh rhiemh      4096 Dec  5 09:00 usr/
drwxr-xr-x 11 rhiemh rhiemh      4096 Dec  5 09:00 var/
```

export로 만들어진 tar은 import로 불러오기


#### docker image save

레이어 구조 포함

```
# docker image save -o save.tar nginx

# tar -xvf save.tar

36bb53b48a5a488898b13f2d1aa8b283970fa2c791771ea973710be659583b7f/
36bb53b48a5a488898b13f2d1aa8b283970fa2c791771ea973710be659583b7f/VERSION
36bb53b48a5a488898b13f2d1aa8b283970fa2c791771ea973710be659583b7f/json
36bb53b48a5a488898b13f2d1aa8b283970fa2c791771ea973710be659583b7f/layer.tar
37e0d9913a44679ff2dd6709586c94269c0309dcc7d1f1a15e9a4a8cf641eacf/
37e0d9913a44679ff2dd6709586c94269c0309dcc7d1f1a15e9a4a8cf641eacf/VERSION
37e0d9913a44679ff2dd6709586c94269c0309dcc7d1f1a15e9a4a8cf641eacf/json
37e0d9913a44679ff2dd6709586c94269c0309dcc7d1f1a15e9a4a8cf641eacf/layer.tar
3964ce7b84589cf9bc585415741b642b7167229e0afde03f502f6c848ed3279d.json
44b545a66656a392b718c5557424c8cc73bfeb9856a041d5d3e5eaf8ad8b36fb/
44b545a66656a392b718c5557424c8cc73bfeb9856a041d5d3e5eaf8ad8b36fb/VERSION
44b545a66656a392b718c5557424c8cc73bfeb9856a041d5d3e5eaf8ad8b36fb/json
44b545a66656a392b718c5557424c8cc73bfeb9856a041d5d3e5eaf8ad8b36fb/layer.tar
731fe0ea4745f8f281b41b1e5816f52aab34fde6e1385d0c67b1e01fc4cee6cf/
731fe0ea4745f8f281b41b1e5816f52aab34fde6e1385d0c67b1e01fc4cee6cf/VERSION
731fe0ea4745f8f281b41b1e5816f52aab34fde6e1385d0c67b1e01fc4cee6cf/json
731fe0ea4745f8f281b41b1e5816f52aab34fde6e1385d0c67b1e01fc4cee6cf/layer.tar
75e637380ce0ec4fbfe15027e07d043469116284482d847bf96dc46904cef37b/
75e637380ce0ec4fbfe15027e07d043469116284482d847bf96dc46904cef37b/VERSION
75e637380ce0ec4fbfe15027e07d043469116284482d847bf96dc46904cef37b/json
75e637380ce0ec4fbfe15027e07d043469116284482d847bf96dc46904cef37b/layer.tar
a9c099edbf24136c64c75e439b4e198e2b0e0b4a1f03ed51670941fe649037c3/
a9c099edbf24136c64c75e439b4e198e2b0e0b4a1f03ed51670941fe649037c3/VERSION
a9c099edbf24136c64c75e439b4e198e2b0e0b4a1f03ed51670941fe649037c3/json
a9c099edbf24136c64c75e439b4e198e2b0e0b4a1f03ed51670941fe649037c3/layer.tar
manifest.json
repositories
```
save로 만들어진 tar은 load로 불러오기


<br><br><br>

### - 불필요한 이미지/컨테이너를 일괄 삭제

`docker system prune [옵션]`

#### 옵션
 - `--all, -a` 사용하지 않는 리소스를 모두 삭제한다.
 - `--force, -f` 강제적으로 삭제한다.

