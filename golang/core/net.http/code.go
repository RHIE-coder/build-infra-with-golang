package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func USAGE_HTTP_HANDLER() {
	handler := func(rw http.ResponseWriter, req *http.Request) {
		fmt.Println("Method : ", req.Method)
		fmt.Println("URL : ", req.URL)
		fmt.Println("Header : ", req.Header)

		b, _ := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		fmt.Println("Body : ", string(b))

		switch req.Method {
		case "POST":
			rw.Write([]byte(`{"result":"success", "message": "post request success !"}`))
		case "GET":
			rw.Write([]byte("get request success !"))
		}
	}

	/*
	   curl -X POST "http://localhost:5000/api/point/mint" -H "Content-Type: application/json" -H "access-token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDYxODU0MTksImlhdCI6MTY3NDY0OTQxOSwiaXNzIjoiY2hvc3VuOmJsb2NrY2hhaW46YnJva2VyIiwic3ViIjoiY2hvc3VuYmMifQ.8fR_7ADmCIYX8PAiQMK88iY8wYDwFP9ced4jViqLFmk" -H "uuid: \"9cb251ca-fe4a-4d96-8ee5-a7bc623ec250\"" --data '{
	       "address": "0xd1104e5ab60ae1573f59919e6e089d63d01ba3bc",
	         "value": 11
	       }'
	*/

	requsetToServer := func() {
		time.Sleep(5 * time.Second)
		bodyData := []byte(`{"name": "John Doe", "age": 30}`)
		body := bytes.NewBuffer(bodyData)
		req, _ := http.NewRequest("POST", "http://localhost:8000", body)
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		var data map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			panic(err)
		}
		json.MarshalIndent(data, "", "  ")
		fmt.Println(data)
	}
	go requsetToServer()
	err := http.ListenAndServe(":8000", http.HandlerFunc(handler))
	if err != nil {
		fmt.Println("Failed to ListenAndServe : ", err)
	}

}

func main() {
	USAGE_HTTP_HANDLER()
}
