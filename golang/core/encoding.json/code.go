package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	str := `response status code : 400, body : {"status":400,"result":false,"message":"The 123123 has already been used."}`
	var data map[string]interface{}
	err := json.Unmarshal([]byte(str[strings.Index(str, "{"):]), &data)
	if err != nil {
		panic(err)
	}
	message, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(strings.Contains(data["message"].(string), "has already been used."))
	fmt.Println(string(message))
}
