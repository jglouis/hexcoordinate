package hex

import (
	"fmt"
	"math"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// Coordinate is hexagonal Coordinate with (U,V) axial coordinates
type Coordinate Vector

// Vector is hexagonal vector with (U,V) axial coordinates
type Vector struct {
	U, V int
}

func (v Vector) String() string {
	return fmt.Sprintf("Vector(%d,%d)", v.U, v.V)
}

// Add computes the sum vector
func Add(vs ...Vector) (ret Vector) {
	for _, v := range vs {
		ret.U += v.U
		ret.V += v.V
	}
	return
}

// Angle returns the angle between two Vectors
func Angle(vector1, vector2 Vector) float64 {
	u1 := float64(vector1.U)
	v1 := float64(vector1.V)
	u2 := float64(vector2.U)
	v2 := float64(vector2.V)
	x1 := math.Sqrt(3) * (u1 + v1/2.0)
	y1 := 1.5 * v1
	x2 := math.Sqrt(3) * (u2 + v2/2.0)
	y2 := 1.5 * v2
	// atan2(vector2.y, vector2.x) - atan2(vector1.y, vector1.x);

	Dot := x1*x2 + y1*y2
	Det := x1*y2 - y1*x2

	return math.Atan2(Det, Dot)
}

// Distance return the distance as a number of hexagon tiles separating two coordinates
func Distance(a, b Coordinate) int {
	return (abs(a.U-b.U) + abs(a.U+a.V-b.U-b.V) + abs(a.V-b.V)) / 2
}

// Orientation represents the orientation of something in a hexagonal vector space
type Orientation int

//go:generate stringer -type=Orientation
// Orientation constants
const (
	N Orientation = iota
	NE
	SE
	S
	SW
	NW
)

var vectorByOrientation = map[Orientation]Vector{
	N:  Vector{0, -1},
	NE: Vector{1, -1},
	SE: Vector{1, 0},
	S:  Vector{0, 1},
	SW: Vector{-1, 1},
	NW: Vector{-1, 0},
}

// Vector gives the hex.Vector correponding to the Orientation
func (o Orientation) Vector() Vector {
	return vectorByOrientation[o]
}

// Rotate returns a rotated orientation
func (o Orientation) Rotate(by60degrees int) Orientation {
	return (o + Orientation(by60degrees)) % 6
}
