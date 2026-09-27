package game

import (
	"fmt"

	"survivor/internal/application/engine"
	"survivor/internal/application/settings"
	"survivor/internal/domain/collision"
	"survivor/internal/domain/player"
	"survivor/internal/domain/sprite"
)

// Use case for creating a new player in the game.

func CreatePlayer(name string, engine *engine.Engine, layer int) *player.Player {
	p := player.New(name, engine.InputManager(), settings.PlayerSpeed)

	tex, _ := engine.ResourceManager().GetTexture(settings.PlayerTexture)
	spriteName := fmt.Sprintf("%s_sprite", name)
	sprite := sprite.New(spriteName, tex, layer)
	p.AddChild(sprite)

	colliderName := fmt.Sprintf("%s_collider", name)
	shape := collision.NewCircle(settings.PlayerColliderRadius)
	mask := collision.NewMask(collision.LayerPlayer, collision.LayerEnemy)
	c := collision.NewCollider(colliderName, mask, shape)
	p.SetCollider(c)
	p.AddChild(c)
	engine.CollisionManager().AddCollider(c)

	p.SetHP(settings.PlayerInitialHP)
	p.SetXP(settings.PlayerInitialXP)
	p.SetLevel(settings.PlayerInitialLevel)

	return p
}
