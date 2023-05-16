# Cross-Origin Resource Sharing

브라우저에서 처리함(응답을 버림, 서버는 어쨋든 응답함).

CORS는 브라우저의 구현 스펙에 포함되는 정책이기 때문에 브라우저를 통하지 않고 서버간 통신을 할 때는 이 정책이 적용되지 않음

## [과정]

 1. 브라우저가 서버에게 요청보냄(Origin 헤더)
 2. 서버가 이 요청에 대한 응답시 `Access-Control-Allow-Origin` 헤더 보냄
 3. 브라우저가 Origin과 Access-Control-Allow-Origin 헤더 비교

## [Preflight Request]

 1. 브라우저는 요청을 한번에 보내지 않고 예비 요청과 본 요청으로 나누어서 서버로 전송
 2. HTTP METHOD: `OPTIONS`

<hr><br><br><br><br>

# [MDN 문서](https://developer.mozilla.org/ko/docs/Web/HTTP/CORS#%EC%A0%91%EA%B7%BC_%EC%A0%9C%EC%96%B4_%EC%8B%9C%EB%82%98%EB%A6%AC%EC%98%A4_%EC%98%88%EC%A0%9C)

## Process

 - preflighted가 적용되는 Content-Type 종류

```
application/json
application/xml
text/plain
multipart/form-data
application/x-www-form-urlencoded
application/octet-stream

By ChatGPT
```

 - Access-Control-Max-Age (en-US) 헤더는 preflight request 요청 결과를 캐시할 수 있는 시간을 나타냅니다