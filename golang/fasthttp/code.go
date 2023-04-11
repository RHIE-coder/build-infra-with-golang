package main

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

// 요청 바디에 담을 데이터를 담는 구조체
type User struct {
	Username string `json:"username"`
	Age      int    `json:"age"`
	Hobby    string `json:"hobby"`
}

func main() {
	// TODO: Connection Pool 설정
	// fasthttp 객체 풀 설정 (unvalid)
	// fasthttp.ConfigureConnPool(fasthttp.ConnPoolConfig{
	// 	MaxConns:     1024,
	// 	MaxIdleTime:  10 * time.Second,
	// 	DialTimeout:  2 * time.Second,
	// 	ReadTimeout:  5 * time.Second,
	// 	WriteTimeout: 5 * time.Second,
	// })

	// POST 요청을 보낼 API 주소
	url := "http://example.com/example"

	// 요청 바디에 담을 데이터
	user := User{
		Username: "john_doe",
		Age:      30,
		Hobby:    "reading",
	}

	// 요청 바디를 JSON 형식으로 인코딩
	reqBody, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	// fasthttp 객체 풀에서 Conn 객체 가져오기
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	// 요청 메서드와 경로 설정
	req.Header.SetMethod("POST")
	req.Header.SetRequestURI(url)

	// 요청 헤더 설정
	req.Header.SetContentType("application/json")
	req.Header.SetContentLength(len(reqBody))

	// 요청 바디 설정
	req.SetBody(reqBody)

	// fasthttp 객체 풀에서 Conn 객체 가져오기
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	// POST 요청 보내기
	if err := fasthttp.Do(req, resp); err != nil {
		panic(err)
	}

	// 응답 결과 출력
	fmt.Println(string(resp.Body()))
}
