package ports

// Key is a hardware-agnostic identifier for a physical input button. It is
// declared here, at the boundary, so both the domain (which binds Keys to
// Actions) and the infrastructure (which resolves Keys against real
// hardware) can depend on it without creating an import cycle between them.
type Key int

const (
	KeyUnknown Key = iota

	KeyW
	KeyA
	KeyS
	KeyD

	KeyUp
	KeyDown
	KeyLeft
	KeyRight

	KeySpace
	KeyEscape
	KeyEnter
)

// MouseButton is a hardware-agnostic identifier for a mouse button.
type MouseButton int

const (
	MouseButtonLeft MouseButton = iota
	MouseButtonRight
	MouseButtonMiddle
)

// GamepadID identifies a connected gamepad. It is opaque and only meant to
// be obtained from InputProvider.ConnectedGamepadIDs, then reused to query
// or bind buttons on that specific controller.
type GamepadID int

// GamepadButton is a hardware-agnostic identifier for a gamepad button,
// following the standard W3C gamepad layout so it maps consistently across
// different controller hardware (Xbox, PlayStation, etc.).
type GamepadButton int

const (
	GamepadButtonSouth GamepadButton = iota // e.g. Xbox A / PlayStation Cross
	GamepadButtonEast                       // e.g. Xbox B / PlayStation Circle
	GamepadButtonWest                       // e.g. Xbox X / PlayStation Square
	GamepadButtonNorth                      // e.g. Xbox Y / PlayStation Triangle

	GamepadButtonLeftBumper
	GamepadButtonRightBumper
	GamepadButtonLeftTrigger
	GamepadButtonRightTrigger

	GamepadButtonBack
	GamepadButtonStart

	GamepadButtonLeftStick
	GamepadButtonRightStick

	GamepadButtonDPadUp
	GamepadButtonDPadDown
	GamepadButtonDPadLeft
	GamepadButtonDPadRight
)

// InputProvider is the outgoing (driven) interface the domain layer uses to
// query the state of physical input devices. It is defined here, at the
// boundary of the application, so that domain/input.Manager can depend on
// it without knowing about any concrete input library. Infrastructure
// packages (e.g. internal/infrastructure/ebiten) implement this interface,
// and the concrete implementation is injected at the composition root
// (cmd/*/main.go) - the same wiring pattern used for TextureProvider.
type InputProvider interface {
	// IsKeyPressed reports whether key is currently held down.
	IsKeyPressed(key Key) bool
	// IsKeyJustPressed reports whether key transitioned to pressed this frame.
	IsKeyJustPressed(key Key) bool
	// IsKeyJustReleased reports whether key transitioned to released this frame.
	IsKeyJustReleased(key Key) bool

	// IsMouseButtonPressed reports whether button is currently held down.
	IsMouseButtonPressed(button MouseButton) bool
	// IsMouseButtonJustPressed reports whether button transitioned to pressed this frame.
	IsMouseButtonJustPressed(button MouseButton) bool
	// IsMouseButtonJustReleased reports whether button transitioned to released this frame.
	IsMouseButtonJustReleased(button MouseButton) bool

	// IsGamepadButtonPressed reports whether button is currently held down on the given gamepad.
	IsGamepadButtonPressed(id GamepadID, button GamepadButton) bool
	// IsGamepadButtonJustPressed reports whether button transitioned to pressed this frame on the given gamepad.
	IsGamepadButtonJustPressed(id GamepadID, button GamepadButton) bool
	// IsGamepadButtonJustReleased reports whether button transitioned to released this frame on the given gamepad.
	IsGamepadButtonJustReleased(id GamepadID, button GamepadButton) bool
	// ConnectedGamepadIDs returns the IDs of all currently connected gamepads.
	ConnectedGamepadIDs() []GamepadID
}
