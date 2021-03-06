package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup

	printSysDetails()
	gr := 100
	wg.Add(gr)
	var m sync.Mutex

	for i := 0; i < gr; i++ {
		go func() {
			m.Lock()
			contr := counter
			contr++
			counter = contr
			fmt.Println("Counter = ", counter)
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter = ", counter)
}

func printSysDetails() {
	fmt.Println("Runtime Arch:", runtime.GOARCH)
	fmt.Println("Runtime Os:", runtime.GOOS)
	fmt.Println("Runtime Num of Go Routine:", runtime.NumGoroutine())
	fmt.Println("Runtime Num of CPU's:", runtime.NumCPU())
}
