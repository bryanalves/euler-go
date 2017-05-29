package euler

import "testing"

func TestProblem5(t *testing.T) {
	out := Problem5()
	expected := 232792560

	if out != expected {
		t.Errorf("Problem5() = %d, should be %d", out, expected)
	}
}
