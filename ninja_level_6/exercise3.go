package main

import (
	"fmt"
)

func main() {
	defer cleanUp()
	fmt.Println("I am in main function and right after calling cleanup function with defer keyword")
}

func cleanUp() {
	defer func() {
		fmt.Println("The cleanUp defer function ran")
	}()
	fmt.Println("In cleanUp function, Main exit successful, all clean-up done")
}