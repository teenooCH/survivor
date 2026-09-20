// Package input translates hardware button state into semantic game Actions.
// It depends only on the ports.InputProvider interface (dependency inversion),
// never on a concrete input library, so gameplay code stays testable and the
// hardware backend can be swapped without touching domain or application code.
package input

import "survivor/internal/ports"

// Manager binds Actions to one or more Keys and resolves their state
// through a ports.InputProvider.
type Manager struct {
	provider ports.InputProvider
	bindings map[Action][]Key
}

// NewManager creates an input Manager driven by the given provider.
// The provider is the outgoing interface implementation, supplied by the
// composition root (e.g. cmd/*/main.go) from the infrastructure layer.
func NewManager(provider ports.InputProvider) *Manager {
	return &Manager{
		provider: provider,
		bindings: make(map[Action][]Key),
	}
}

// BindKey binds one or more Keys to an Action. Existing bindings for the
// Action are preserved, so multiple calls add alternative keys (e.g. WASD
// and arrow keys both triggering movement).
func (m *Manager) BindKey(action Action, keys ...Key) {
	m.bindings[action] = append(m.bindings[action], keys...)
}

// ClearBindings removes all key bindings for the given Action.
func (m *Manager) ClearBindings(action Action) {
	delete(m.bindings, action)
}

// IsActionPressed reports whether any Key bound to action is currently held.
func (m *Manager) IsActionPressed(action Action) bool {
	for _, key := range m.bindings[action] {
		if m.provider.IsKeyPressed(key) {
			return true
		}
	}

	return false
}

// IsActionJustPressed reports whether any Key bound to action transitioned
// to pressed this frame.
func (m *Manager) IsActionJustPressed(action Action) bool {
	for _, key := range m.bindings[action] {
		if m.provider.IsKeyJustPressed(key) {
			return true
		}
	}

	return false
}

// IsActionJustReleased reports whether any Key bound to action transitioned
// to released this frame.
func (m *Manager) IsActionJustReleased(action Action) bool {
	for _, key := range m.bindings[action] {
		if m.provider.IsKeyJustReleased(key) {
			return true
		}
	}

	return false
}
