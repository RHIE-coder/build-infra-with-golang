# 밴치마킹

## # 설치

```
git clone https://github.com/wg/wrk
cd wrk
make
```

## # 실행

```
./wrk -t12 -c400 -d30s http://192.168.56.1:3344/

Running 30s test @ http://192.168.56.1:3344/
  12 threads and 400 connections
```

### - 기본값

```
Running 10s test @ http://192.168.56.1:3344
  2 threads and 10 connections
```