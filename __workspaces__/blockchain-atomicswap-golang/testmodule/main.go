package main

import (
	"fmt"

	memory "github.com/gofiber/storage/memory/v2"
)

func main() {

	store := memory.New()

	err := store.Set("key1", []byte("test"), 100000000000)

	if err != nil {
		panic(err)
	}

	result, err := store.Get("key1")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(result)) //test
}
