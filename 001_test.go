package euler

import "testing"

func TestProblem1(t *testing.T) {
	out := Problem1()
	expected := 233168

	if out != expected {
		t.Errorf("Problem1() = %d, should be %d", out, expected)
	}
}
