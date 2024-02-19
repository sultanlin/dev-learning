package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Sultan", "")
		want := "Hello, Sultan"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)

	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Seniorita", "French")
		want := "Bonjour, Seniorita"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// TODO: Understand testing helpers
	t.Helper()
	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
	if got != want {
	}
}
