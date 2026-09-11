package camera_test

import (
	"testing"

	"github.com/teenooCH/survivor/internal/entities/camera"
	"github.com/teenooCH/survivor/internal/entities/graph"
	"github.com/teenooCH/survivor/internal/entities/node2D"
	"github.com/teenooCH/survivor/internal/entities/transform"
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

func (m *mockImage) DrawImage(image graph.Image, options transform.Transform) {}

func (m *mockImage) Dimensions() (width, height float64) {
	return m.w, m.h // Example dimensions
}

func (m *mockImage) SetDimensions(width, height float64) {
	m.w = width
	m.h = height
}
