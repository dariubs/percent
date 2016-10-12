package percent

import (
  "testing"
)

func TestCalculate(t *testing.T) {
  part, all := 20, 50
  pcent := Calculate(part, all)

  if pcent != 40 {
    t.Fatalf("Wrong percent!")
  }
}
