package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	printSysDetails()
	wg.Add(2)

	fmt.Println("In Main: Before calling go routines")
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("InFoo Go Routine", i)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("InBar Go Routine", i)
		}
		wg.Done()
	}()
	printSysDetails()
	wg.Wait()
	fmt.Println("In Main: Done")
	printSysDetails()
}

func printSysDetails() {
	fmt.Println("Runtime Arch:", runtime.GOARCH)
	fmt.Println("Runtime Os:", runtime.GOOS)
	fmt.Println("Runtime Num of Go Routine:", runtime.NumGoroutine())
	fmt.Println("Runtime Num of CPU's:", runtime.NumCPU())
}
