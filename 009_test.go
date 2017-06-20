package euler

import "testing"

func TestProblem9(t *testing.T) {
	out := Problem9()
	expected := 31875000

	if out != expected {
		t.Errorf("Problem9() = %d, should be %d", out, expected)
	}
}
