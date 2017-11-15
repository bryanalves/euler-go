package euler

import "testing"

func TestProblem14(t *testing.T) {
	out := Problem14()
	expected := 837799

	if out != expected {
		t.Errorf("Problem14() = %d, should be %d", out, expected)
	}
}
