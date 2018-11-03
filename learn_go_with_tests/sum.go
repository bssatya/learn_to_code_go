package main

// func Sum takes five integers and returns the sume of all the 5 integers
func Sum(numbers []int) (sum int) {
	for _, n := range numbers {
		sum += n
	}
	return
}
