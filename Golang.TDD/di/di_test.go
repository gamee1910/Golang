package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T){
	buffer := bytes.Buffer{}
	Greet(&buffer, "Game")

	got := buffer.String()
	want := "Hello, Game"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestLogError(t *testing.T){
	buffer := bytes.Buffer{}
	LogError(&buffer, "File not found")

	got := buffer.String()
	want := "ERROR: File not found"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}