package main

import "testing"

func TestProblem7(t *testing.T) {
    out := problem_7()
    expected := 104743

    if out != expected {
        t.Errorf("problem_7() = %d, should be %d", out, expected)
    }

}
