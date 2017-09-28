package euler

import "testing"

func TestProblem11(t *testing.T) {
	out := Problem11()
	expected := 70600674

	if out != expected {
		t.Errorf("Problem11() = %d, should be %d", out, expected)
	}
}
