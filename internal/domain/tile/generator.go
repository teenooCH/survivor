package tile

// ChunkGenerator produces the tile grid for any chunk coordinate on demand,
// which is what allows a TileMap to cover an effectively infinite world
// without ever holding it all in memory at once.
type ChunkGenerator interface {
	// ChunkSize returns the width and height of a chunk in tiles.
	ChunkSize() (width, height int)
	// Generate returns the chunk at the given chunk coordinate.
	Generate(coord ChunkCoord) *Chunk
}

// PatternGenerator is a ChunkGenerator that repeats a fixed Pattern for
// every chunk, producing an infinitely tiled version of that pattern.
type PatternGenerator struct {
	pattern Pattern
}

func NewPatternGenerator(pattern Pattern) *PatternGenerator {
	return &PatternGenerator{pattern: pattern}
}

func (g *PatternGenerator) ChunkSize() (width, height int) {
	return g.pattern.Width(), g.pattern.Height()
}

func (g *PatternGenerator) Generate(coord ChunkCoord) *Chunk {
	tiles := make([]int, len(g.pattern.tiles))
	copy(tiles, g.pattern.tiles)

	return &Chunk{
		Coord:  coord,
		Width:  g.pattern.Width(),
		Height: g.pattern.Height(),
		Tiles:  tiles,
	}
}
