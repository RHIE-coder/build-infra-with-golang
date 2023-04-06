package main

import (
	"fmt"
	"reflect"
)

func main() {
	type favContextKey string
	k := favContextKey("language")
	fmt.Println(reflect.TypeOf(k)) //main.favContextKey
}
