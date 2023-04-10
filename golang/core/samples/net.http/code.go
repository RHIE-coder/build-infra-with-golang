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
