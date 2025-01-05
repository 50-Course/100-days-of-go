package main

import "testing"

// our little helper funciton that does
// somehting think of it like our little assert
func assertInvariant(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("error! expected %q, but got %q", want, got)
	}
}

func TestMain(t *testing.T) {
	// name is supplied, greet people name
	t.Run("Greets people by their name", func(t *testing.T) {
		got := greet("Santa", "")
		want := "Hello Santa"

		assertInvariant(t, got, want)
	})

	t.Run("ensures greeting fallsback to 'Hello, World' when name is not provided",
		func(t *testing.T) {
			// empty string
			got := greet("", "")
			want := "Hello World!"

			assertInvariant(t, got, want)
		})

	t.Run("cannot greet user in 'alternate language'", func(t *testing.T) {
		got := greet("Elodie", "Spanish")

		want := "Hola Elodie"
		assertInvariant(t, got, want)
	})

	t.Run("can greet user in second language", func(t *testing.T) {
		got := greet("Rotweller", "French")
		want := "Bonjour Rotweller"

		assertInvariant(t, got, want)
	})
}
