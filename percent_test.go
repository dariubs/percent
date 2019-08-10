package percent

import (
	"testing"
)

func TestPercent(t *testing.T) {
	pct, total := 25, 200
	part := Percent(pct, total)

	if part != 50.0 {
		t.Fatalf("%d is wrong number for %d percent", int(part), pct)
	}
}

func TestPercentFloat(t *testing.T) {
	pct, total := 25.0, 200.0
	part := PercentFloat(pct, total)

	if part != 50.0 {
		t.Fatalf("%d is wrong number for %f percent!", int(part), pct)
	}
}

func TestPercentOf(t *testing.T) {
	part, total := 300, 2400
	pct := PercentOf(part, total)

	if pct != 12.5 {
		t.Fatalf("%f is wrong percent!", pct)
	}
}

func TestPercentOfFloat(t *testing.T) {
	part, total := 300.0, 2400.0
	pct := PercentOfFloat(part, total)

	if pct != 12.5 {
		t.Fatalf("%f is wrong percent!", pct)
	}
}

func TestChange(t *testing.T) {
	before, after := 20, 60
	pct := Change(before, after)

	if int(pct) != 200.0 {
		t.Fatalf("%f is wrong percent!", pct)
	}
}

func TestChangeFloat(t *testing.T) {
	before, after := 20.0, 60.0
	pct := ChangeFloat(before, after)

	if int(pct) != 200.0 {
		t.Fatalf("%f is wrong percent!", pct)
	}
}
