package node2D_test

import (
	"testing"

	"survivor/internal/entities/node2D"
	"survivor/internal/entities/transform"
	"survivor/internal/entities/vector"
)

func TestNode2D_GetWorldTransform(t *testing.T) {
	type testNode struct {
		name string
		x, y float64 // local position
	}

	tests := []struct {
		name    string     // description of this test case
		parents []testNode // parent nodes
		want    transform.Transform
	}{
		{"No parents", []testNode{}, transform.New(vector.New(1, 1), vector.New(0, 0), 0)},
		{"Single parent", []testNode{{name: "parent1", x: 2, y: 2}}, transform.New(vector.New(3, 3), vector.New(0, 0), 0)},
		{"Two parents", []testNode{
			{name: "parent1", x: 2, y: 2},
			{name: "parent2", x: 3, y: 3},
		}, transform.New(vector.New(6, 6), vector.New(0, 0), 0)},
		{"Five parents", []testNode{
			{name: "parent1", x: 2, y: 2},
			{name: "parent2", x: 3, y: 3},
			{name: "parent3", x: 4, y: 4},
			{name: "parent4", x: 5, y: 5},
			{name: "parent5", x: 6, y: 6},
		}, transform.New(vector.New(21, 21), vector.New(0, 0), 0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := node2D.New("start")
			start.SetPosition(1, 1)

			current := start

			for _, p := range tt.parents {
				parent := node2D.New(p.name)
				parent.SetPosition(p.x, p.y)
				parent.AddChild(current)
				current = parent
			}

			got := start.GetWorldTransform()

			if got != tt.want {
				t.Errorf("GetWorldTransform()\ngot  %v\nwant %v", got, tt.want)
			}
		})
	}
}
