package assets

import (
	"survivor/internal/application/resource"
	"survivor/internal/application/settings"
)

func GameManifest() resource.Manifest {
	return resource.Manifest{
		Textures: []resource.TextureSpec{
			{Key: settings.PlayerTexture, Path: Player},
			{Key: settings.EnemyTexture, Path: Enemy},
		},
		TileSets: []resource.TileSetSpec{
			{
				Key:        settings.FloorTileset,
				Path:       Spritesheet,
				TileWidth:  settings.TileWidth,
				TileHeight: settings.TileHeight,
				Spacing:    settings.TileSpacing,
				Count:      settings.FloorTileCount,
			},
		},
		Patterns: []resource.PatternSpec{
			{Key: settings.FloorMapPattern, Path: FloorMap},
		},
	}
}
