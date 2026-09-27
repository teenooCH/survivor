package ebiten

import (
	"log"

	"survivor/internal/application/game"
	"survivor/internal/domain/graph"

	"github.com/hajimehoshi/ebiten/v2"
)

// ebGame implements ebiten.Game and drives the survivor engine.
type ebGame struct {
	game   *game.Game
	width  int
	height int
}

func (g *ebGame) Update() error {
	g.game.Update()
	return nil
}

func (g *ebGame) Draw(screen *ebiten.Image) {
	g.game.Draw(newEbitenImage(screen))
}

func (g *ebGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.width, g.height
}

func GetNewImage(width, height int) graph.Image {
	return newEbitenImage(ebiten.NewImage(width, height))
}

func RunGame(g *game.Game, title string, width, height int) {
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle(title)

	if err := ebiten.RunGame(&ebGame{game: g, width: width, height: height}); err != nil {
		log.Fatal(err)
	}
}
