package graph

import "survivor/internal/domain/transform"

// Image is the interface for any image that can be drawn on top of another image.
type Image interface {
	// DrawImage draws the image on top the receiver image
	// with the given options.
	DrawImage(image Image, options DrawOpt)
	// Dimensions returns the width and height in pixels.
	Dimensions() (width, height float64)
	SetDimensions(width, height float64)
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
