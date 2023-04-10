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

func SumAllTails(numbers ...[]int) []int {
	var sums []int
	for _, collection := range numbers {
		if len(collection) == 0 {
			sums = append(sums, 0)
		} else {
			tail := collection[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
