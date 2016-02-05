package hex

import "testing"

func TestAddition(t *testing.T) {
	v1 := Vector{0, 1}
	v2 := Vector{2, 3}
	want := Vector{2, 4}
	got := Add(v1, v2)

	if got != want {
		t.Errorf("Add %s and %s, got %s expected %s", v1, v2, got, want)
	}
}

func TestRotation(t *testing.T) {
	tests := map[Orientation]Orientation{
		N.Rotate(7):  NE,
		S.Rotate(-1): SE,
	}

	for got, want := range tests {
		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	}
}
