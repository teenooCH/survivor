package input

import "survivor/internal/ports"

// MouseButton re-exports ports.MouseButton so gameplay code can bind
// Actions without importing the ports package directly.
type MouseButton = ports.MouseButton

const (
	MouseButtonLeft   = ports.MouseButtonLeft
	MouseButtonRight  = ports.MouseButtonRight
	MouseButtonMiddle = ports.MouseButtonMiddle
)
