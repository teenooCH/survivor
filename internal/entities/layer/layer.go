package layer

import (
	"slices"

	"github.com/teenooCH/survivor/internal/entities/graph"
	"github.com/teenooCH/survivor/internal/entities/transform"
	"github.com/teenooCH/survivor/internal/pkg/stack"
)

// Layers manages draw order: lower layer index are drawn first (background).
// Within a layer, nodes are drawn in LIFO order (last pushed = drawn first).
type Layers struct {
	layers []*stack.Stack[func()]
}

func New(defaultLayerCount int) *Layers {
	l := &Layers{make([]*stack.Stack[func()], 0, defaultLayerCount)}
	l.ensureLayers(defaultLayerCount)

	return l
}

// ensureLayers ensures that the layers slice has at least idx+1 layers.
func (l *Layers) ensureLayers(idx int) {
	for idx >= len(l.layers) {
		l.layers = append(l.layers, stack.New[func()]())
	}
}

// AddNode adds a node to the specified layer index. If the layer does not exist, it is created.
func (l *Layers) AddNode(layerIndex int, node graph.Drawable,
	target graph.Image, op transform.Transform,
) {
	l.ensureLayers(layerIndex)

	f := func() {
		node.Draw(target, op)
	}
	l.layers[layerIndex].Push(f)
}

// DrawAll draws all layers in order, from lowest to highest index.
// Within a layer, nodes are drawn in LIFO order (last pushed = drawn first).
func (l *Layers) DrawAll() {
	for layer := range slices.Values(l.layers) {
		for !layer.IsEmpty() {
			if f, ok := layer.Pop(); ok {
				f()
			}
		}
	}
}
