package main

// Sum calculates the total from a slice of numbers.
func Sum(numbers []int) int {
	add := func(result, item int) int { return result + item }
	return Reduce(numbers, add, 0)
}

func SumAll(numbers ...[]int) []int {
	add := func(result, items []int) []int {
		return append(result, Sum(items))
	}
	return Reduce(numbers, add, []int{})
}

// SumAllTails calculates the sums of all but the
// first number given a collection of slices.
func SumAllTails(numbers ...[]int) []int {
	sumTail := func(result, collection []int) []int {
		if len(collection) == 0 {
			return append(result, 0)
		} else {
			tail := collection[1:] // everything after the 0 index item
			return append(result, Sum(tail))
		}
	}
	return Reduce(numbers, sumTail, []int{})
}

func Reduce[A any](collection []A, function func(A, A) A, starterValue A) A {
	var result = starterValue
	for _, x := range collection {
		result = function(result, x)
	}
	return result
}
