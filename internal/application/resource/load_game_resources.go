package resource

import (
	"fmt"
	"slices"
)

// Use case LoadGameResources loads all the game resources
// specified in the manifest into the resource manager.

type TextureSpec struct {
	Key  string
	Path string
}
type TileSetSpec struct {
	Key        string
	Path       string
	TileWidth  int
	TileHeight int
	Spacing    int
	Count      int
}
type PatternSpec struct {
	Key  string
	Path string
}
type Manifest struct {
	Textures []TextureSpec
	TileSets []TileSetSpec
	Patterns []PatternSpec
}

func LoadGameResources(rm *Manager, manifest Manifest) error {
	for texture := range slices.Values(manifest.Textures) {
		if err := rm.LoadTexture(texture.Key, texture.Path); err != nil {
			return fmt.Errorf("failed to load %s texture: %w", texture.Key, err)
		}
	}

	for tileset := range slices.Values(manifest.TileSets) {
		if err := rm.LoadTileset(
			tileset.Key, tileset.Path,
			tileset.TileWidth, tileset.TileHeight,
			tileset.Spacing, tileset.Count,
		); err != nil {
			return fmt.Errorf("failed to load tileset %s: %w", tileset.Key, err)
		}
	}

	for pattern := range slices.Values(manifest.Patterns) {
		if err := rm.LoadPattern(pattern.Key, pattern.Path); err != nil {
			return fmt.Errorf("failed to load pattern %s: %w", pattern.Key, err)
		}
	}

	return nil
}
