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

// Update the game state, e.g., player, enemy, collisions, etc.
func (g *Game) Update() {
	if g.gameOver {
		return
	}

	g.player.Update()
	g.enemy.Update()
	g.engine.CollisionManager().ProcessCollisions()

	g.engine.Update()
}

func (g *Game) Draw(screen graph.Image) {
	g.engine.Draw(screen)
	g.hud.Draw(screen)
}

func (g *Game) SetGameOver() {
	g.gameOver = true
	g.gameOverWidget.SetVisible(true)
}
