package main

import (
	"fmt"
	"reflect"
)

func USAGE_METHOD_OF_TypeOf() {
	type favContextKey string
	k := favContextKey("language")
	fmt.Println(reflect.TypeOf(k)) //main.favContextKey
}

func main() {
}
