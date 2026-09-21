package ports

import "survivor/internal/domain/graph"

type TextureProvider interface {
	LoadTexture(name string) (graph.Image, error)

	// LoadTileset loads a spritesheet and slices it into count tiles of
	// tileWidth x tileHeight pixels stacked vertically, separated by
	// spacing pixels, returning one texture per tile index in order.
	LoadTileset(name string, tileWidth, tileHeight, spacing, count int) ([]graph.Image, error)
}
