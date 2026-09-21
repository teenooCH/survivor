package tile

import "survivor/internal/domain/vector"

// Viewport is the portion of the world a TileMap needs to keep loaded.
// domain/camera.Camera satisfies this interface.
type Viewport interface {
	GetPosition() vector.Vector
	GetWidth() uint
	GetHeight() uint
}
