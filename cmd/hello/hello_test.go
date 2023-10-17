package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Hello to humans", func(t *testing.T) {
		got := SayHello("Suraj")
		want := "Hello, Suraj"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Hello to everyone when no human is present", func(t *testing.T) {
		got := SayHello("")
		want := "Hello, world"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
