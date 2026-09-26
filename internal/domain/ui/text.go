package ui

import (
	"image/color"

	"survivor/internal/domain/graph"
	"survivor/internal/domain/vector"
)

// Text is a HUD widget that draws a text string at a fixed screen
// position, with a parametrizable color, size, background and border.
type Text struct {
	text       string
	color      color.Color
	position   vector.Vector
	size       float64
	background Background
	border     Border
	visible    bool
}

// NewText creates a HUD text widget at (x, y) with white text, size 16,
// a transparent background and no border.
func NewText(text string, x, y float64) *Text {
	return &Text{
		text:       text,
		color:      color.White,
		position:   vector.New(x, y),
		size:       16,
		background: NewTransparentBackground(),
		border:     NoBorder(),
		visible:    true,
	}
}

func (t *Text) GetText() string     { return t.text }
func (t *Text) SetText(text string) { t.text = text }

func (t *Text) GetColor() color.Color    { return t.color }
func (t *Text) SetColor(clr color.Color) { t.color = clr }

func (t *Text) GetPosition() vector.Vector { return t.position }
func (t *Text) SetPosition(x, y float64)   { t.position = vector.New(x, y) }

func (t *Text) GetSize() float64     { return t.size }
func (t *Text) SetSize(size float64) { t.size = size }

func (t *Text) GetBackground() Background   { return t.background }
func (t *Text) SetBackground(bg Background) { t.background = bg }

func (t *Text) GetBorder() Border       { return t.border }
func (t *Text) SetBorder(border Border) { t.border = border }

func (t *Text) IsVisible() bool         { return t.visible }
func (t *Text) SetVisible(visible bool) { t.visible = visible }

// Draw renders the text, and its optional background/border, onto target
// if target supports text rendering (see graph.TextRenderer).
func (t *Text) Draw(target graph.Image) {
	if !t.visible {
		return
	}

	renderer, ok := target.(graph.TextRenderer)
	if !ok {
		return
	}

	renderer.DrawText(graph.TextOpt{
		Text:  t.text,
		X:     t.position.X(),
		Y:     t.position.Y(),
		Size:  t.size,
		Color: t.color,
		Background: graph.BackgroundOpt{
			Color:       t.background.Color(),
			Transparent: t.background.Transparent(),
		},
		Border: graph.BorderOpt{
			Color:   t.border.Color(),
			Width:   t.border.Width(),
			Visible: t.border.Visible(),
		},
	})
}
