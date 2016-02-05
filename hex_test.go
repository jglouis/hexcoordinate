package hex

import "testing"

func TestAddition(t *testing.T) {
	v1 := Vector{0, 1}
	v2 := Vector{2, 3}
	expected := Vector{2, 4}
	got := Add(v1, v2)

	if got != expected {
		t.Errorf("Add %s and %s, got %s expected %s", v1, v2, got, expected)
	}
}
