package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	var resBody interface{}
	resBody = "hello"
	resContents := fmt.Sprintf("%v", resBody)
	fmt.Println(resContents)
	typecheck := reflect.TypeOf(resContents)
	fmt.Println(typecheck)
}

func test() {
	option := "send,receive"
	result := strings.Split(option, ",")
	fmt.Println(result)
	fmt.Println(contains(result, "sed"))
}

func contains(s []string, substr string) bool {
	for _, v := range s {
		if v == substr {
			return true
		}
	}
	return false
}
