package main

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbers ...[]int) []int {
	var sums []int
	for _, collection := range numbers {
		sums = append(sums, Sum(collection))
	}
	return sums
}
