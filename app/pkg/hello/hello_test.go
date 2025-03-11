package hello

import "testing"

func TestHelloWorld(t *testing.T) {
	want := "Hello, World!"
	got := Hello()
	
	if got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}