package euler

import "testing"

func TestProblem2(t *testing.T) {
	out := Problem2()
	expected := 4613732

	if out != expected {
		t.Errorf("Problem2() = %d, should be %d", out, expected)
	}
}
