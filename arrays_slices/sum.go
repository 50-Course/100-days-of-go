package arrays_slices

import (
	_ "errors"
	_ "fmt"
)

func Sum(arr [5]int) int {
	count := 0

	if len(arr) > 5 {
	}

	for _, n := range arr {
		count += n
	}
	return count
}
