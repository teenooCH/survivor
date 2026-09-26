package ui

import "image/color"

// Background describes the fill behind a HUD widget: either a solid color
// or fully transparent (no fill drawn).
type Background struct {
	color       color.Color
	transparent bool
}

// NewBackground creates an opaque/solid colored background.
func NewBackground(clr color.Color) Background {
	return Background{color: clr}
}

// NewTransparentBackground creates a background that draws no fill.
func NewTransparentBackground() Background {
	return Background{transparent: true}
}

func (b Background) Color() color.Color { return b.color }
func (b Background) Transparent() bool  { return b.transparent }

// Border describes an optional frame drawn around a HUD widget.
type Border struct {
	color   color.Color
	width   float64
	visible bool
}

// NewBorder creates a visible border with the given color and stroke width.
func NewBorder(clr color.Color, width float64) Border {
	return Border{color: clr, width: width, visible: width > 0}
}

// NoBorder creates a border that is not drawn.
func NoBorder() Border {
	return Border{}
}

func (b Border) Color() color.Color { return b.color }
func (b Border) Width() float64     { return b.width }
func (b Border) Visible() bool      { return b.visible }
