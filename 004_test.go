package main

import "testing"

func TestProblem4(t *testing.T) {
    out := problem_4()
    expected := 906609

    if out != expected {
      t.Errorf("problem_4() = %d, should be %d", out, expected)
  }
}
