package hex

import "fmt"

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
	N:  Vector{0, 1},
	NE: Vector{-1, 1},
	SE: Vector{-1, 0},
	S:  Vector{0, -1},
	SW: Vector{1, -1},
	NW: Vector{1, 0},
}

// Vector gives the hex.Vector correponding to the Orientation
func (o Orientation) Vector() Vector {
	return vectorByOrientation[o]
}

// Rotate returns a rotated orientation
func (o Orientation) Rotate(by60degrees int) Orientation {
	return (o + Orientation(by60degrees)) % 6
}
