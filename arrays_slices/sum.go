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
