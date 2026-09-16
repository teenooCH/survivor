package transform

import "survivor/internal/domain/vector"

// Transformanble ist the interface for any node that has a transform.
type Transformable interface {
	GetTransform() Transform
	SetTransform(t Transform)
	GetWorldTransform() Transform
}

// Transform holds position, pivot, rotation (radians), and scale.
// It is the data that Node2D uses for its local and world transforms.
//   - position: Where the node is placed in its parent’s coordinate system.
//     For a sprite at (100, 50), the sprite’s origin (or pivot)
//     will appear 100 pixels right and 50 down from the parent’s origin.
//     This is the translation that moves the node.
//   - pivot: The point around which scaling and rotation happen.
//     For a sprite, (0, 0) is typically the top-left corner;
//     if you set the pivot to the centre of the texture, rotating
//     scales around that centre. The transform pipeline moves the pivot
//     to the origin, applies scale and rotation, then moves it back.
//   - rotation: The angle in radians (0 = no rotation, π/2 = 90° clockwise
//     in a Y-down coordinate system). The node is rotated around its pivot.
//   - scale: Stretch or shrink along X and Y. (1, 1) means no scaling;
//     (2, 2) doubles the size; (0.5, 1) halves the width.
//     Scale is applied around the pivot.
type Transform struct {
	position vector.Vector
	pivot    vector.Vector
	rotation float64 // in radians
	scale    vector.Vector
}

func New(position, pivot vector.Vector, rotation float64) Transform {
	return Transform{
		position: position,
		pivot:    pivot,
		rotation: rotation,
		scale:    vector.New(1, 1),
	}
}

// NewZero creates a transform with position (0, 0), pivot (0, 0), rotation 0, and scale (1, 1).
func NewZero() Transform {
	return Transform{
		position: vector.New(0, 0),
		pivot:    vector.New(0, 0),
		rotation: 0,
		scale:    vector.New(1, 1),
	}
}

func (t *Transform) Position() vector.Vector { return t.position }
func (t *Transform) Pivot() vector.Vector    { return t.pivot }
func (t *Transform) Rotation() float64       { return t.rotation }
func (t *Transform) Scale() vector.Vector    { return t.scale }

func (t *Transform) SetPosition(x, y float64) { t.position = vector.New(x, y) }
func (t *Transform) SetPivot(x, y float64)    { t.pivot = vector.New(x, y) }
func (t *Transform) SetRotation(rot float64)  { t.rotation = rot }
func (t *Transform) SetScale(x, y float64)    { t.scale = vector.New(x, y) }

func (t *Transform) Translate(x, y float64) { t.position = t.position.Translate(x, y) }
func (t *Transform) Rotate(radians float64) { t.rotation += radians }
func (t *Transform) ScaleBy(x, y float64)   { t.scale = vector.New(t.scale.X()*x, t.scale.Y()*y) }

// Concat combines two transforms: the receiver t (typically the
// parent’s world transform) and the argument other (the child’s
// local transform). The result is t × other. See Chapter 2.
func (t *Transform) Concatenate(other Transform) {
	sx := other.position.X() * t.scale.X()
	sy := other.position.Y() * t.scale.Y()
	rotated := vector.New(sx, sy).Rotate(t.rotation)

	t.Translate(rotated.X(), rotated.Y())
	t.ScaleBy(other.scale.X(), other.scale.Y())
	t.Rotate(other.rotation)
	t.SetPivot(other.Pivot().X(), other.Pivot().Y())
}
