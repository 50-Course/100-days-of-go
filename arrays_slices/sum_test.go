package arrays_slices

import "testing"

// so my understanding now is all objects inherit from the `interface` 'Object'
//
// the concept of object or Object as reffered to as languages like python, java or TS
// does not exist in Go, instead there are interface which as as the topmost-level
// 'interface' (or let me call it object) for now, of which all other types, or 'objects'
// inherits from
func assertEqual(t testing.TB, got interface{}, want int) {
	t.Helper() // acts as python's functools.wraps

	if got != want {
		t.Errorf("AssertionError: expected %d for an answer but received %d, given %v", want, got, nums)
	}
}

func TestSum(t *testing.T) {
	nums := [5]int{1, 2, 3, 4, 5}

	got := Sum(nums)
	want := 15

	// t.Run("Array of of 5 integers results in the expected value", func(t *testing.T) {
	// })
	// t.Run("Array of values greater than capacity of five results in an error", func(t *testing.T) {
	// })
	// t.Run("Arrays of a combination of postive and negative integers results in an expected value", func(t *testing.T) {
	// })
	// t.Run("Array of all negative values added together results in a negative value", func(t *testing.T) {
	// })
}
