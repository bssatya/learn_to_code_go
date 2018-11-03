package main

// func Sum takes five integers and returns the sume of all the 5 integers
func Sum(numbers []int) (sum int) {
	for _, n := range numbers {
		sum += n
	}
	return
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	lengthOfNumbers := len(numbersToSum)
	sums = make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return
}
