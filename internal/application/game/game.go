// Package game assembles and coordinates the application's gameplay.
// CreateGame builds a game from loaded resources and an input provider,
// wiring together the engine, world, player, enemy, tilemap, and game-over UI.
//
// Game.Update advances the world, camera, tilemap streaming, and collision
// processing in order. Game.Draw renders the world and HUD without updating
// game state. Platform-specific input and rendering implementations are
// supplied by the caller.
package game

import (
	"survivor/internal/application/engine"
	"survivor/internal/domain/enemy"
	"survivor/internal/domain/graph"
	"survivor/internal/domain/player"
	"survivor/internal/domain/ui"
)

type Game struct {
	engine *engine.Engine
	player *player.Player
	enemy  *enemy.Enemy
	hud    *ui.HUD

	gameOver       bool
	gameOverWidget *ui.Text
}

func NewGame(engine *engine.Engine,
	player *player.Player, enemy *enemy.Enemy,
	gameOverWidget *ui.Text,
) *Game {
	return &Game{
		engine:         engine,
		player:         player,
		enemy:          enemy,
		hud:            ui.NewHUD(),
		gameOverWidget: gameOverWidget,
	}
}

func (g *Game) HUD() *ui.HUD { return g.hud }

// Update the game state in a specific order.
func (g *Game) Update() {
	if g.gameOver {
		return
	}

	// World updates updateable nodes such as the player, enemy,
	// and other entities exactly once per frame.
	g.engine.UpdateWorld()

	// Camera follows the newly updated player.
	g.engine.UpdateCamera()

	// Tilemap streaming uses the current camera position.
	g.engine.UpdateStreaming()

	// Collision checks use post-movement positions.
	g.engine.ProcessCollisions()
}

func (g *Game) Draw(screen graph.Image) {
	g.engine.DrawWorld(screen)
	g.hud.Draw(screen)
}

func (g *Game) SetGameOver() {
	g.gameOver = true
	g.gameOverWidget.SetVisible(true)
}
