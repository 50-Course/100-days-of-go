package iteration

import (
	"testing"
)

// Helper function that does assert two arguments are the same
func assertEqual(t testing.TB, got, expected string) {
	t.Helper()
	if got != expected {
		t.Errorf("AssertionError: Expected %v but got %v\n", expected, got)
	}
}

func TestRepeat(t *testing.T) {
	t.Run("given a value, repeat it five times", func(t *testing.T) {
		// setup the test data

		repeated := Repeat("a")
		want := "aaaaa"

		assertEqual(t, repeated, want)
	})

	t.Run("given a value, repeat N times, as provided by the caller", func(t *testing.T) {
		repeated := Repeat("a", 2)
		want := "aa"

		assertEqual(t, repeated, want)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
