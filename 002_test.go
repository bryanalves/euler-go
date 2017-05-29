package main

import "testing"

func TestProblem2(t *testing.T) {
	out := problem_2()
	expected := 4613732

	if out != expected {
		t.Errorf("problem_2() = %d, should be %d", out, expected)
	}
}
