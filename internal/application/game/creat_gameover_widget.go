package game

import (
	"image/color"

	"survivor/internal/domain/ui"
)

// Use case for creating a "Game Over" widget.

func CreateGameOverWidget(x, y float64) *ui.Text {
	hudText := ui.NewText("Game Over", x, y)
	hudText.SetColor(color.RGBA{R: 255, G: 0, B: 0, A: 255})
	hudText.SetSize(32)
	hudText.SetBackground(ui.NewTransparentBackground())
	hudText.SetVisible(false)

	return hudText
}
