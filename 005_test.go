
package main

import "testing"

func TestProblem5(t *testing.T) {
    out := problem_5()
    expected := 232792560

    if out != expected {
      t.Errorf("problem_5() = %d, should be %d", out, expected)
  }
}
