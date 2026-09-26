package graph

import (
	"image/color"

	"survivor/internal/domain/transform"
)

// Image is the interface for any image that can be drawn on top of another image.
type Image interface {
	// DrawImage draws the image on top the receiver image
	// with the given options.
	DrawImage(image Image, options DrawOpt)
	// Dimensions returns the width and height in pixels.
	Dimensions() (width, height float64)
	SetDimensions(width, height float64)
}

// TextRenderer is an optional capability of an Image that can render text
// and simple filled/stroked shapes in screen space. It is implemented by
// the infrastructure's render target and used by the ui package to draw
// HUD widgets (text now, graphs/bars later) directly on top of the world.
type TextRenderer interface {
	// DrawText renders a text widget (with optional background and border)
	// described by opts onto the receiver image.
	DrawText(opts TextOpt)
	// FillRect draws a filled rectangle, e.g. for a graph/bar widget.
	FillRect(x, y, width, height float64, clr color.Color)
	// StrokeRect draws a rectangle outline, e.g. for a widget's border.
	StrokeRect(x, y, width, height, strokeWidth float64, clr color.Color)
}

// TextOpt represents the options for drawing a HUD text widget.
type TextOpt struct {
	Text       string
	X, Y       float64
	Size       float64
	Color      color.Color
	Background BackgroundOpt
	Border     BorderOpt
}

// BackgroundOpt describes the fill behind a HUD widget.
type BackgroundOpt struct {
	Color       color.Color
	Transparent bool
}

// BorderOpt describes an optional frame drawn around a HUD widget.
type BorderOpt struct {
	Color   color.Color
	Width   float64
	Visible bool
}

// Drawable is the interface for any object that can be drawn on top of an image.
type Drawable interface {
	transform.Transformable
	GetLayer() int
	// Draw draws the image of the Drawable node on the
	// given target image with the provided options.
	Draw(target Image, options DrawOpt)
}

// DrawOpt represents the options for DrawImage.
type DrawOpt struct {
	Transform transform.Transform
}
