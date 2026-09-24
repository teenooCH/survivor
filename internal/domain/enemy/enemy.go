package enemy

import (
	"math"

	"survivor/internal/domain/collision"
	"survivor/internal/domain/node2D"
	"survivor/internal/domain/vector"
)

type PositionProvider interface {
	GetPosition() vector.Vector
}

type Enemy struct {
	node2D.Node2D

	hp    float64
	speed float64

	collider *collision.Collider

	targetNode PositionProvider
}

func New(name string) *Enemy {
	e := &Enemy{
		Node2D: *node2D.New(name),
	}

	return e
}

func (e *Enemy) HP() float64                   { return e.hp }
func (e *Enemy) Speed() float64                { return e.speed }
func (e *Enemy) Collider() *collision.Collider { return e.collider }
func (e *Enemy) Target() PositionProvider      { return e.targetNode }

func (e *Enemy) SetHP(hp float64)                  { e.hp = hp }
func (e *Enemy) SetSpeed(speed float64)            { e.speed = speed }
func (e *Enemy) SetCollider(c *collision.Collider) { e.collider = c }
func (e *Enemy) SetTarget(target PositionProvider) { e.targetNode = target }

// move towards a given target position.
func (e *Enemy) moveToward(target vector.Vector) {
	pos := e.GetPosition()
	dX, dY := target.X()-pos.X(), target.Y()-pos.Y()

	if dX == 0 && dY == 0 {
		return
	}

	length := math.Hypot(dX, dY)
	dX /= length
	dY /= length

	e.SetPosition(pos.X()+dX*e.speed, pos.Y()+dY*e.speed)
}

func (e *Enemy) Update() {
	if e.targetNode != nil {
		e.moveToward(e.targetNode.GetPosition())
	}
}
