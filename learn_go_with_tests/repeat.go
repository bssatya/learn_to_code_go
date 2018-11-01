package main

// func Repeat takes a string and returns a string which repeated form of the input string
func Repeat(s string) (repeat string) {
	for i := 0; i < 5; i++ {
		repeat += s
	}

	return
}
