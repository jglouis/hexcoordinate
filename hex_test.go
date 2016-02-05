package hex

import "testing"

func TestAddition(t *testing.T) {
	v1 := Vector{0, 1}
	v2 := Vector{2, 3}
	expected := Vector{2, 4}

	if Add(v1, v2) != expected {
		t.Error()
	}
}
