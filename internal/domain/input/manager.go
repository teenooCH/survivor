// Package input translates hardware button state into semantic game Actions.
// It depends only on the ports.InputProvider interface (dependency inversion),
// never on a concrete input library, so gameplay code stays testable and the
// hardware backend can be swapped without touching domain or application code.
package input

import "survivor/internal/ports"

// bindingKind identifies which physical input device a binding refers to.
type bindingKind int

const (
	bindingKindKey bindingKind = iota
	bindingKindMouseButton
	bindingKindGamepadButton
)

// binding is a single physical input bound to an Action. A gamepad binding
// without an explicit ID matches any currently connected gamepad, so games
// don't need to know a controller's ID up front.
type binding struct {
	kind          bindingKind
	key           Key
	mouseButton   MouseButton
	gamepadButton GamepadButton
	gamepadID     GamepadID
	anyGamepad    bool
}

// Manager binds Actions to one or more physical inputs (keys, mouse
// buttons, or gamepad buttons) and resolves their state through a
// ports.InputProvider.
type Manager struct {
	provider ports.InputProvider
	bindings map[Action][]binding
}

// NewManager creates an input Manager driven by the given provider.
// The provider is the outgoing interface implementation, supplied by the
// composition root (e.g. cmd/*/main.go) from the infrastructure layer.
func NewManager(provider ports.InputProvider) *Manager {
	return &Manager{
		provider: provider,
		bindings: make(map[Action][]binding),
	}
}

// BindKey binds one or more Keys to an Action. Existing bindings for the
// Action are preserved, so multiple calls add alternative inputs (e.g. WASD
// and arrow keys both triggering movement).
func (m *Manager) BindKey(action Action, keys ...Key) {
	for _, key := range keys {
		m.bindings[action] = append(m.bindings[action], binding{kind: bindingKindKey, key: key})
	}
}

// BindMouseButton binds one or more mouse buttons to an Action.
func (m *Manager) BindMouseButton(action Action, buttons ...MouseButton) {
	for _, button := range buttons {
		m.bindings[action] = append(m.bindings[action], binding{kind: bindingKindMouseButton, mouseButton: button})
	}
}

// BindGamepadButton binds a gamepad button to an Action. If no GamepadID is
// given, the binding matches the button on any currently connected gamepad,
// so single-player games can support a controller without knowing its ID in
// advance. Pass explicit IDs (see ConnectedGamepadIDs) to restrict the
// binding to specific controllers, e.g. for local multiplayer.
func (m *Manager) BindGamepadButton(action Action, button GamepadButton, ids ...GamepadID) {
	if len(ids) == 0 {
		m.bindings[action] = append(m.bindings[action], binding{
			kind: bindingKindGamepadButton, gamepadButton: button, anyGamepad: true,
		})

		return
	}

	for _, id := range ids {
		m.bindings[action] = append(m.bindings[action], binding{
			kind: bindingKindGamepadButton, gamepadButton: button, gamepadID: id,
		})
	}
}

// ClearBindings removes all bindings (keys, mouse and gamepad buttons) for
// the given Action.
func (m *Manager) ClearBindings(action Action) {
	delete(m.bindings, action)
}

// ConnectedGamepadIDs returns the IDs of all currently connected gamepads.
func (m *Manager) ConnectedGamepadIDs() []GamepadID {
	return m.provider.ConnectedGamepadIDs()
}

// IsActionPressed reports whether any input bound to action is currently held.
func (m *Manager) IsActionPressed(action Action) bool {
	for _, b := range m.bindings[action] {
		if m.isBindingPressed(b) {
			return true
		}
	}

	return false
}

// IsActionJustPressed reports whether any input bound to action transitioned
// to pressed this frame.
func (m *Manager) IsActionJustPressed(action Action) bool {
	for _, b := range m.bindings[action] {
		if m.isBindingJustPressed(b) {
			return true
		}
	}

	return false
}

// IsActionJustReleased reports whether any input bound to action transitioned
// to released this frame.
func (m *Manager) IsActionJustReleased(action Action) bool {
	for _, b := range m.bindings[action] {
		if m.isBindingJustReleased(b) {
			return true
		}
	}

	return false
}

func (m *Manager) isBindingPressed(b binding) bool {
	switch b.kind {
	case bindingKindKey:
		return m.provider.IsKeyPressed(b.key)
	case bindingKindMouseButton:
		return m.provider.IsMouseButtonPressed(b.mouseButton)
	case bindingKindGamepadButton:
		return m.isGamepadBindingActive(b, m.provider.IsGamepadButtonPressed)
	default:
		return false
	}
}

func (m *Manager) isBindingJustPressed(b binding) bool {
	switch b.kind {
	case bindingKindKey:
		return m.provider.IsKeyJustPressed(b.key)
	case bindingKindMouseButton:
		return m.provider.IsMouseButtonJustPressed(b.mouseButton)
	case bindingKindGamepadButton:
		return m.isGamepadBindingActive(b, m.provider.IsGamepadButtonJustPressed)
	default:
		return false
	}
}

func (m *Manager) isBindingJustReleased(b binding) bool {
	switch b.kind {
	case bindingKindKey:
		return m.provider.IsKeyJustReleased(b.key)
	case bindingKindMouseButton:
		return m.provider.IsMouseButtonJustReleased(b.mouseButton)
	case bindingKindGamepadButton:
		return m.isGamepadBindingActive(b, m.provider.IsGamepadButtonJustReleased)
	default:
		return false
	}
}

// isGamepadBindingActive resolves a gamepad binding against check, expanding
// wildcard (anyGamepad) bindings to every currently connected gamepad ID.
func (m *Manager) isGamepadBindingActive(b binding, check func(GamepadID, GamepadButton) bool) bool {
	if !b.anyGamepad {
		return check(b.gamepadID, b.gamepadButton)
	}

	for _, id := range m.provider.ConnectedGamepadIDs() {
		if check(id, b.gamepadButton) {
			return true
		}
	}

	return false
}
