package main

import "testing"

func TestProblem3(t *testing.T) {
	out := problem_3(600851475143)
	expected := 6857

	if out != expected {
		t.Errorf("problem_3(600851475143) = %d, should be %d", out, expected)
	}
}
