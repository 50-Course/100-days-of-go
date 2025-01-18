package arrays_slices

import (
	_ "errors"
	_ "fmt"
)

func Sum(arr []int) int {
	count := 0

	for _, n := range arr {
		count += n
	}
	return count
}

func SumAll(numsToSum ...[]int) []int {
	if numsToSum == nil {
		return []int{}
	}

	var sums []int

	for _, num := range numsToSum {
		sums = append(sums, Sum(num))
	}
	return sums
}

func SumAllTails(numsToSum ...[]int) []int {
	if numsToSum == nil {
		return []int{}
	}

	var sums []int

	// tails are slices or arrays beginning from the second element to the end
	// of the slice or array
	for _, nums := range numsToSum {
		tail := nums[1:]
		sums = append(sums, Sum(tail))
	}
	return sums
}
