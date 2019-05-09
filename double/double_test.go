package double

import "testing"

func TestNumber(t *testing.T) {
	expected := 4
	got := Number(2)

	if expected != got {
		t.Errorf("Ecpected %d got %d\n", expected, got)
	}
}
