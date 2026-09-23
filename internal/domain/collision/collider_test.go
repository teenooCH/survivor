package collision_test

import (
	"testing"

	"survivor/internal/domain/collision"
	"survivor/internal/domain/node2D"
)

func TestCollider_OverlapsWith(t *testing.T) {
	type pos struct {
		x, y float64
	}

	first := collision.NewCollider(
		"first",
		collision.NewMask(collision.LayerPlayer, collision.LayerEnemy),
		collision.NewCircle(10),
	)
	second := collision.NewCollider(
		"second",
		collision.NewMask(collision.LayerEnemy, collision.LayerPlayer),
		collision.NewCircle(15),
	)

	firstParent := node2D.New("first-parent")
	firstParent.AddChild(first)

	secondParent := node2D.New("second-parent")
	secondParent.AddChild(second)

	tests := []struct {
		first  pos
		second pos
		want   bool
	}{
		{
			first:  pos{x: 0, y: 0},
			second: pos{x: 10, y: 0},
			want:   true,
		},
		{
			first:  pos{x: 0, y: 0},
			second: pos{x: 30, y: 0},
			want:   false,
		},
		{
			first:  pos{x: 100, y: 200},
			second: pos{x: 110, y: 215},
			want:   true,
		},
		{
			first:  pos{x: 100, y: 200},
			second: pos{x: 130, y: 215},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			firstParent.SetPosition(tt.first.x, tt.first.y)
			secondParent.SetPosition(tt.second.x, tt.second.y)

			got := first.OverlapsWith(second)
			if got != tt.want {
				t.Errorf("OverlapsWith() = %v, want %v", got, tt.want)
			}
		})
	}
}
