package player

import (
	"survivor/internal/domain/node2D"
)

type Player struct {
	node2D.Node2D
	hp    float64
	xp    int
	level int
}

func New(name string) *Player {
	p := &Player{
		Node2D: *node2D.New(name),
	}

	return p
}

func (p *Player) HP() float64 { return p.hp }
func (p *Player) XP() int     { return p.xp }
func (p *Player) Level() int  { return p.level }

func (p *Player) SetHP(hp float64)   { p.hp = hp }
func (p *Player) SetXP(xp int)       { p.xp = xp }
func (p *Player) SetLevel(level int) { p.level = level }
