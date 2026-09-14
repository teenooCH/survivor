package game

// Use case for creating a new player in the game.
import (
	"fmt"

	"survivor/internal/entities/player"
	"survivor/internal/entities/sprite"
	"survivor/internal/game/engine"
	"survivor/internal/game/settings"
)

func CreatePlayer(name string, engine *engine.Engine, layer int) *player.Player {
	p := player.New(settings.PlayerName)

	tex, _ := engine.ResourceManager().GetTexture(settings.PlayerTexture)
	spriteName := fmt.Sprintf("%s_sprite", name)
	sprite := sprite.New(spriteName, tex, layer)
	p.AddChild(sprite)

	p.SetHP(settings.PlayerInitialHP)
	p.SetXP(settings.PlayerInitialXP)
	p.SetLevel(settings.PlayerInitialLevel)

	engine.World().AddNode(layer, p)

	return p
}
