package ebiten

import (
	"survivor/internal/domain/graph"

	"github.com/hajimehoshi/ebiten/v2"
)

// ebitenImage adapts *ebiten.Image to the graph.Image interface.
type ebitenImage struct {
	img *ebiten.Image
}

func newEbitenImage(img *ebiten.Image) *ebitenImage { return &ebitenImage{img: img} }

func (e *ebitenImage) DrawImage(src graph.Image, options graph.DrawOpt) {
	source, ok := src.(*ebitenImage)
	if !ok {
		return
	}

	tr := options.Transform
	scale := tr.Scale()
	pos := tr.Position()

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale.X(), scale.Y())
	op.GeoM.Rotate(tr.Rotation())
	op.GeoM.Translate(pos.X(), pos.Y())

	e.img.DrawImage(source.img, op)
}

func (e *ebitenImage) Dimensions() (width, height float64) {
	b := e.img.Bounds()
	return float64(b.Dx()), float64(b.Dy())
}

func (e *ebitenImage) SetDimensions(width, height float64) {
	// ebiten images have a fixed size once created; nothing to do here.
}
