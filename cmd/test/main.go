// Command test is a manual integration test that wires the player and
// engine packages together and renders a static player sprite to a
// 640x480 window using ebiten as the graphics backend.
package main

import (
	"log"

	"survivor/internal/application/engine"
	"survivor/internal/application/game"
	"survivor/internal/application/resource"
	"survivor/internal/application/settings"
	"survivor/internal/domain/camera"
	"survivor/internal/domain/world"
	"survivor/internal/infrastructure/assets"
	"survivor/internal/infrastructure/ebiten"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

func main() {
	textures := resource.NewManager(ebiten.NewTextureProvider(assets.FS))
	if err := textures.LoadTexture(assets.Player, settings.PlayerTexture); err != nil {
		log.Fatalf("failed to load player texture: %v", err)
	}

	camera := camera.New(screenWidth, screenHeight)
	world := world.New(camera)
	eng := engine.New(world, textures, ebiten.NewInputProvider())

	game.RegisterDefaultBindings(eng)

	p := game.CreatePlayer(settings.PlayerName, eng, 0)
	p.SetPosition(screenWidth/2, screenHeight/2)
	eng.World().AddNode(p, 0)

	ebiten.RunGame(eng)
}
