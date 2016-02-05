package hex

// Vector is hexagonal vector with (U,V) axial coordinates
type Vector struct {
	U, V int
}

// Add computes the sum vector
func Add(vs ...Vector) (ret Vector) {
	for _, v := range vs {
		ret.U += v.U
		ret.V += v.V
	}
	return
}
