// Command test is a manual integration test that wires the player and
// engine packages together and renders a static player sprite to a
// 640x480 window using ebiten as the graphics backend.
package main

import (
	"bytes"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"survivor/internal/entities/engine"
	"survivor/internal/entities/graph"
	"survivor/internal/entities/player"
	"survivor/internal/game/settings"
	"survivor/internal/infrastructure/assets"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

// ebitenImage adapts *ebiten.Image to the graph.Image interface.
type ebitenImage struct {
	img *ebiten.Image
}

func newEbitenImage(img *ebiten.Image) *ebitenImage { return &ebitenImage{img: img} }

func (e *ebitenImage) DrawImage(src graph.Image, options graph.DrawOpt) {
	source, ok := src.(*ebitenImage)
	if !ok {
		return
	}

	tr := options.Transform
	scale := tr.Scale()
	pos := tr.Position()

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale.X(), scale.Y())
	op.GeoM.Rotate(tr.Rotation())
	op.GeoM.Translate(pos.X(), pos.Y())

	e.img.DrawImage(source.img, op)
}

func (e *ebitenImage) Dimensions() (width, height float64) {
	b := e.img.Bounds()
	return float64(b.Dx()), float64(b.Dy())
}

func (e *ebitenImage) SetDimensions(width, height float64) {
	// ebiten images have a fixed size once created; nothing to do here.
}

// textureManager is a resource.Manager backed by ebiten images loaded
// from the embedded assets filesystem.
type textureManager struct {
	textures map[string]graph.Image
}

func newTextureManager() *textureManager {
	return &textureManager{textures: make(map[string]graph.Image)}
}

func (m *textureManager) LoadTexture(path, name string) error {
	data, err := assets.FS.ReadFile(path)
	if err != nil {
		return err
	}

	src, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return err
	}

	m.textures[name] = newEbitenImage(ebiten.NewImageFromImage(src))

	return nil
}

func (m *textureManager) GetTexture(name string) (graph.Image, bool) {
	tex, ok := m.textures[name]
	return tex, ok
}

// game implements ebiten.Game and drives the survivor engine.
type game struct {
	engine *engine.Engine
}

func (g *game) Update() error {
	g.engine.Update()
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	g.engine.Draw(newEbitenImage(screen))
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	textures := newTextureManager()
	if err := textures.LoadTexture(assets.Player, settings.PlayerTexture); err != nil {
		log.Fatalf("failed to load player texture: %v", err)
	}

	surface := newEbitenImage(ebiten.NewImage(screenWidth, screenHeight))
	eng := engine.New(surface, textures)

	p := player.NewPlayer(eng)
	p.SetPosition(screenWidth/2, screenHeight/2)
	eng.World().AddNode(0, p)

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Survivor - Player Draw Test")

	if err := ebiten.RunGame(&game{engine: eng}); err != nil {
		log.Fatal(err)
	}
}
