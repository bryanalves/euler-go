package euler

import "testing"

func TestProblem13(t *testing.T) {
	out := Problem13()
	expected := "5537376230"

	if out != expected {
		t.Errorf("Problem13() = %d, should be %d", out, expected)
	}
}
