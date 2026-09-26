// Command test is a manual integration test that wires the player and
// engine packages together and renders a static player sprite to a
// 640x480 window using ebiten as the graphics backend.
package main

import (
	"image/color"
	"log"

	"survivor/internal/application/engine"
	"survivor/internal/application/game"
	"survivor/internal/application/resource"
	"survivor/internal/application/settings"
	"survivor/internal/domain/camera"
	"survivor/internal/domain/ui"
	"survivor/internal/domain/world"
	"survivor/internal/infrastructure/assets"
	"survivor/internal/infrastructure/ebiten"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

func main() {
	textures := resource.NewManager(ebiten.NewTextureProvider(assets.FS), assets.NewMapProvider(assets.FS))
	if err := textures.LoadTexture(assets.Player, settings.PlayerTexture); err != nil {
		log.Fatalf("failed to load player texture: %v", err)
	}

	if err := textures.LoadTexture(assets.Enemy, settings.EnemyTexture); err != nil {
		log.Fatalf("failed to load enemy texture: %v", err)
	}

	if err := textures.LoadTileset(
		assets.Spritesheet, settings.FloorTileset,
		settings.TileWidth, settings.TileHeight, settings.TileSpacing, settings.FloorTileCount,
	); err != nil {
		log.Fatalf("failed to load floor tileset: %v", err)
	}

	if err := textures.LoadPattern(assets.FloorMap, settings.FloorMapPattern); err != nil {
		log.Fatalf("failed to load floor map: %v", err)
	}

	camera := camera.New(screenWidth, screenHeight)
	world := world.New(camera)
	eng := engine.New(world, textures, ebiten.NewInputProvider())

	game.RegisterDefaultBindings(eng)

	if _, err := game.CreateTileMap("floor", eng, camera, 0); err != nil {
		log.Fatalf("failed to create tilemap: %v", err)
	}

	p := game.CreatePlayer(settings.PlayerName, eng, 1)
	p.SetPosition(screenWidth/2, screenHeight/2)
	eng.World().AddNode(p, 1)
	camera.SetFollow(p)

	e := game.CreateEnemy("enemy", eng, 1)
	e.SetPosition(screenWidth/2+150, screenHeight/2)
	e.SetSpeed(settings.EnemySpeed)
	eng.World().AddNode(e, 1)
	e.SetTarget(p)

	hudText := ui.NewText("HP: 100", 10, 10)
	hudText.SetColor(color.White)
	hudText.SetSize(16)
	hudText.SetBackground(ui.NewBackground(color.RGBA{A: 160}))
	hudText.SetBorder(ui.NewBorder(color.White, 1))
	eng.HUD().AddWidget(hudText)

	ebiten.RunGame(eng)
}
