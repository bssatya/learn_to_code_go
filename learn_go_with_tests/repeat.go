package main

// func Repeat takes a string and returns a string which repeated form of the input string
func Repeat(s string, times int) (repeat string) {
	for i := 0; i < times; i++ {
		repeat += s
	}

	return
}
