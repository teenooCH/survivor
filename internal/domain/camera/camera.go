// package camera provides the screen view into the world space.
// The camera represents the physical screen and is capable of
// following a node. Also it is possible to "shake" the screen.
package camera

import (
	"math/rand/v2"

	"survivor/internal/domain/graph"
	"survivor/internal/domain/node2D"
	"survivor/internal/domain/transform"
)

// Camera provides a viewport into the world and can follow a transformable Node.
type Camera struct {
	node2D.Node2D
	width        uint
	height       uint
	surface      graph.Image
	nodeToFollow transform.Transformable
	shakeMag     float64 // current screen-shake magnitude in pixels; decays each frame (ch13)
}

// New creates a camera with the given dimensions from the surface.
func New(surface graph.Image) *Camera {
	w, h := surface.Dimensions()

	return &Camera{
		Node2D:  *node2D.New("camera"),
		width:   uint(w),
		height:  uint(h),
		surface: surface,
	}
}

// GetSurface returns the offscreen image the scene is drawn to.
func (c *Camera) GetSurface() graph.Image {
	return c.surface
}

// GetWidth returns the camera viewport width in pixels.
func (c *Camera) GetWidth() uint { return c.width }

// GetHeight returns the camera viewport height in pixels.
func (c *Camera) GetHeight() uint { return c.height }

// SetFollow sets the node to follow. Pass nil to disable.
func (c *Camera) SetFollow(node transform.Transformable) {
	c.nodeToFollow = node
}

// Update updates camera position to follow the target node (center on screen).
// Also a shake is applied if set.
func (c *Camera) Update() {
	if c.nodeToFollow == nil {
		return
	}

	// Center the target in the view: camera top-left = target center - half screen
	wt := c.nodeToFollow.GetWorldTransform()
	pos := wt.Position()
	c.SetPosition(pos.X()-float64(c.width)/2, pos.Y()-float64(c.height)/2)

	// Apply a decaying screen shake on top of the follow position (ch13). Because
	// the world and particles are both drawn relative to the camera position, they
	// shake together while the screen-space HUD stays still.
	if c.shakeMag > 0.4 {
		ox := (rand.Float64()*2 - 1) * c.shakeMag
		oy := (rand.Float64()*2 - 1) * c.shakeMag
		c.SetPosition(c.GetPosition().X()+ox, c.GetPosition().Y()+oy)
		c.shakeMag *= 0.85
	} else {
		c.shakeMag = 0
	}
}

// Shake starts (or strengthens) a brief screen shake. magnitude is in pixels and
// decays to zero over a few frames. Call it on impactful events such as an enemy
// death or the player taking damage.
func (c *Camera) Shake(magnitude float64) {
	if magnitude > c.shakeMag {
		c.shakeMag = magnitude
	}
}

// ApplyOffset modifies tr so world coords are drawn relative to camera position.
func (c *Camera) ApplyOffset(tr *transform.Transform) {
	pos := c.GetPosition()
	tr.Translate(-pos.X(), -pos.Y())
}

// DrawToScreen draws the camera surface to the screen.
// It's not called Draw so it doesn't implement Drawable.
func (c *Camera) DrawToScreen(screen graph.Image) {
	screen.DrawImage(c.surface, graph.DrawOpt{Transform: transform.NewZero()})
}

// GetWorldCoords converts screen coordinates (e.g. mouse) to world coordinates.
func (c *Camera) GetWorldCoords(screenX, screenY int) (worldX, worldY float64) {
	pos := c.GetPosition()
	return float64(screenX) + pos.X(), float64(screenY) + pos.Y()
}
