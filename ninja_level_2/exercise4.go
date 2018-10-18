package main

import (
	"fmt"
)

func main () {
	i := 42
	fmt.Printf("%b\t#%x\t%d\n", i, i, i)
	j := i << 1
	fmt.Printf("%b\t#%x\t%d\n", j, j, j)
}