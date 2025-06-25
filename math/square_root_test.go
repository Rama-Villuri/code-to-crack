package math

import "testing"

func TestSquareRoot(t *testing.T) {
	got := SquareRoot(16)
	want := 4.0

	if got != want {
		t.Errorf("SquareRoot(16) = %v; want %v", got, want)
	}
}
