package main

import "fmt"

func AddressCheck() {
	arr := []int{1, 2, 3, 4}
	input := []int{5, 6, 7, 8}
	arr = append(arr, input...)
	fmt.Printf("%p\n", &arr[4])
	fmt.Printf("%p\n", &input[0])
	fmt.Println(input[0])
	fmt.Println(arr[4])

	// input[0] != arr[4]
}

func Cut() {
	limit := 6
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr[:limit], len(arr[:limit]))
}

func main() {
	Cut()
}
