package tile

import "survivor/internal/domain/graph"

// TileSet maps tile indices (as used in a Pattern/Chunk) to their textures.
type TileSet struct {
	textures []graph.Image
}

func NewTileSet(textures []graph.Image) *TileSet {
	return &TileSet{textures: textures}
}

// Texture returns the texture for the given tile index.
// Returns false if the index is out of range.
func (t *TileSet) Texture(index int) (graph.Image, bool) {
	if index < 0 || index >= len(t.textures) {
		return nil, false
	}

	return t.textures[index], true
}
