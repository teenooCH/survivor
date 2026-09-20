package input

import "survivor/internal/ports"

// Key re-exports ports.Key so gameplay code can bind Actions without
// importing the ports package directly.
type Key = ports.Key

const (
	KeyUnknown = ports.KeyUnknown

	KeyW = ports.KeyW
	KeyA = ports.KeyA
	KeyS = ports.KeyS
	KeyD = ports.KeyD

	KeyUp    = ports.KeyUp
	KeyDown  = ports.KeyDown
	KeyLeft  = ports.KeyLeft
	KeyRight = ports.KeyRight

	KeySpace  = ports.KeySpace
	KeyEscape = ports.KeyEscape
	KeyEnter  = ports.KeyEnter
)
