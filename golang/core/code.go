package main

import (
	"fmt"
	"strings"
)

func main() {
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
