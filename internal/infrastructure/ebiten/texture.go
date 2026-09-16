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
