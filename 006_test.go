package euler

import "testing"

func TestProblem6(t *testing.T) {
	x := sum_squares(10)
	if x != 385 {
		t.Errorf("sum_squares fails: %d", x)
	}

	x = square_sums(10)
	if x != 3025 {
		t.Errorf("square_sums fails: %d", x)
	}

	out := Problem6()
	expected := 25164150

	if out != expected {
		t.Errorf("Problem6() = %d, should be %d", out, expected)
	}
}
