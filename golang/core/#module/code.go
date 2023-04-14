package main

import "fmt"

func doSomething() {
	fmt.Println("do something")
}

func namedFactory() func() {
	innerFunc := func() {
		fmt.Println("this is named function")
	}

	return innerFunc
}

func anonyFactory() func() {
	return func() {
		fmt.Println("this is anonymous function")
	}
}

func printFuncPointer(fptr func()) {
	fmt.Println(&fptr)
}

func WHAT_IS_FUNC_POINT() {
	printFuncPointer(doSomething)
	printFuncPointer(doSomething)
	printFuncPointer(doSomething)
}

func main() {
	WHAT_IS_FUNC_POINT()
}
