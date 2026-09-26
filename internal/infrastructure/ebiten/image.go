package ebiten

import (
	"image/color"

	"survivor/internal/domain/graph"

	"github.com/hajimehoshi/bitmapfont/v4"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// defaultFontSize is the natural pixel size of hudFace's glyphs. TextOpt.Size
// is applied as a scale factor relative to this size.
const defaultFontSize = 12

// textPadding is the space in pixels kept between the text and the edge of
// its background/border box.
const textPadding = 4

// hudFace is the shared font used to render HUD text.
var hudFace = text.NewGoXFace(bitmapfont.Face)

// ebitenImage adapts *ebiten.Image to the graph.Image interface.
type ebitenImage struct {
	img *ebiten.Image
}

var _ graph.TextRenderer = (*ebitenImage)(nil)

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

// DrawText renders a HUD text widget: an optional background fill, an
// optional border, and the text itself, in that order.
func (e *ebitenImage) DrawText(opts graph.TextOpt) {
	size := opts.Size
	if size <= 0 {
		size = defaultFontSize
	}
	scale := size / defaultFontSize

	textWidth, textHeight := text.Measure(opts.Text, hudFace, defaultFontSize)
	boxWidth := textWidth*scale + 2*textPadding
	boxHeight := textHeight*scale + 2*textPadding

	if !opts.Background.Transparent && opts.Background.Color != nil {
		e.FillRect(opts.X, opts.Y, boxWidth, boxHeight, opts.Background.Color)
	}

	if opts.Border.Visible {
		e.StrokeRect(opts.X, opts.Y, boxWidth, boxHeight, opts.Border.Width, opts.Border.Color)
	}

	clr := opts.Color
	if clr == nil {
		clr = color.White
	}

	drawOpt := &text.DrawOptions{}
	drawOpt.GeoM.Scale(scale, scale)
	drawOpt.GeoM.Translate(opts.X+textPadding, opts.Y+textPadding)
	drawOpt.ColorScale.SetWithColor(clr)

	text.Draw(e.img, opts.Text, hudFace, drawOpt)
}

// FillRect draws a filled rectangle, e.g. for a graph/bar widget.
func (e *ebitenImage) FillRect(x, y, width, height float64, clr color.Color) {
	if clr == nil {
		return
	}

	vector.FillRect(e.img, float32(x), float32(y), float32(width), float32(height), clr, false)
}

// StrokeRect draws a rectangle outline, e.g. for a widget's border.
func (e *ebitenImage) StrokeRect(x, y, width, height, strokeWidth float64, clr color.Color) {
	if clr == nil || strokeWidth <= 0 {
		return
	}

	vector.StrokeRect(e.img, float32(x), float32(y), float32(width), float32(height), float32(strokeWidth), clr, false)
}
