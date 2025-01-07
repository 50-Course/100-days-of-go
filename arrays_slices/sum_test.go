package arrays_slices

import "testing"

func TestSum(t *testing.T) {
	nums := [5]int{1, 2, 3, 4, 5}

	got := Sum(nums)
	want := 15

	if got != want {
		t.Errorf("AssertionError: expected %d for an answer but received %d, given %v", want, got, nums)
	}
}
