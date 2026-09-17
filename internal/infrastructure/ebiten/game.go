package ebiten

import (
	"log"

	"survivor/internal/application/engine"
	"survivor/internal/domain/graph"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

// ebGame implements ebiten.Game and drives the survivor engine.
type ebGame struct {
	engine *engine.Engine
}

func (g *ebGame) Update() error {
	g.engine.Update()
	return nil
}

func (g *ebGame) Draw(screen *ebiten.Image) {
	g.engine.Draw(newEbitenImage(screen))
}

func (g *ebGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func NewEbGame(engine *engine.Engine) *ebGame {
	return &ebGame{
		engine: engine,
	}
}

func GetNewImage(width, height int) graph.Image {
	return newEbitenImage(ebiten.NewImage(width, height))
}

func RunGame(eng *engine.Engine) {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Survivor - Player Draw Test")

	if err := ebiten.RunGame(&ebGame{engine: eng}); err != nil {
		log.Fatal(err)
	}
}
