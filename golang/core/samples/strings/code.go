package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "The quick brown fox jumps over the lazy dog"
	substr1 := "brown"
	substr2 := "cat"
	fmt.Println(strings.Contains(str, substr1)) // true
	fmt.Println(strings.Contains(str, substr2)) // false
}
