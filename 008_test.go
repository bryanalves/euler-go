package main

import "testing"

func TestProblem8(t *testing.T) {
	out := problem_8()
	expected := 23514624000

	if out != expected {
		t.Errorf("problem_8() = %d, should be %d", out, expected)
	}
}
