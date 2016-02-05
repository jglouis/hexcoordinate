package hex

import (
	"math"
	"testing"
)

const EPSILON float64 = 0.0001

func floatEquals(a, b float64) bool {
	if (a-b) < EPSILON && (b-a) < EPSILON {
		return true
	}
	return false
}

func TestVectorAddition(t *testing.T) {
	v1 := Vector{0, 1}
	v2 := Vector{2, 3}
	v3 := Vector{3, 4}

	tests := []struct {
		inputs   []Vector
		expected Vector
	}{
		{[]Vector{v1, v2}, Vector{2, 4}},
		{[]Vector{v1, v2, v3}, Vector{5, 8}},
	}

	for _, test := range tests {
		got := Add(test.inputs...)
		if got != test.expected {
			t.Errorf("Adding vectors %s got %s but want %s", test.inputs, got, test.expected)
		}
	}
}

func TestVectorRotateBy60(t *testing.T) {
	tests := []struct {
		input    int
		expected Vector
	}{
		{0, N.Vector()},
		{1, NE.Vector()},
		{-1, NW.Vector()},
		{7, NE.Vector()},
		{-7, NW.Vector()},
		{-13, NW.Vector()},
	}

	for _, test := range tests {
		got := N.Vector().RotateBy60(test.input)
		if got != test.expected {
			t.Errorf("N.Vector()RotateBy60(%d) returned %s but expected %s", test.input, got, test.expected)
		}
	}
}

func TestVectorAngle(t *testing.T) {
	v0 := Vector{0, 0}

	tests := []struct {
		input1, input2 Vector
		expected       float64
	}{
		{v0, v0, 0},
		{N.Vector(), NE.Vector(), math.Pi / 3},
		{N.Vector(), NW.Vector(), -math.Pi / 3},
		{N.Vector(), S.Vector(), math.Pi},
	}

	for _, test := range tests {
		got := Angle(test.input1, test.input2)
		if !floatEquals(got, test.expected) {
			t.Errorf("Angle(%s,%s) returned %f but expected %f", test.input1, test.input2, got, test.expected)
		}
	}
}

func TestVectorString(t *testing.T) {
	tests := map[Vector]string{
		Vector{0, 0}:  "Vector(0,0)",
		Vector{-1, 0}: "Vector(-1,0)",
		Vector{4, 5}:  "Vector(4,5)",
	}

	for input, expected := range tests {
		got := input.String()
		if got != expected {
			t.Errorf("%s.String() returned %s but expected %s", input, got, expected)
		}
	}
}

func TestCoordinateString(t *testing.T) {
	tests := map[Coordinate]string{
		Coordinate{0, 0}:  "Coordinate(0,0)",
		Coordinate{-1, 0}: "Coordinate(-1,0)",
		Coordinate{4, 5}:  "Coordinate(4,5)",
	}

	for input, expected := range tests {
		got := input.String()
		if got != expected {
			t.Errorf("%s.String() returned %s but expected %s", input, got, expected)
		}
	}
}

func TestCoordinateDistance(t *testing.T) {
	tests := []struct {
		input1, input2 Coordinate
		expected       int
	}{
		{Coordinate{0, 0}, Coordinate{0, 1}, 1},
		{Coordinate{0, 0}, Coordinate{0, 0}, 0},
		{Coordinate{3, 0}, Coordinate{-3, 2}, 6},
	}

	for _, test := range tests {
		got := Distance(test.input1, test.input2)
		if got != test.expected {
			t.Errorf("Distance(%s,%s) returned %d but expected %d", test.input1, test.input2, got, test.expected)
		}
	}
}

func TestOrientationRotate(t *testing.T) {
	tests := []struct {
		inputOrientation Orientation
		inputN           int
		expected         Orientation
	}{
		{N, 7, NE},
		{S, -1, SE},
	}

	for _, test := range tests {
		got := test.inputOrientation.RotateBy60(test.inputN)
		if got != test.expected {
			t.Errorf("%s.RotateBy60(%d) returned %s but expected %s", test.inputOrientation, test.inputN, got, test.expected)
		}
	}
}

func TestOrientationVector(t *testing.T) {
	tests := map[Orientation]Vector{
		N:  Vector{0, -1},
		NE: Vector{1, -1},
		SE: Vector{1, 0},
		S:  Vector{0, 1},
		SW: Vector{-1, 1},
		NW: Vector{-1, 0},
	}

	for input, expected := range tests {
		got := input.Vector()
		if got != expected {
			t.Errorf("%s.Vector() returned %s but expected %s", input, got, expected)
		}
	}
}

func TestOrientationString(t *testing.T) {
	tests := map[Orientation]string{
		N:               "N",
		NE:              "NE",
		SE:              "SE",
		S:               "S",
		SW:              "SW",
		NW:              "NW",
		Orientation(-1): "Orientation(-1)",
	}

	for input, expected := range tests {
		got := input.String()
		if got != expected {
			t.Errorf("%s.String() returned %s but expected %s", input, got, expected)
		}
	}
}
