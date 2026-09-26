// Package ui provides screen-space HUD widgets (e.g. text) that are drawn
// directly onto the render target, independent of the world's camera.
package ui

import "survivor/internal/domain/graph"

// Widget is a single HUD element that can draw itself in screen space.
// Future widgets (e.g. a health graph/bar) implement this interface to
// become usable by HUD alongside Text.
type Widget interface {
	Draw(target graph.Image)
}
