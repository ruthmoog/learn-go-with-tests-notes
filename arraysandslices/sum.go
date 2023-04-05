package main

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbers ...[]int) []int {
	size := len(numbers)
	sums := make([]int, size)

	for i, collection := range numbers {
		sums[i] = Sum(collection)
	}
	return sums
}
