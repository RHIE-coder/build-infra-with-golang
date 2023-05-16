# Dockerfile

## # 개념

Docker 상에서 작동시킬 컨테이너의 구성 정보를 기술하기 위한 파일

`docker build` 명령어를 통해 Dockerfile에 기술된 구성 정보를 바탕으로 Docker 이미지 작성

## # 기본 구문

 - `FROM`: 베이스 이미지 지정
 - `RUN`: 명령 실행
 - `CMD`: 컨테이너 실행 명령
 - `LABEL`: 라벨 설정
 - `EXPOSE`: 포트 익스포트
 - `ENV`: 환경 변수
 - `ADD`: 파일/디렉토리 추가
 - `COPY`: 파일 복사
 - `ENTRYPOINT`: 컨테이너 실행 명령
 - `VOLUME`: 볼륨 마운트
 - `USER`: 사용자 지정
 - `WORKDIR`: 작업 디렉토리
 - `ARG`: Dockerfile 안의 변수
 - `ONBUILD`: 빌드 완료 후 실행되는 명령
 - `STOPSIGNAL`: 시스템 콜 시그널 설정
 - `HEALTHCHECK`: 컨테이너의 헬스 체크
 - `SHELL`: 기본 쉘 설정

