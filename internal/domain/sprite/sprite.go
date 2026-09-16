package sprite

import (
	"survivor/internal/domain/graph"
	"survivor/internal/domain/node2D"
)

type Sprite struct {
	node2D.Node2D
	texture graph.Image
	layer   int
	visible bool
}

// New creates a new Sprite instance with the given name, texture, and layer.
// If the texture is not nil, it sets the pivot to the center of the texture.
func New(name string, texture graph.Image, layer int) *Sprite {
	s := &Sprite{
		Node2D:  *node2D.New(name),
		texture: texture,
		layer:   layer,
		visible: true,
	}

	s.centerPivot()

	return s
}

func (s *Sprite) GetTexture() graph.Image { return s.texture }
func (s *Sprite) SetTexture(texture graph.Image) {
	s.texture = texture
	s.centerPivot()
}

func (s *Sprite) GetLayer() int      { return s.layer }
func (s *Sprite) SetLayer(layer int) { s.layer = layer }

func (s *Sprite) IsVisible() bool         { return s.visible }
func (s *Sprite) SetVisible(visible bool) { s.visible = visible }

func (s *Sprite) Draw(target graph.Image, options graph.DrawOpt) {
	if s.texture == nil || !s.visible {
		return
	}

	target.DrawImage(s.texture, options)
}

// centerPivot sets the pivot of the sprite to the center of its texture.
func (s *Sprite) centerPivot() {
	if s.texture == nil {
		return
	}

	width, height := s.texture.Dimensions()
	s.SetPivot(width/2, height/2)
}
