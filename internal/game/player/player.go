package player

import (
	"survivor/internal/entities/engine"
	"survivor/internal/entities/node2D"
	"survivor/internal/entities/sprite"
	"survivor/internal/game/settings"
)

type Player struct {
	node2D.Node2D
	sprite *sprite.Sprite
	engine *engine.Engine
}

func NewPlayer(engine *engine.Engine) *Player {
	p := &Player{
		Node2D: *node2D.New(settings.PlayerName),
		engine: engine,
	}
	p.SetPosition(0, 0)

	tex, _ := engine.ResourceManager().GetTexture(settings.PlayerTexture)
	sprite := sprite.New("", tex, 0)

	p.AddChild(sprite)
	p.sprite = sprite

	return p
}

// Sprite returns the player's sprite node so it can be registered
// with the world for rendering.
func (p *Player) Sprite() *sprite.Sprite { return p.sprite }
