package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

type RequestBody struct {
	Result  string `json:"result"`
	Message string `json:"message"`
}

func USAGE_HTTP_HANDLER() {
	handler := func(rw http.ResponseWriter, req *http.Request) {
		fmt.Println("Method : ", req.Method)
		fmt.Println("URL : ", req.URL)
		fmt.Println("Header : ", req.Header)

		// b, _ := ioutil.ReadAll(req.Body)
		var buf bytes.Buffer
		newBody := io.TeeReader(req.Body, &buf)

		// TODO
		// fmt.Printf("%p\n", newReqBody)
		// fmt.Printf("%p\n", req.Body)

		b, _ := ioutil.ReadAll(req.Body)

		reqBody := RequestBody{}
		err := json.NewDecoder(newBody).Decode(&reqBody)

		if err != nil {
			panic(err)
		}

		defer req.Body.Close()
		fmt.Println("Body : ", string(b))
		b2, _ := json.Marshal(reqBody)
		fmt.Println("Body2: ", string(b2))

		switch req.Method {
		case "POST":
			fmt.Println("res data is written")
			// rw.Write([]byte(`{"result":"success", "message": "post request success !"}`))
			fmt.Fprint(rw, `{"result":"success", "message": "post request success !"}`)
			fmt.Println("wait start")
			// buf := rw.(*bytes.Buffer) //*bytes.Buffer does not implement http.ResponseWriter (missing method Header)
			// responseString := buf.String()
			// fmt.Println(responseString)
			time.Sleep(3 * time.Second)
			fmt.Println("wait end")
		case "GET":
			rw.Write([]byte("get request success !"))
		}
	}

	reqToServer := func() {
		time.Sleep(3 * time.Second)
		bodyData := []byte(`{"name": "John Doe", "age": 30}`)
		body := bytes.NewBuffer(bodyData)
		req, _ := http.NewRequest("POST", "http://localhost:8000", body)
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		fmt.Println("request!!")
		resp, err := client.Do(req)
		fmt.Println("response!!")
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
	go reqToServer()
	err := http.ListenAndServe(":8000", http.HandlerFunc(handler))
	if err != nil {
		fmt.Println("Failed to ListenAndServe : ", err)
	}
}

func USAGE_IP() {
	ipAddrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatalf("fail to get ip address")
	}
	for _, addr := range ipAddrs {
		ipnet, ok := addr.(*net.IPNet)
		if !ok {
			log.Fatalf("fail to assert")
		}
		fmt.Println(ipnet.IP)
		fmt.Println(ipnet.Mask)
		fmt.Println(ipnet.Mask.Size())
		fmt.Println(ipnet.IP.IsLoopback())
		fmt.Println(ipnet.IP.To4())
		fmt.Println(ipnet.IP.To16())
		fmt.Println("=====")
	}
}

func USAGE_CIDR() {
	ip := net.ParseIP("172.23.37.104")
	mask := net.CIDRMask(20, 32)

	ipnet := &net.IPNet{
		IP:   ip,
		Mask: mask,
	}

	cidr := ipnet.String()
	fmt.Println(cidr)
}

func main() {
	// USAGE_IP()
	USAGE_CIDR()
	// USAGE_HTTP_HANDLER()
}
