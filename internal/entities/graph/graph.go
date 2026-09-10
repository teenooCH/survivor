package graph

import "github.com/teenooCH/survivor/internal/entities/transform"

// Image is the interface for any image that can be drawn on top of another image.
type Image interface {
	// DrawImage draws the image on top the receiver image
	// with the given options.
	DrawImage(image Image, options transform.Transform)
	// Dimensions returns the width and height in pixels.
	Dimensions() (width, height float64)
}
