package tile

// ChunkCoord identifies a chunk's position on the infinite chunk grid.
type ChunkCoord struct {
	X, Y int
}

// Chunk holds the tile indices for one chunk of the tilemap.
type Chunk struct {
	Coord         ChunkCoord
	Width, Height int
	Tiles         []int
}

// TileAt returns the tile index at the given local (in-chunk) coordinate.
func (c *Chunk) TileAt(x, y int) int {
	return c.Tiles[y*c.Width+x]
}
