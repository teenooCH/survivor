package game

import (
	"fmt"

	"survivor/internal/application/engine"
	"survivor/internal/application/settings"
	"survivor/internal/domain/collision"
	"survivor/internal/domain/enemy"
	"survivor/internal/domain/sprite"
)

// Use case for creating a new enemy.

func CreateEnemy(name string, engine *engine.Engine, layer int) *enemy.Enemy {
	e := enemy.New(name)

	tex, _ := engine.ResourceManager().GetTexture(settings.EnemyTexture)
	spriteName := fmt.Sprintf("%s_sprite", name)
	sprite := sprite.New(spriteName, tex, layer)
	sprite.SetScale(settings.EnemyScaleFactor, settings.EnemyScaleFactor)
	e.AddChild(sprite)

	colliderName := fmt.Sprintf("%s_collider", name)
	shape := collision.NewCircle(settings.EnemyColliderRadius)
	mask := collision.NewMask(collision.LayerEnemy, collision.LayerPlayer)
	c := collision.NewCollider(colliderName, mask, shape)
	e.SetCollider(c)
	e.AddChild(c)
	engine.CollisionManager().AddCollider(c)

	engine.World().AddNode(e, layer)

	return e
}
