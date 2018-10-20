package main

import (
	"fmt"
	"runtime"
)


func main() {
	fmt.Println("Runtime Arch:", runtime.GOARCH)
	fmt.Println("Runtime Os:", runtime.GOOS)
	fmt.Println("Runtime Num of CPU's:", runtime.NumCPU())
}