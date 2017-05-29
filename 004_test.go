package euler

import "testing"

func TestProblem4(t *testing.T) {
	out := Problem4()
	expected := 906609

	if out != expected {
		t.Errorf("Problem4() = %d, should be %d", out, expected)
	}
}
