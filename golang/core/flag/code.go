package main

import (
	"flag"
	"fmt"
	"os"
)

/*
	옵션이 주어지지 않아도 기본값으로 할당되고 에러가 발생하지 않음
*/
func USAGE_NewFlagSet() {
	fs := flag.NewFlagSet("core", flag.ExitOnError)

	var port int
	fs.IntVar(&port, "port", 8080, "port number")

	err := fs.Parse(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("port:", port)
}

/*
	옵션이 주어지지 않으면 에러 발생
*/
func USAGE_METHODS_OF_flag() {
	var port int
	flag.IntVar(&port, "port", -1, "port number")

	flag.Parse()
	// flag 값이 없으면 에러 처리
	if flag.Lookup("port").Value.String() == "-1" {
		flag.PrintDefaults()
		fmt.Println("port is required")
		return
	}
}

// TODO: Visit()
// TODO: Value()

func main() {

}
