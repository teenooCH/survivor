package assets

import (
	"survivor/internal/application/resource"
	"survivor/internal/application/settings"
)

func GameManifest() resource.Manifest {
	return resource.Manifest{
		Textures: []resource.TextureSpec{
			{Key: settings.PlayerTexture, Path: "sprites/player.png"},
			{Key: settings.EnemyTexture, Path: "sprites/enemy.png"},
		},
		TileSets: []resource.TileSetSpec{
			{
				Key:        settings.FloorTileset,
				Path:       "sprites/spritesheet.png",
				TileWidth:  settings.TileWidth,
				TileHeight: settings.TileHeight,
				Spacing:    settings.TileSpacing,
				Count:      settings.FloorTileCount,
			},
		},
		Patterns: []resource.PatternSpec{
			{Key: settings.FloorMapPattern, Path: "maps/floor.map"},
		},
	}
}
