package main

import (
	"fmt"
	"runtime"
)

var Version = "dev"
var BuildDate = "---"

func main() {
	fmt.Println("hello world, github! (", Version, BuildDate, runtime.GOOS, runtime.GOARCH, ")")
}
