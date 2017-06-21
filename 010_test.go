package euler

import "testing"

func TestProblem10(t *testing.T) {
	out := Problem10()
	expected := 142913828922

	if out != expected {
		t.Errorf("Problem10() = %d, should be %d", out, expected)
	}
}
