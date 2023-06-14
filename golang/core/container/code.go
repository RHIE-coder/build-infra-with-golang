package main

import (
	"container/list"
	"fmt"
)

func USAGE_LIST() {
	li := list.New()
	li.PushBack(10)  // 10
	li.PushBack(20)  // 10 -> 20
	li.PushBack(30)  // 10 -> 20 -> 30
	li.PushFront(40) // 40 -> 10 -> 20 -> 30
	li.PushFront(50) // 50 -> 40 -> 10 -> 20 -> 30

	for el := li.Front(); el != nil; el = el.Next() {
		fmt.Println(el)
		fmt.Println(el.Value)
	}
}

func main() {
	USAGE_LIST()
}
