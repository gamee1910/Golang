package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Game")
	want := "Hello, Game"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
