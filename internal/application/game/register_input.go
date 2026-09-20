package game

import (
	"survivor/internal/application/engine"
	"survivor/internal/domain/input"
)

// RegisterDefaultBindings binds the default keyboard controls (WASD and
// arrow keys for movement) to the engine's input Manager. This use case is
// invoked once from the composition root (e.g. cmd/*/main.go) after the
// engine has been created.
func RegisterDefaultBindings(engine *engine.Engine) {
	im := engine.InputManager()

	im.BindKey(input.ActionMoveUp, input.KeyW, input.KeyUp)
	im.BindKey(input.ActionMoveDown, input.KeyS, input.KeyDown)
	im.BindKey(input.ActionMoveLeft, input.KeyA, input.KeyLeft)
	im.BindKey(input.ActionMoveRight, input.KeyD, input.KeyRight)
}
