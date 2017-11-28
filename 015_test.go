package euler

import "testing"

func TestProblem15(t *testing.T) {
	out := Problem15()
	expected := int64(137846528820)

	if out != expected {
		t.Errorf("Problem15() = %d, should be %d", out, expected)
	}
}
