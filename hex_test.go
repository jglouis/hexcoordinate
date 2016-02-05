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

	tests := map[Vector]Vector{
		Add(v1, v2):     Vector{2, 4},
		Add(v1, v2, v3): Vector{5, 8},
	}

	for got, want := range tests {
		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	}
}

func TestVectorAngle(t *testing.T) {
	v0 := Vector{0, 0}
	v1 := Vector{0, -1}
	v2 := Vector{0, 1}
	v3 := Vector{-1, 1}

	tests := map[float64]float64{
		Angle(v0, v0): 0,
		Angle(v1, v2): math.Pi,
		Angle(v2, v3): math.Pi / 6, // should not pass!
	}

	for got, want := range tests {
		if !floatEquals(got, want) {
			t.Errorf("got %f but want %f", got, want)
		}
	}
}

func TestVectorString(t *testing.T) {
	tests := map[string]string{
		Vector{0, 0}.String():  "Vector(0,0)",
		Vector{-1, 0}.String(): "Vector(-1,0)",
		Vector{4, 5}.String():  "Vector(4,5)",
	}

	for got, want := range tests {
		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	}
}

func TestCoordinateDistance(t *testing.T) {
	tests := map[int]int{
		Distance(Coordinate{0, 0}, Coordinate{0, 1}):  1,
		Distance(Coordinate{0, 0}, Coordinate{0, 0}):  0,
		Distance(Coordinate{3, 0}, Coordinate{-3, 2}): 6,
	}

	for got, want := range tests {
		if got != want {
			t.Errorf("got %d but want %d", got, want)
		}
	}
}

func TestOrientationRotate(t *testing.T) {
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

func TestOrientationVector(t *testing.T) {
	tests := map[Vector]Vector{
		N.Vector():  Vector{0, 1},
		NE.Vector(): Vector{-1, 1},
		SE.Vector(): Vector{-1, 0},
		S.Vector():  Vector{0, -1},
		SW.Vector(): Vector{1, -1},
		NW.Vector(): Vector{1, 0},
	}

	for got, want := range tests {
		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	}
}

func TestOrientationString(t *testing.T) {
	tests := map[string]string{
		N.String():               "N",
		NE.String():              "NE",
		SE.String():              "SE",
		S.String():               "S",
		SW.String():              "SW",
		NW.String():              "NW",
		Orientation(-1).String(): "Orientation(-1)",
	}

	for got, want := range tests {
		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	}
}
