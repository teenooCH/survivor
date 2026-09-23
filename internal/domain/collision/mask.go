package collision

const (
	LayerPlayer = 1 << iota
	LayerEnemy
	LayerProjectile
	LayerPickup
)

type Mask struct {
	layer         uint
	collisionMask uint
}

func NewMask(layer, collisionMask uint) Mask {
	return Mask{
		layer:         layer,
		collisionMask: collisionMask,
	}
}

// CanCollideWith returns true if this mask can collide
// with the other mask.
func (m Mask) CanCollideWith(other Mask) bool {
	return m.collisionMask&other.layer != 0
}

func (m Mask) Layer() uint {
	return m.layer
}

func (m Mask) CollisionMask() uint {
	return m.collisionMask
}
