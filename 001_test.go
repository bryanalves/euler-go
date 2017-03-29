package main

import "testing"

func TestProblem1(t *testing.T) {
    out := problem_1(1000)
    expected := 233168

    if out != expected {
      t.Errorf("problem_1(1000) = %d, should be %d", out, expected)
  }
}
