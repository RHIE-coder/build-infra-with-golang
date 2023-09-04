package internal

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

func HttpRequester(method string, url string, data interface{}) ([]byte, error) {

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.Header.SetMethod(method)
	req.Header.SetRequestURI(url)
	req.Header.SetContentType("application/json")

	if data != nil {
		reqBody, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		req.Header.SetContentLength(len(reqBody))
		req.SetBody(reqBody)
	}

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := fasthttp.Do(req, resp); err != nil {
		return nil, err
	}

	return resp.Body(), nil
}
