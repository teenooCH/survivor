package tile_test

import (
	"testing"

	"survivor/internal/domain/tile"
)

func TestPatternGenerator(t *testing.T) {
	pattern, err := tile.ParsePattern([]byte("0 1\n1 0\n"))
	if err != nil {
		t.Fatalf("ParsePattern() error = %v", err)
	}

	g := tile.NewPatternGenerator(pattern)

	width, height := g.ChunkSize()
	if width != 2 || height != 2 {
		t.Fatalf("ChunkSize() = %dx%d, want 2x2", width, height)
	}

	// Generate returns the same pattern content for any chunk coordinate,
	// which is what makes the tilemap tile infinitely in every direction.
	for _, coord := range []tile.ChunkCoord{{X: 0, Y: 0}, {X: -3, Y: 5}, {X: 42, Y: -7}} {
		c := g.Generate(coord)

		if c.Coord != coord {
			t.Errorf("Generate(%v).Coord = %v, want %v", coord, c.Coord, coord)
		}

		if c.Width != 2 || c.Height != 2 {
			t.Errorf("Generate(%v) size = %dx%d, want 2x2", coord, c.Width, c.Height)
		}

		if c.TileAt(0, 0) != 0 || c.TileAt(1, 0) != 1 || c.TileAt(0, 1) != 1 || c.TileAt(1, 1) != 0 {
			t.Errorf("Generate(%v) tiles = %v, want pattern repeated", coord, c.Tiles)
		}
	}
}

func TestPatternGeneratorReturnsIndependentChunks(t *testing.T) {
	pattern, err := tile.ParsePattern([]byte("0 1\n"))
	if err != nil {
		t.Fatalf("ParsePattern() error = %v", err)
	}

	g := tile.NewPatternGenerator(pattern)

	a := g.Generate(tile.ChunkCoord{X: 0, Y: 0})
	b := g.Generate(tile.ChunkCoord{X: 0, Y: 0})

	a.Tiles[0] = 99

	if b.TileAt(0, 0) != 0 {
		t.Fatalf("mutating one generated chunk affected another: got %d, want 0", b.TileAt(0, 0))
	}
}
