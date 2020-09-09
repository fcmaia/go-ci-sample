package main

import "testing"

func TestSum1(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}
func TestSum2(t *testing.T) {
	total := Sum(-1, 5)
	if total != 4 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 4)
	}
}
