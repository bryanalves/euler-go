package euler

import "testing"

func TestProblem12(t *testing.T) {
	out := Problem12()
	expected := 76576500

	if out != expected {
		t.Errorf("Problem12() = %d, should be %d", out, expected)
	}
}
