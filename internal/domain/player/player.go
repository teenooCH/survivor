package player

import (
	"math"

	"survivor/internal/domain/collision"
	"survivor/internal/domain/input"
	"survivor/internal/domain/node2D"
)

type Player struct {
	node2D.Node2D

	hp    float64
	xp    int
	level int
	speed float64

	collider *collision.Collider

	input *input.Manager
}

// New creates a Player driven by the given input Manager.
// speed is the movement distance in pixels applied per Update call.
func New(name string, inputManager *input.Manager, speed float64) *Player {
	p := &Player{
		Node2D: *node2D.New(name),
		speed:  speed,
		input:  inputManager,
	}

	return p
}

func (p *Player) HP() float64                   { return p.hp }
func (p *Player) XP() int                       { return p.xp }
func (p *Player) Level() int                    { return p.level }
func (p *Player) Collider() *collision.Collider { return p.collider }

func (p *Player) SetHP(hp float64)                         { p.hp = hp }
func (p *Player) SetXP(xp int)                             { p.xp = xp }
func (p *Player) SetLevel(level int)                       { p.level = level }
func (p *Player) SetCollider(collider *collision.Collider) { p.collider = collider }

// Update reads the bound movement Actions and moves the player accordingly.
// It implements world.Updatable so the scene graph drives it automatically.
func (p *Player) Update() {
	if p.input == nil {
		return
	}

	dx, dy := 0.0, 0.0

	if p.input.IsActionPressed(input.ActionMoveUp) {
		dy -= 1
	}

	if p.input.IsActionPressed(input.ActionMoveDown) {
		dy += 1
	}

	if p.input.IsActionPressed(input.ActionMoveLeft) {
		dx -= 1
	}

	if p.input.IsActionPressed(input.ActionMoveRight) {
		dx += 1
	}

	if dx == 0 && dy == 0 {
		return
	}

	length := math.Hypot(dx, dy)
	dx /= length
	dy /= length

	pos := p.GetPosition()
	p.SetPosition(pos.X()+dx*p.speed, pos.Y()+dy*p.speed)
}
