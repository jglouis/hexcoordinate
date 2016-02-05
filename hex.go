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
