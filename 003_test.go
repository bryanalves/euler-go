package euler

import "testing"

func TestProblem3(t *testing.T) {
	out := Problem3(600851475143)
	expected := 6857

	if out != expected {
		t.Errorf("Problem3(600851475143) = %d, should be %d", out, expected)
	}
}
