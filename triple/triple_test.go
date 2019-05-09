package triple

import "testing"

func TestNumber(t *testing.T) {
	expected := 6
	got := Number(2)

	if expected != got {
		t.Errorf("Ecpected %d got %d\n", expected, got)
	}
}
