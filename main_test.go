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

func TestGreetingPrefix(t *testing.T) {
	// testing we can work with different prefixes
	t.Run("Given a language, generates a prefix", func(t *testing.T) {
		got := greetingPrefix("Spanish")
		want := "Hola"
		assertInvariant(t, got, want)
	})

	t.Run("Given an invalid lanugage, returns a empty  value", func(t *testing.T) {
		got := greetingPrefix("Yoruba")
		want := "Hello"

		assertInvariant(t, got, want)
	})
}

func TestMain(t *testing.T) {
	// name is supplied, greet people name
	t.Run("Greets people by their name", func(t *testing.T) {
		got := Hello("Santa", "")
		want := "Hello Santa"

		assertInvariant(t, got, want)
	})

	t.Run("ensures greeting fallsback to 'Hello, World' when name is not provided",
		func(t *testing.T) {
			// empty string
			got := Hello("", "")
			want := "Hello World!"

			assertInvariant(t, got, want)
		})

	t.Run("cannot greet user in 'alternate language'", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")

		want := "Hola Elodie"
		assertInvariant(t, got, want)
	})

	t.Run("can greet user in second language", func(t *testing.T) {
		got := Hello("Rotweller", "French")
		want := "Bonjour Rotweller"

		assertInvariant(t, got, want)
	})
}
