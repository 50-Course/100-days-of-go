package integers

import (
	"fmt"
	"testing"
)

// This introduces testable examples that can be loaded in
// go doc. The idea is to write a function with the name
// "Example"-followed by the original function name IN
//
// ...the test file
func ExampleAdd() {
	sum := Add(1, 1)
	fmt.Println(sum)
}

func TestAdder(t *testing.T) {
	t.Run("Adding two values results in an expected sum", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		if sum != expected {
			t.Errorf("Expected %d, but got %d instead", expected, sum)
		}
	})

	t.Run("Adding negative value to negative value, results in right value", func(t *testing.T) {
		sum := Add(-2, -5)
		expected := -7

		if sum != expected {
			t.Errorf("Expected %d, but got %d instead", expected, sum)
		}
	})

	t.Run("Adding two zero values results in zero", func(t *testing.T) {
		sum := Add(0, 0)
		expected := 0

		if sum != expected {
			t.Errorf("Expected %d, but got %d instead", expected, sum)
		}
	})
}
