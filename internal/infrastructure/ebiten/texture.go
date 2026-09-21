package ebiten

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/fs"

	"survivor/internal/domain/graph"

	"github.com/hajimehoshi/ebiten/v2"
)

type TextureProvider struct {
	fsys fs.FS
}

func NewTextureProvider(fsys fs.FS) *TextureProvider {
	return &TextureProvider{
		fsys: fsys,
	}
}

func (p *TextureProvider) LoadTexture(name string) (graph.Image, error) {
	data, err := fs.ReadFile(p.fsys, name)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", name, err)
	}

	src, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("failed to decode texture %s: %w", name, err)
	}

	return newEbitenImage(ebiten.NewImageFromImage(src)), nil
}

// LoadTileset loads a spritesheet and slices it into count tiles of
// tileWidth x tileHeight pixels stacked vertically, separated by spacing
// pixels, returning one independent texture per tile index in order.
func (p *TextureProvider) LoadTileset(
	name string, tileWidth, tileHeight, spacing, count int,
) ([]graph.Image, error) {
	data, err := fs.ReadFile(p.fsys, name)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", name, err)
	}

	src, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("failed to decode tileset %s: %w", name, err)
	}

	sheet := ebiten.NewImageFromImage(src)

	textures := make([]graph.Image, count)
	for i := range count {
		y := i * (tileHeight + spacing)
		rect := image.Rect(0, y, tileWidth, y+tileHeight)

		tile := ebiten.NewImage(tileWidth, tileHeight)
		tile.DrawImage(sheet.SubImage(rect).(*ebiten.Image), nil)

		textures[i] = newEbitenImage(tile)
	}

	return textures, nil
}
