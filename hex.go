// Package hex provides tools for handling axial hexagonal coordinates.
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

// Coordinate represents hexagonal Coordinates with (U,V) axial coordinates
type Coordinate Vector

func (c Coordinate) String() string {
	return fmt.Sprintf("Hexagonal Coordinate(%d,%d)", c.U, c.V)
}

// Vector represents a hexagonal vector with (U,V) axial coordinates
type Vector struct {
	U, V int
}

func (v Vector) String() string {
	return fmt.Sprintf("Hexagonal Vector(%d,%d)", v.U, v.V)
}

// Add returns a Vector that is the sum of all the given Vectors
func Add(vs ...Vector) (ret Vector) {
	for _, v := range vs {
		ret.U += v.U
		ret.V += v.V
	}
	return
}

// RotateBy60 takes a Vector and returns another Vector, rotated by n*60 degrees.
func (v Vector) RotateBy60(n int) Vector {
	// Transform to cubic coordinates
	x, z := v.U, v.V
	y := -x - z

	// Shifting coordinates
	nShift := n % 3
	if n < 0 {
		nShift = abs(n-1) % 3
	}
	switch nShift {
	case 1:
		x, y, z = z, x, y
	case 2:
		x, y, z = y, z, x
	}

	// Flip signs every 60 degrees
	if abs(n%2) == 1 {
		x, y, z = -x, -y, -z
	}

	return Vector{x, z}
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

// Distance returns the distance as a number of hexagon tiles separating two coordinates
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

// RotateBy60 returns a rotated Orientation
func (o Orientation) RotateBy60(n int) Orientation {
	return (o + Orientation(n)) % 6
}

// GridOrientation indicates how hewagons are oriented in the grid (flat or pointy top)
type GridOrientation int

// GridOrientation constants
const (
	PointyTop GridOrientation = iota
	FlatTop
)

// HexToPixel converts hexagonal coordinates into pixel coordinates
func (c Coordinate) HexToPixel(gridOrientation GridOrientation) PixelCoordinate {
	var x, y float64
	switch gridOrientation {
	case PointyTop:
		x = math.Sqrt(3) * (float64(c.U) + float64(c.V)/2)
		y = 1.5 * float64(c.V)
	case FlatTop:
		x = 1.5 * float64(c.U)
		y = math.Sqrt(3) * (float64(c.V) + float64(c.U)/2)
	}
	return PixelCoordinate{X: x, Y: y}
}

// PixelToHex converts pixel coordinates into hexagonal coordinates
func (p PixelCoordinate) PixelToHex(gridOrientation GridOrientation) Coordinate {
	var u, v float64
	switch gridOrientation {
	case PointyTop:
		u = p.X*math.Sqrt(3)/3 - p.Y/3
		v = p.Y * 2.0 / 3.0
	case FlatTop:
		u = p.X * 2.0 / 3.0
		v = p.Y*math.Sqrt(3)/3 - p.X/3
	}
	return Coordinate{U: round(u), V: round(v)}
}

func round(f float64) int {
	return int(math.Floor(f + .5))
}
