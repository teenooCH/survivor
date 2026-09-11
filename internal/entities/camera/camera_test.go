package camera_test

import (
	"math"
	"testing"

	"github.com/teenooCH/survivor/internal/entities/camera"
	"github.com/teenooCH/survivor/internal/entities/graph"
	"github.com/teenooCH/survivor/internal/entities/node2D"
	"github.com/teenooCH/survivor/internal/entities/vector"
)

func TestCamera_Update_FollowTheNode(t *testing.T) {
	t.Run("", func(t *testing.T) {
		nf := node2D.New("testNode")
		nf.SetPosition(1000, 750)

		c := camera.New(newMockImage(200, 100))
		c.Update()
		comparePosition(t, c.GetPosition(), vector.New(0, 0))

		c.SetFollow(nf)
		c.Update()
		comparePosition(t, c.GetPosition(), vector.New(900, 700))

		nf.SetPosition(2000, 1500)
		c.Update()
		comparePosition(t, c.GetPosition(), vector.New(1900, 1450))
	})
}

func TestCamera_Shake_DecaysAndTurnsOff(t *testing.T) {
	nf := node2D.New("testNode")
	nf.SetPosition(1000, 750)

	c := camera.New(newMockImage(200, 100))
	c.SetFollow(nf)

	const magnitude = 10.0
	c.Shake(magnitude)

	basePos := vector.New(900, 700)
	bound := magnitude

	// While the shake is still active, each update must offset the camera by no
	// more than the current (decaying) magnitude, and the magnitude shrinks by
	// a factor of 0.85 every cycle.
	for bound > 0.4 {
		c.Update()

		offX := c.GetPosition().X() - basePos.X()
		offY := c.GetPosition().Y() - basePos.Y()

		if math.Abs(offX) > bound || math.Abs(offY) > bound {
			t.Fatalf("expected shake offset within +/-%v, got (%v, %v)", bound, offX, offY)
		}

		bound *= 0.85
	}

	// Once the magnitude has decayed below the threshold, the shake must turn
	// off entirely and the camera should exactly track the followed node.
	c.Update()
	comparePosition(t, c.GetPosition(), basePos)

	c.Update()
	comparePosition(t, c.GetPosition(), basePos)
}

func comparePosition(t *testing.T, a, b vector.Vector) {
	t.Helper()

	if a != b {
		t.Fatalf("expected position to be %v, got %v", b, a)
	}
}

type mockImage struct {
	w, h float64
}

func newMockImage(w, h float64) *mockImage {
	m := &mockImage{}
	m.SetDimensions(w, h)

	return m
}

func (m *mockImage) DrawImage(_ graph.Image, _ graph.DrawOpt) {}

func (m *mockImage) Dimensions() (width, height float64) {
	return m.w, m.h // Example dimensions
}

func (m *mockImage) SetDimensions(width, height float64) {
	m.w = width
	m.h = height
}
