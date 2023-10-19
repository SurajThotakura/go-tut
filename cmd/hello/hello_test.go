package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Hello to humans", func(t *testing.T) {
		got := SayHello("Suraj", "English")
		want := "Hello, Suraj"

		formattedMessage(t, got, want)
	})
	t.Run("Hello in Telugu", func(t *testing.T) {
		got := SayHello("Suraj", "Telugu")
		want := "Namaskaram, Suraj"

		formattedMessage(t, got, want)
	})
	t.Run("Hello in Hindi", func(t *testing.T) {
		got := SayHello("Suraj", "Hindi")
		want := "Namaste, Suraj"

		formattedMessage(t, got, want)
	})
	t.Run("Hello to everyone when no human is present", func(t *testing.T) {
		got := SayHello("", "")
		want := "Hello, world"

		formattedMessage(t, got, want)
	})
}

func formattedMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
