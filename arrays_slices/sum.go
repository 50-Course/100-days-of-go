package arrays_slices

func Sum(arr [5]int) int {
	count := 0

	for _, n := range arr {
		count += n
	}
	return count
}
