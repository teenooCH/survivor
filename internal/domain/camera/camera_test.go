package camera_test

import (
	"testing"

	"survivor/internal/domain/camera"
	"survivor/internal/domain/node2D"
	"survivor/internal/domain/vector"
)

func TestCamera_Update_FollowTheNode(t *testing.T) {
	t.Run("", func(t *testing.T) {
		nf := node2D.New("testNode")
		nf.SetPosition(1000, 750)

		c := camera.New(200, 100)
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
