package euler

import "testing"

func TestProblem8(t *testing.T) {
	out := Problem8()
	expected := 23514624000

	if out != expected {
		t.Errorf("Problem8() = %d, should be %d", out, expected)
	}
}
