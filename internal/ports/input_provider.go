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
}
