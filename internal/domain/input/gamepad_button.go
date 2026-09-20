package input

import "survivor/internal/ports"

// GamepadID and GamepadButton re-export their ports counterparts so
// gameplay code can bind Actions without importing the ports package directly.
type (
	GamepadID     = ports.GamepadID
	GamepadButton = ports.GamepadButton
)

const (
	GamepadButtonSouth = ports.GamepadButtonSouth
	GamepadButtonEast  = ports.GamepadButtonEast
	GamepadButtonWest  = ports.GamepadButtonWest
	GamepadButtonNorth = ports.GamepadButtonNorth

	GamepadButtonLeftBumper   = ports.GamepadButtonLeftBumper
	GamepadButtonRightBumper  = ports.GamepadButtonRightBumper
	GamepadButtonLeftTrigger  = ports.GamepadButtonLeftTrigger
	GamepadButtonRightTrigger = ports.GamepadButtonRightTrigger

	GamepadButtonBack  = ports.GamepadButtonBack
	GamepadButtonStart = ports.GamepadButtonStart

	GamepadButtonLeftStick  = ports.GamepadButtonLeftStick
	GamepadButtonRightStick = ports.GamepadButtonRightStick

	GamepadButtonDPadUp    = ports.GamepadButtonDPadUp
	GamepadButtonDPadDown  = ports.GamepadButtonDPadDown
	GamepadButtonDPadLeft  = ports.GamepadButtonDPadLeft
	GamepadButtonDPadRight = ports.GamepadButtonDPadRight
)
