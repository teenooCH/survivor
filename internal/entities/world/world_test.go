package world_test

import (
	"testing"

	"github.com/teenooCH/survivor/internal/entities/graph"
	"github.com/teenooCH/survivor/internal/entities/node2D"
	"github.com/teenooCH/survivor/internal/entities/transform"
	"github.com/teenooCH/survivor/internal/entities/world"
)

// fakeImage is a minimal graph.Image used as the draw target in tests.
type fakeImage struct{}

func (fakeImage) DrawImage(image graph.Image, options graph.DrawOpt) {}
func (fakeImage) Dimensions() (width, height float64)                { return 100, 100 }
func (fakeImage) SetDimensions(width, height float64)                {}

// drawableNode is a node2D.Node2D that also implements graph.Drawable.
// Each Draw call appends the node's name to the shared order slice.
type drawableNode struct {
	node2D.Node2D
	layer int
	order *[]string
}

func newDrawableNode(name string, layer int, order *[]string) *drawableNode {
	return &drawableNode{Node2D: *node2D.New(name), layer: layer, order: order}
}

func (d *drawableNode) GetLayer() int { return d.layer }

func (d *drawableNode) Draw(target graph.Image, options graph.DrawOpt) {
	*d.order = append(*d.order, d.GetName())
}

func TestWorldDrawOnlyDrawsDrawableNodesInLayerOrder(t *testing.T) {
	var order []string

	w := world.New(fakeImage{})

	// worldLayer 0: two drawable nodes with different GetLayer values plus a non-drawable node.
	nodeHigh := newDrawableNode("high", 5, &order)
	nodeLow := newDrawableNode("low", 1, &order)
	nonDrawable := node2D.New("plain")

	w.AddNode(0, nodeLow)
	w.AddNode(0, nodeHigh)
	w.AddNode(0, nonDrawable)

	// worldLayer 1: a single drawable node, must be drawn after everything in worldLayer 0.
	nodeTop := newDrawableNode("top", 0, &order)
	w.AddNode(1, nodeTop)

	w.Draw(fakeImage{})

	want := []string{"low", "high", "top"}
	if len(order) != len(want) {
		t.Fatalf("Draw() order = %v, want %v", order, want)
	}

	for i := range want {
		if order[i] != want[i] {
			t.Fatalf("Draw() order = %v, want %v", order, want)
		}
	}
}

func TestWorldDrawSkipsNonDrawableNodes(t *testing.T) {
	var order []string

	w := world.New(fakeImage{})

	nonDrawable := node2D.New("plain")
	w.AddNode(0, nonDrawable)

	w.Draw(fakeImage{})

	if len(order) != 0 {
		t.Fatalf("Draw() order = %v, want empty", order)
	}
}

var _ transform.Transformable = (*drawableNode)(nil)
