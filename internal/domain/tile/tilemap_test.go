package tile_test

import (
	"testing"

	"survivor/internal/domain/graph"
	"survivor/internal/domain/tile"
	"survivor/internal/domain/transform"
	"survivor/internal/domain/vector"
)

// fakeViewport is a minimal tile.Viewport used to drive TileMap in tests.
type fakeViewport struct {
	pos           vector.Vector
	width, height uint
}

func (v *fakeViewport) GetPosition() vector.Vector { return v.pos }
func (v *fakeViewport) GetWidth() uint             { return v.width }
func (v *fakeViewport) GetHeight() uint            { return v.height }

// fakeImage is a minimal graph.Image used as the draw target in tests.
type fakeImage struct{ draws int }

func (f *fakeImage) DrawImage(graph.Image, graph.DrawOpt) { f.draws++ }
func (f *fakeImage) Dimensions() (width, height float64)  { return 0, 0 }
func (f *fakeImage) SetDimensions(width, height float64)  {}

func newTileMap(t *testing.T, viewport tile.Viewport) *tile.TileMap {
	t.Helper()

	pattern, err := tile.ParsePattern([]byte("0 1\n1 0\n"))
	if err != nil {
		t.Fatalf("ParsePattern() error = %v", err)
	}

	tileSet := tile.NewTileSet([]graph.Image{&fakeImage{}, &fakeImage{}})
	generator := tile.NewPatternGenerator(pattern)

	// tile size 10x10, chunk size 2x2 -> chunk is 20x20 pixels, no margin.
	return tile.New("floor", tileSet, generator, viewport, 10, 10, 0, 0)
}

func TestTileMapUpdateLoadsChunksCoveringViewport(t *testing.T) {
	viewport := &fakeViewport{pos: vector.New(0, 0), width: 25, height: 25}
	tm := newTileMap(t, viewport)

	tm.Update()

	// viewport spans pixels [0,25) in both axes with 20px chunks -> chunks 0 and 1.
	if got := tm.LoadedChunkCount(); got != 4 {
		t.Fatalf("LoadedChunkCount() = %d, want 4", got)
	}
}

func TestTileMapUpdateUnloadsChunksOutOfView(t *testing.T) {
	viewport := &fakeViewport{pos: vector.New(0, 0), width: 25, height: 25}
	tm := newTileMap(t, viewport)

	tm.Update()

	if got := tm.LoadedChunkCount(); got != 4 {
		t.Fatalf("LoadedChunkCount() after first Update() = %d, want 4", got)
	}

	// Move far away: none of the originally loaded chunks should remain.
	viewport.pos = vector.New(2000, 2000)

	tm.Update()

	if got := tm.LoadedChunkCount(); got != 4 {
		t.Fatalf("LoadedChunkCount() after moving = %d, want 4", got)
	}
}

func TestTileMapDrawDrawsLoadedTiles(t *testing.T) {
	viewport := &fakeViewport{pos: vector.New(0, 0), width: 20, height: 20}
	tm := newTileMap(t, viewport)

	tm.Update() // loads the 2x2 chunks covering the 20x20 viewport (chunk size == viewport size)

	target := &fakeImage{}
	tm.Draw(target, graph.DrawOpt{Transform: transform.NewZero()})

	// 4 chunks of 2x2 tiles each -> 16 draw calls.
	if target.draws != 16 {
		t.Fatalf("Draw() issued %d DrawImage calls, want 16", target.draws)
	}
}

var _ graph.Drawable = (*tile.TileMap)(nil)
