package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/valyala/fasthttp"
)

type FasthttpClient struct {
	origin        string
	beforeRequest func(*fasthttp.Request)
	afterResponse func(*fasthttp.Response)
}

type Requester struct {
	Method string
	URL    string
	Data   interface{}
}

func NewClient(origin string) *FasthttpClient {
	return &FasthttpClient{
		origin: origin,
	}
}

func (client *FasthttpClient) SetBeforeRequest(hookFunc func(*fasthttp.Request)) {
	// req.Header.SetContentType("application/json")
	client.beforeRequest = hookFunc
}

func (client *FasthttpClient) SetAfterResponse(hookFunc func(*fasthttp.Response)) {
	client.afterResponse = hookFunc
}

func (client *FasthttpClient) Dial(requester *Requester) {
	var err error

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	if client.beforeRequest != nil {
		client.beforeRequest(req)
	}

	req.Header.SetMethod(requester.Method)
	req.Header.SetRequestURI(client.origin + requester.URL)

	var reqBody []byte
	if requester.Data != nil {
		reqBody, err = json.Marshal(requester.Data)
	}

	if err != nil {
		panic(err)
	}

	if len(reqBody) > 0 {
		req.SetBody(reqBody)
	}

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := fasthttp.Do(req, resp); err != nil {
		panic(err)
	}

	if client.afterResponse != nil {
		client.afterResponse(resp)
	}

	var resData map[string]interface{}
	json.Unmarshal(resp.Body(), &resData)
	JsonIndent2, err := json.MarshalIndent(resData, "", "  ")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(JsonIndent2))

	// 파일 생성 및 쓰기
	file, err := os.Create("result.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = file.Write(JsonIndent2)
	if err != nil {
		fmt.Println(err)
		return
	}
}
