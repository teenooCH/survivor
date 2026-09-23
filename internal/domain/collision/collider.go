package collision

import "survivor/internal/domain/node2D"

type Collider struct {
	node2D.Node2D
	mask             Mask
	shape            *Circle
	collisionHandler func(other *Collider)
}

func NewCollider(name string, mask Mask, shape *Circle) *Collider {
	return &Collider{
		Node2D: *node2D.New(name),
		mask:   mask,
		shape:  shape,
	}
}

func (c *Collider) Mask() Mask {
	return c.mask
}

func (c *Collider) Shape() *Circle {
	return c.shape
}

func (c *Collider) CollisionHandler() func(other *Collider) {
	return c.collisionHandler
}

func (c *Collider) SetCollisionHandler(callback func(other *Collider)) {
	c.collisionHandler = callback
}

func (c *Collider) CanCollideWith(other *Collider) bool {
	return c.mask.CanCollideWith(other.mask)
}

// OverlapsWith returns true if this collider overlaps with other collider.
func (c *Collider) OverlapsWith(other *Collider) bool {
	dx := c.GetWorldPosition().X() - other.GetWorldPosition().X()
	dy := c.GetWorldPosition().Y() - other.GetWorldPosition().Y()
	distanceSquared := dx*dx + dy*dy
	radiusSum := c.shape.Radius() + other.shape.Radius()

	return distanceSquared < radiusSum*radiusSum
}
