package environment

import (
	"fmt"
	"runtime"
)

func LoadEnvFile() {
	_, execPath, _, _ := runtime.Caller(0)
	fmt.Println(execPath)
}
