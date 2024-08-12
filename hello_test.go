package main 

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Logan")
	want := "Hello, Logan!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}