# Docker

## # Docker의 기능

 - Docker 이미지를 만드는 기능:build
 - Docker 이미지를 공유하는 기능:ship
 - Docker 컨테이너를 작동시키는 기능:run

### - build

Docker 이미지의 정체는 어플리케이션의 실행에 필요한 파일들이 저장된 디렉토리.

수동으로 만들거나 Dockerfile을 이용해 만듬.

Docker 이미지는 겹쳐서 사용할 수 있음.

<br><br><br>

### - ship

Docker 이미지는 Docker Registry에 공유할 수 있음(Docker Hub)

#### Automated Build

Docker Hub와 Github와 연동하여 Github 상에서는 Dockerfile을 관리하고 Docker 이미지를 자동으로 생성하여 Docker Hub에 공개

<br><br><br>

### - run

Docker 이미지만 있다면 Docker가 설치된 환경이라면 어디서든지 컨테이너를 작동시킬 수 있음.

또한, Docker 이미지를 가지고 여러개의 컨테이너를 기동시킬 수도 있음.

제품 환경에서는 모든 Docker 컨테이너를 한 대의 호스트 머신(물리 머신)에서 작동시키는 일은 드물며, 시스템의 트래픽 증감이나 가용성 요건, 신뢰성 요건 등을 고려한 후에 여러 대의 호스트 머신으로 된 분산 환경을 구축.

컨테이너 관리에 대해서 오케스트레이션 툴을 이용하는 것이 일반적



 <hr><br><br><br><br><br>

## # Docker 컴포넌트

### - Docker Engine 

도커의 핵심 기능. 이미지 생성, 컨테이너 기동, Dockerfile에 의한 이미지 생성 등

<br><br><br>

### - Docker Registry

Docker 이미지를 공개 및 공유하기 위한 레지스트리 기능

Docker Hub도 Docker Registry 기능을 사용하고 있음.

<br><br><br>

### - Docker Compose

컨테이너 일원 관리

<br><br><br>

### - Docker Machine

Docker 실행 환경 구축

<br><br><br>

### - Docker Swarm

클러스터 관리

여러 Docker 호스트를 클러스트화하기 위한 툴

Docker Swarm에서는 클러스터를 관리하거나 API를 제공하는 역할을 Manager가 담당

Docker 컨테이너를 실행하는 역할을 Node가 담당.

K8S도 이용할 수 있음

<hr><br><br><br><br><br>

## # Docker의 작동 구조

### - namespace: 컨테이너 구획화하기

#### PID namespace

PID: Linux에서 각 프로세스에 할당된 고유 ID이며 namespace가 다른 프로세스끼리는 서로 엑세스할 수 없음

#### Network namespace

네트워크 리소스(IP, Port, Routing Table 등)를 격리된 namespace마다 독립적으로 가질 수 있음

덕분에, 호스트OS 상에 이미 사용중인 포트가 있더라도 컨테이너 안에서 동일한 포트 번호를 사용할 수 있음

#### UID namespace

UID, GID 별로 독립적으로 가짐

#### MOUNT namespace

마운트란 컴퓨터에 연결된 기기나 기억장치를 OS에 인식시켜 이용가능한 상태로 만드는 것을 말함

격리된 파일시스템을 만들 수 있음.

### UTS namespace

namespace별로 독자적으로 호스트명이나 도메인명을 가질 수 있음

### IPC namespace

프로세스 간의 통신 오브젝트를 namespace별로 독립적으로 가질 수 있음

 <hr><br><br><br><br><br>

## # 릴리즈 관리 장치(cgroups)

Docker는 물리 머신 상의 자원을 여러 컨테이너가 공유하여 작동한다.

이 떄 Linux 커널의 기능인 contrl groups 기능을 사용하여 자원의 할당 등을 관리한다.

 - `cpu`: CPU 사용량 제한
 - `cpuacct`: CPU 사용량 통계 정보를 제공
 - `cpuset`: CPU나 메모리 배치를 제어
 - `memory`: 메모리나 스왑 사용량을 제한
 - `devices`: 디바이스에 대한 액세스 허가/거부
 - `freezer`: 그룹에 속한 프로세스 정지/재개
 - `net_cls`: 네트워크 제어 태그를 부가
 - `blkio`: 블록 디바이스 입출력량 제어

 <hr><br><br><br><br><br>

## # 네트워크 구성(가상 브리지/가상 NIC)

리눅스는 DOcker를 설치하면 서버의 물리 NIC가 docker0이라는 가상 브리지 네트워크로 연결됨.

도커 컨테이너가 실행되면 컨테이너에 172.17.0.0/16이라는 서브넷 마스크를 가진 프라이빗 IP주소가 eth0으로 자동으로 할당됨

### - NAT(Network Address Translation)

 프라이빗 IP --> NAT가 가지고 있는 글로벌 주소

 주소 변환에에 의해 프라이빗 네트워크상의 컴퓨터와 인터넷 상의 서버간의 통신이 성립

 but, 글로벌 IP 주소와 프라이빗 IP 주소를 1:1로 변환하기 때문에 동시에 여러 클라이언트가 엑세스 할 수 없음

### - NAPT(Network Address Port Translation)

 프라이빗 IP 주소와 함께 포트 번호도 같이 변환하는 기술

 하나의 글로벌 IP주소와 여러 개의 프라이빗 IP 주소를 변환할 수 있음

 Linux에서 NAPT를 구축하는 것을 IP 마스커레이드(mascarade)라고 함-가면무도회

 <hr><br><br><br><br><br>

## # Docker 활용하기

### - 버전 확인

`docker version`

### - 도커 실행 환경 확인

`docker system info`

실행 환경의 상세 설정이 표시됨

<br><br><br>

### - 도커 디스크 이용 상황

`docker system df`

<br><br><br>

### - 도커 이미지 다운로드

`docker pull <image>`

```sh
docker pull nginx
```

#### 도커 이미지를 사용하여 가동 및 정지

```sh
$ docker container run --name webserver -d -p 80:80 nginx  # 이미지 --> 컨테이너
$ docker container ps # 컨테이너 확인
$ docker container stats webserver # 컨테이너 가동 확인
$ docker stop webserver # 컨테이너 정지
$ docker start webserver # 컨테이너 가동
```




