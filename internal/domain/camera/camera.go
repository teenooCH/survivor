// package camera provides the screen view into the world space.
// The camera represents the physical screen and is capable of
// following a node.
package camera

import (
	"survivor/internal/domain/node2D"
	"survivor/internal/domain/transform"
)

// Camera provides a viewport into the world and can follow a transformable Node.
type Camera struct {
	node2D.Node2D
	width        uint
	height       uint
	nodeToFollow transform.Transformable
}

// New creates a camera with the given dimensions.
func New(width, height uint) *Camera {
	return &Camera{
		Node2D: *node2D.New("camera"),
		width:  width,
		height: height,
	}
}

// GetWidth returns the camera viewport width in pixels.
func (c *Camera) GetWidth() uint { return c.width }

// GetHeight returns the camera viewport height in pixels.
func (c *Camera) GetHeight() uint { return c.height }

// GetHeight returns the camera viewport height in pixels.
func (c *Camera) SetFollow(node transform.Transformable) {
	c.nodeToFollow = node
}

// Update updates camera position to follow the target node (center on screen).
func (c *Camera) Update() {
	if c.nodeToFollow == nil {
		return
	}

	// Center the target in the view: camera top-left = target center - half screen
	wt := c.nodeToFollow.GetWorldTransform()
	pos := wt.Position()
	c.SetPosition(pos.X()-float64(c.width)/2, pos.Y()-float64(c.height)/2)
}

// ScreenToWorldCoords converts screen coordinates (e.g. mouse) to world coordinates.
func (c *Camera) ScreenToWorldCoords(screenX, screenY int) (worldX, worldY float64) {
	pos := c.GetPosition()
	return float64(screenX) + pos.X(), float64(screenY) + pos.Y()
}
