package arrays_slices

import (
	"reflect"
	"slices"
	"testing"
)

// so my understanding now is all objects inherit from the `interface` 'Object'
//
// the concept of object or Object as reffered to as languages like python, java or TS
// does not exist in Go, instead there are interface which as as the topmost-level
// 'interface' (or let me call it object) for now, of which all other types, or 'objects'
// inherits from
func assertEqual(t testing.TB, got, want int) {
	t.Helper() // acts as python's functools.wraps

	if got != want {
		t.Errorf("AssertionError: expected %d for an answer but received %d", want, got)
	}
}

func TestSum(t *testing.T) {
	t.Run("collection of of 5 integers results in the expected value", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}

		got := Sum(nums)
		want := 15
		assertEqual(t, got, want)
	})

	t.Run("a collection of a combination of postive and negative integers results in an expected value", func(t *testing.T) {
		numsCombo := []int{1, -2, 3, -4, -5}

		got := Sum(numsCombo)
		// expected = -7
		want := -7

		assertEqual(t, got, want)
	})

	t.Run("A collection of any number of integers results in the expected value", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		got := Sum(nums)
		want := 55
		assertEqual(t, got, want)
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !slices.Equal(got, want) {
		t.Errorf("expected %v but received %v", want, got)
	}

	t.Run("zero variadic numbers of slices passed in will results in an empty slice", func(t *testing.T) {
		got := SumAll()
		want := []int{}

		if !slices.Equal(got, want) {
			t.Errorf("expected %v but received %v", want, got)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{2, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %v but received %v", want, got)
	}
}
