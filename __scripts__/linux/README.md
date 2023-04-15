# 위치 확인

## [type]

```sh
[bash$] type ifconfig
ifconfig is /sbin/ifconfig
```

## [which]

```sh
[bash$] which ifconfig
/sbin/ifconfig
```

## [whereis]

```sh
[bash]# whereis ifconfig
ifconfig: /sbin/ifconfig /usr/share/man/man8/ifconfig.8.gz
```

## [find]

```sh
[bash]# find / -name ifconfig
/sbin/ifconfig
```

<br><br><br><br><br>

# 인자를 기억해서 여러 터미널에 적용시키기

## [STEP 1 : 명령어 기억하기]
 - `>`  명령어: `w`
 - `>>` 명령어: `w+`
 - `readlink`: 상대 경로 --> 절대 경로

```sh
// copy path for sharing
[bash]# readlink -f ./workspace/notebook > cpfs
[bash]# cat cpfs
/home/rhiemh/workspace/notebook
```

## [STEP 2 : 명령어 실행 - 실패]

### - 원인 

 - xargs가 파이프라인 왼쪽 결과를 공백을 기준으로 $1, $2, $3 으로 나눈다
 - xargs가 나눈 결과를 bash -c 명령어에 넣는다
 - $0은 맨 오른쪽의 bash가 된다.

```sh
[bash]# echo 1 2 3  | xargs bash -c 'echo "$1"' bash
1
[bash]# echo 1 2 3  | xargs bash -c 'echo "$2"' bash
2
[bash]# echo 1 2 3  | xargs bash -c 'echo "$3"' bash
3
```

 - 이렇게도 쓰일 수 있다

```sh
[bash]# cat cpfs | xargs -I {} bash -c 'echo "{}"'
/home/rhiemh/workspace/notebook
```

### - 실패할 수 밖에 없는 이유

 - 현재 쉘이 아니라 다른 쉘 프로세스에서 돌아가기 때문

```sh
// cpfs에 명시된 경로에 파일은 생성되나 현재 쉘에는 영향이 없다
[bash]# cat ~/cpfs | xargs -I {} bash -c 'cd "{}" && touch testtest' bash
```

## [STEP 3: 완전 간단헀었다]

```sh
cd $(cat ~/cpfs)
```

끝...

<br><br><br><br><br>