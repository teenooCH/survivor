package ebiten

import (
	"log"

	"survivor/internal/application/game"
	"survivor/internal/domain/graph"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

// ebGame implements ebiten.Game and drives the survivor engine.
type ebGame struct {
	game *game.Game
}

func (g *ebGame) Update() error {
	g.game.Update()
	return nil
}

func (g *ebGame) Draw(screen *ebiten.Image) {
	g.game.Draw(newEbitenImage(screen))
}

func (g *ebGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func NewEbGame(game *game.Game) *ebGame {
	return &ebGame{
		game: game,
	}
}

func GetNewImage(width, height int) graph.Image {
	return newEbitenImage(ebiten.NewImage(width, height))
}

func RunGame(g *game.Game) {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Survivor - Player Draw Test")

	if err := ebiten.RunGame(&ebGame{game: g}); err != nil {
		log.Fatal(err)
	}
}
