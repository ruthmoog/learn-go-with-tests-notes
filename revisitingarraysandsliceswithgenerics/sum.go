package main

// Sum calculates the total from a slice of numbers.
func Sum(numbers []int) int {
	add := func(result, item int) int { return result + item }
	return Reduce(numbers, add, 0)
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

func Reduce[A any](collection []A, function func(A, A) A, starterValue A) A {
	var result = starterValue
	for _, x := range collection {
		result = function(result, x)
	}
	return result
}
