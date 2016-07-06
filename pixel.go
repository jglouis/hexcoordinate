package hex

import "fmt"

// PixelCoordinate represents pixel Coordinates (X,Y)
type PixelCoordinate PixelVector

// PixelVector represents a vector (X,Y)
type PixelVector struct {
	X, Y float64
}

func (c PixelCoordinate) String() string {
	return fmt.Sprintf("Pixel Coordinate(%.2f,%.2f)", c.X, c.Y)
}
