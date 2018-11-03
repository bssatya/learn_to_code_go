package main

// func Sum takes five integers and returns the sume of all the 5 integers
func Sum(numbers []int) (sum int) {
	for _, n := range numbers {
		sum += n
	}
	return
}

// func SumAll takes variadic slices of int and returns a slice containing the sum of each slice
func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return
}

// func SumAllTails takes variadic slices of int and returns a slice containing the sum of each slice tails
func SumAllTails(slicesToSum ...[]int) (tailSums []int) {
	numberOfSlices := len(slicesToSum)
	tailSums = make([]int, numberOfSlices)

	for i, slice := range slicesToSum {
		var sum int
		if len(slice) == 0 {
			sum = 0
		} else {
			tail := slice[1:]
			sum = Sum(tail)
		}
		tailSums[i] = sum
	}
	return
}
