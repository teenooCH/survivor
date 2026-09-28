package tile

import (
	"math"

	"survivor/internal/domain/graph"
	"survivor/internal/domain/node2D"
	"survivor/internal/domain/transform"
)

// TileMap renders an effectively infinite grid of tiles around a Viewport
// (typically the camera). Chunks intersecting the viewport (plus a small
// margin) are generated on demand and chunks that fall out of view are
// discarded again, so the map never needs to be fully held in memory.
type TileMap struct {
	node2D.Node2D

	tileSet   *TileSet
	generator ChunkGenerator
	viewport  Viewport

	tileWidth, tileHeight   int
	chunkWidth, chunkHeight int
	loadMargin              int // extra chunks kept loaded around the viewport

	layer  int
	loaded map[ChunkCoord]*Chunk
}

// NewMap creates a TileMap that keeps chunks around viewport loaded, using
// generator to produce chunk content and tileSet to resolve tile indices
// to textures. tileWidth/tileHeight is the size of one tile in pixels.
func NewMap(
	name string, tileSet *TileSet, generator ChunkGenerator, viewport Viewport,
	tileWidth, tileHeight, loadMargin, layer int,
) *TileMap {
	chunkWidth, chunkHeight := generator.ChunkSize()

	return &TileMap{
		Node2D:      *node2D.New(name),
		tileSet:     tileSet,
		generator:   generator,
		viewport:    viewport,
		tileWidth:   tileWidth,
		tileHeight:  tileHeight,
		chunkWidth:  chunkWidth,
		chunkHeight: chunkHeight,
		loadMargin:  loadMargin,
		layer:       layer,
		loaded:      make(map[ChunkCoord]*Chunk),
	}
}

func (m *TileMap) GetLayer() int { return m.layer }

// LoadedChunkCount returns the number of chunks currently kept in memory.
func (m *TileMap) LoadedChunkCount() int { return len(m.loaded) }

func (m *TileMap) chunkPixelSize() (width, height float64) {
	return float64(m.chunkWidth * m.tileWidth), float64(m.chunkHeight * m.tileHeight)
}

// visibleChunkRange returns the inclusive range of chunk coordinates that
// intersect the viewport, expanded by loadMargin chunks on every side.
func (m *TileMap) visibleChunkRange() (min, max ChunkCoord) {
	pos := m.viewport.GetPosition()
	chunkW, chunkH := m.chunkPixelSize()

	min = ChunkCoord{
		X: int(math.Floor(pos.X()/chunkW)) - m.loadMargin,
		Y: int(math.Floor(pos.Y()/chunkH)) - m.loadMargin,
	}
	max = ChunkCoord{
		X: int(math.Floor((pos.X()+float64(m.viewport.GetWidth()))/chunkW)) + m.loadMargin,
		Y: int(math.Floor((pos.Y()+float64(m.viewport.GetHeight()))/chunkH)) + m.loadMargin,
	}

	return min, max
}

// UpdateTileMap loads chunks that entered the viewport and unloads chunks that
// left it.
func (m *TileMap) UpdateTileMap() {
	min, max := m.visibleChunkRange()

	for y := min.Y; y <= max.Y; y++ {
		for x := min.X; x <= max.X; x++ {
			coord := ChunkCoord{X: x, Y: y}
			if _, ok := m.loaded[coord]; !ok {
				m.loaded[coord] = m.generator.Generate(coord)
			}
		}
	}

	for coord := range m.loaded {
		if coord.X < min.X || coord.X > max.X || coord.Y < min.Y || coord.Y > max.Y {
			delete(m.loaded, coord)
		}
	}
}

// Draw renders every tile of every currently loaded chunk. options.Transform
// provides the world-to-screen offset for the tilemap (in practice just the
// camera translation, since the tilemap itself stays at the world origin
// with no scale or rotation).
func (m *TileMap) Draw(target graph.Image, options graph.DrawOpt) {
	offset := options.Transform.Position()

	for _, chunk := range m.loaded {
		originX := float64(chunk.Coord.X*m.chunkWidth*m.tileWidth) + offset.X()
		originY := float64(chunk.Coord.Y*m.chunkHeight*m.tileHeight) + offset.Y()

		for ty := range chunk.Height {
			for tx := range chunk.Width {
				texture, ok := m.tileSet.Texture(chunk.TileAt(tx, ty))
				if !ok {
					continue
				}

				tr := transform.NewZero()
				tr.SetPosition(originX+float64(tx*m.tileWidth), originY+float64(ty*m.tileHeight))

				target.DrawImage(texture, graph.DrawOpt{Transform: tr})
			}
		}
	}
}

var _ graph.Drawable = (*TileMap)(nil)
