package game

import (
	"survivor/internal/domain/collision"
	"survivor/internal/domain/player"
)

// Use case for wiring player callbacks.
// This function connects the player's actions to the appropriate game logic, such as movement, attacks, and interactions with other game entities.
func WirePlayerCallbacks(player *player.Player, game *Game) {
	player.Collider().SetCollisionHandler(func(other *collision.Collider) {
		game.SetGameOver()
	})
}
