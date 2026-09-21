package game

import (
	"fmt"

	"survivor/internal/application/engine"
	"survivor/internal/application/settings"
	"survivor/internal/domain/graph"
	"survivor/internal/domain/tile"
)

// Use case for creating the infinite floor tilemap.

// CreateTileMap builds the tilemap from the floor pattern and tile textures
// already loaded into the engine's resource Manager (see
// resource.Manager.LoadPattern and LoadTileset), and adds it to the world
// on the given layer. viewport is typically the engine's camera; the
// tilemap streams chunks in and out of memory as the viewport moves.
func CreateTileMap(name string, engine *engine.Engine, viewport tile.Viewport, layer int) (*tile.TileMap, error) {
	pattern, ok := engine.ResourceManager().GetPattern(settings.FloorMapPattern)
	if !ok {
		return nil, fmt.Errorf("tilemap: pattern %q not loaded", settings.FloorMapPattern)
	}

	textures := make([]graph.Image, settings.FloorTileCount)

	for i := range textures {
		texName := fmt.Sprintf("%s_%d", settings.FloorTileset, i)

		tex, ok := engine.ResourceManager().GetTexture(texName)
		if !ok {
			return nil, fmt.Errorf("tilemap: texture %q not loaded", texName)
		}

		textures[i] = tex
	}

	tileSet := tile.NewTileSet(textures)
	generator := tile.NewPatternGenerator(pattern)

	tm := tile.New(
		name, tileSet, generator, viewport,
		settings.TileWidth, settings.TileHeight, settings.TileMapLoadMargin, layer,
	)

	engine.World().AddNode(tm, layer)

	return tm, nil
}
