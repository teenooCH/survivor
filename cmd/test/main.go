// Command test is a manual integration test that wires the player and
// engine packages together and renders a static player sprite to a
// 640x480 window using ebiten as the graphics backend.
package main

import (
	"log"

	"survivor/internal/application/game"
	"survivor/internal/application/resource"
	"survivor/internal/infrastructure/assets"
	"survivor/internal/infrastructure/ebiten"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

func main() {
	rm := resource.NewManager(ebiten.NewTextureProvider(assets.FS), assets.NewMapProvider(assets.FS))
	if err := resource.LoadGameResources(rm, assets.GameManifest()); err != nil {
		log.Fatalf("failed to load game resources: %v", err)
	}

	g := game.CreateGame(rm, ebiten.NewInputProvider(), screenWidth, screenHeight)

	ebiten.RunGame(g)
}
