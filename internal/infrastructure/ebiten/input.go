package ebiten

import (
	"survivor/internal/ports"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// keyMapping translates the domain's hardware-agnostic ports.Key into a
// concrete ebiten.Key. It is the single place where the domain's Key type
// meets the ebiten library.
var keyMapping = map[ports.Key]ebiten.Key{
	ports.KeyW: ebiten.KeyW,
	ports.KeyA: ebiten.KeyA,
	ports.KeyS: ebiten.KeyS,
	ports.KeyD: ebiten.KeyD,

	ports.KeyUp:    ebiten.KeyArrowUp,
	ports.KeyDown:  ebiten.KeyArrowDown,
	ports.KeyLeft:  ebiten.KeyArrowLeft,
	ports.KeyRight: ebiten.KeyArrowRight,

	ports.KeySpace:  ebiten.KeySpace,
	ports.KeyEscape: ebiten.KeyEscape,
	ports.KeyEnter:  ebiten.KeyEnter,
}

// mouseButtonMapping translates ports.MouseButton into a concrete ebiten.MouseButton.
var mouseButtonMapping = map[ports.MouseButton]ebiten.MouseButton{
	ports.MouseButtonLeft:   ebiten.MouseButtonLeft,
	ports.MouseButtonRight:  ebiten.MouseButtonRight,
	ports.MouseButtonMiddle: ebiten.MouseButtonMiddle,
}

// gamepadButtonMapping translates ports.GamepadButton into ebiten's standard
// gamepad layout, so the same binding works consistently across different
// controller hardware.
var gamepadButtonMapping = map[ports.GamepadButton]ebiten.StandardGamepadButton{
	ports.GamepadButtonSouth: ebiten.StandardGamepadButtonRightBottom,
	ports.GamepadButtonEast:  ebiten.StandardGamepadButtonRightRight,
	ports.GamepadButtonWest:  ebiten.StandardGamepadButtonRightLeft,
	ports.GamepadButtonNorth: ebiten.StandardGamepadButtonRightTop,

	ports.GamepadButtonLeftBumper:   ebiten.StandardGamepadButtonFrontTopLeft,
	ports.GamepadButtonRightBumper:  ebiten.StandardGamepadButtonFrontTopRight,
	ports.GamepadButtonLeftTrigger:  ebiten.StandardGamepadButtonFrontBottomLeft,
	ports.GamepadButtonRightTrigger: ebiten.StandardGamepadButtonFrontBottomRight,

	ports.GamepadButtonBack:  ebiten.StandardGamepadButtonCenterLeft,
	ports.GamepadButtonStart: ebiten.StandardGamepadButtonCenterRight,

	ports.GamepadButtonLeftStick:  ebiten.StandardGamepadButtonLeftStick,
	ports.GamepadButtonRightStick: ebiten.StandardGamepadButtonRightStick,

	ports.GamepadButtonDPadUp:    ebiten.StandardGamepadButtonLeftTop,
	ports.GamepadButtonDPadDown:  ebiten.StandardGamepadButtonLeftBottom,
	ports.GamepadButtonDPadLeft:  ebiten.StandardGamepadButtonLeftLeft,
	ports.GamepadButtonDPadRight: ebiten.StandardGamepadButtonLeftRight,
}

// InputProvider implements ports.InputProvider using the ebiten library.
// It is the outgoing interface's concrete adapter and is injected into
// domain/input.NewManager from the composition root (e.g. cmd/*/main.go).
type InputProvider struct{}

func NewInputProvider() *InputProvider {
	return &InputProvider{}
}

func (p *InputProvider) IsKeyPressed(key ports.Key) bool {
	ebKey, ok := keyMapping[key]
	if !ok {
		return false
	}

	return ebiten.IsKeyPressed(ebKey)
}

func (p *InputProvider) IsKeyJustPressed(key ports.Key) bool {
	ebKey, ok := keyMapping[key]
	if !ok {
		return false
	}

	return inpututil.IsKeyJustPressed(ebKey)
}

func (p *InputProvider) IsKeyJustReleased(key ports.Key) bool {
	ebKey, ok := keyMapping[key]
	if !ok {
		return false
	}

	return inpututil.IsKeyJustReleased(ebKey)
}

func (p *InputProvider) IsMouseButtonPressed(button ports.MouseButton) bool {
	ebButton, ok := mouseButtonMapping[button]
	if !ok {
		return false
	}

	return ebiten.IsMouseButtonPressed(ebButton)
}

func (p *InputProvider) IsMouseButtonJustPressed(button ports.MouseButton) bool {
	ebButton, ok := mouseButtonMapping[button]
	if !ok {
		return false
	}

	return inpututil.IsMouseButtonJustPressed(ebButton)
}

func (p *InputProvider) IsMouseButtonJustReleased(button ports.MouseButton) bool {
	ebButton, ok := mouseButtonMapping[button]
	if !ok {
		return false
	}

	return inpututil.IsMouseButtonJustReleased(ebButton)
}

func (p *InputProvider) IsGamepadButtonPressed(id ports.GamepadID, button ports.GamepadButton) bool {
	ebButton, ok := gamepadButtonMapping[button]
	if !ok {
		return false
	}

	return ebiten.IsStandardGamepadButtonPressed(ebiten.GamepadID(id), ebButton)
}

func (p *InputProvider) IsGamepadButtonJustPressed(id ports.GamepadID, button ports.GamepadButton) bool {
	ebButton, ok := gamepadButtonMapping[button]
	if !ok {
		return false
	}

	return inpututil.IsStandardGamepadButtonJustPressed(ebiten.GamepadID(id), ebButton)
}

func (p *InputProvider) IsGamepadButtonJustReleased(id ports.GamepadID, button ports.GamepadButton) bool {
	ebButton, ok := gamepadButtonMapping[button]
	if !ok {
		return false
	}

	return inpututil.IsStandardGamepadButtonJustReleased(ebiten.GamepadID(id), ebButton)
}

// ConnectedGamepadIDs returns the IDs of all currently connected gamepads,
// letting the domain layer discover and bind controllers without any
// hardware-specific setup.
func (p *InputProvider) ConnectedGamepadIDs() []ports.GamepadID {
	ebIDs := ebiten.AppendGamepadIDs(nil)

	ids := make([]ports.GamepadID, len(ebIDs))
	for i, ebID := range ebIDs {
		ids[i] = ports.GamepadID(ebID)
	}

	return ids
}
