package ui

import (
	"slices"

	"survivor/internal/domain/graph"
)

// HUD owns a flat collection of screen-space widgets (currently Text,
// later also graphs/bars) and draws them on top of the world, unaffected
// by the camera.
type HUD struct {
	widgets []Widget
}

// NewHUD creates an empty HUD.
func NewHUD() *HUD {
	return &HUD{widgets: make([]Widget, 0)}
}

// AddWidget adds a widget to the HUD.
func (h *HUD) AddWidget(w Widget) {
	h.widgets = append(h.widgets, w)
}

// RemoveWidget removes the given widget from the HUD, if present.
func (h *HUD) RemoveWidget(w Widget) bool {
	i := slices.Index(h.widgets, w)
	if i == -1 {
		return false
	}

	h.widgets = slices.Delete(h.widgets, i, i+1)

	return true
}

// Draw draws all widgets onto target in screen space.
func (h *HUD) Draw(target graph.Image) {
	for _, w := range h.widgets {
		w.Draw(target)
	}
}
