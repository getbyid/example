package main

import (
	"fmt"
	"runtime"
)

var (
    version = "dev"
    commit  = "none"
    date    = "unknown"
)

func main() {
	fmt.Println("hello world, github! (", version, commit, date, ")")
}
