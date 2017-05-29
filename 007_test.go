package euler

import "testing"

func TestProblem7(t *testing.T) {
	out := Problem7()
	expected := 104743

	if out != expected {
		t.Errorf("Problem7() = %d, should be %d", out, expected)
	}

}
