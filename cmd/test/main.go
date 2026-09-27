// Command test is a manual integration test that wires the player and
// engine packages together and renders a static player sprite to a
// 640x480 window using ebiten as the graphics backend.
package main

import (
	"log"

	"survivor/internal/application/game"
	"survivor/internal/application/resource"
	"survivor/internal/application/settings"
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

	g := game.CreateGame(textures, ebiten.NewInputProvider(), screenWidth, screenHeight)

	ebiten.RunGame(g)
}
