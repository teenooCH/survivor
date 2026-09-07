// package vector provides a simple 2D vector implementation with basic operations such as rotation.
// It can be used as position, velocity, or direction in 2D space.
// The Vector structure serves as:
//   - Point/Position: Representing a specific location in 2D space.
//   - Direction/offset: Indicating a direction or offset from a point.
//   - Delta/velocity: Representing the speed and direction of movement in 2D space.
package vector

import "math"

type Vector struct {
	x, y float64
}

func New(x, y float64) Vector { return Vector{x, y} }

func (v Vector) X() float64 { return v.x }
func (v Vector) Y() float64 { return v.y }

// Translate translates the vector by x,y and returns a new Vector instance.
func (v Vector) Translate(x, y float64) Vector { return Vector{v.x + x, v.y + y} }

// Rotate rotates the vector by the given angle in radians and returns a new Vector instance.
func (v Vector) Rotate(radians float64) Vector {
	s, c := math.Sin(radians), math.Cos(radians)

	return Vector{
		x: v.x*c - v.y*s,
		y: v.x*s + v.y*c,
	}
}

// IsEqual checks if two Vector instances are equal within
// a small epsilon to account for floating-point inaccuracies.
func (v Vector) IsEqual(b Vector) bool {
	const epsilon = 1e-9
	return math.Abs(v.X()-b.X()) < epsilon && math.Abs(v.Y()-b.Y()) < epsilon
}
