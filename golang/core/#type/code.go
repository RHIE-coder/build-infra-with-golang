package main

import (
	"fmt"
	"reflect"
)

type alice string
type Stuff alice

func CallMe(who alice) {
	fmt.Println("hello " + who)
}

func GetData() string {
	return "alice"
}

func main() {
	var msg alice = alice(GetData())
	// var msg alice = interface{}(GetData()).(alice) // error
	fmt.Println(reflect.TypeOf(msg))
}
