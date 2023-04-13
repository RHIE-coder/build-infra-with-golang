package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"

	"golang/external/fiber/swagger"
)

func selectExec() string {
	var exec string
	flag.StringVar(&exec, "exec", "", "the execution entrypoint name")

	flag.Parse()
	// flag 값이 없으면 에러 처리
	if flag.Lookup("exec").Value.String() == "" {
		flag.PrintDefaults()
		panic("the exec option is required")
	}

	return exec
}

func main() {

	entrypoint := selectExec()

	app := fiber.New()

	// TODO: make entrypoint array for multiple execution
	switch entrypoint {
	case "swagger":
		swagger.App(app)
	}

	app.Listen(":5555")
}
