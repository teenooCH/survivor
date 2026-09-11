package world

import (
	"slices"

	"github.com/teenooCH/survivor/internal/entities/graph"
	"github.com/teenooCH/survivor/internal/pkg/stack"
)

// callbackStacks manages draw order: lower layer index are drawn first (background).
// Within a layer, nodes are drawn in LIFO order (last pushed = drawn first).
type callbackStacks struct {
	stacks []*stack.Stack[func()]
}

func newCallbackStack() *callbackStacks {
	l := &callbackStacks{make([]*stack.Stack[func()], 0)}

	return l
}

// addCallback adds a callback to the specified layer index.
// The callback will draw the node on the target image
// with the specified draw options.
// If the layer does not exist, it is created.
func (cb *callbackStacks) addCallback(layerIndex int, node graph.Drawable,
	target graph.Image, op graph.DrawOpt,
) {
	if layerIndex >= len(cb.stacks) {
		cb.addLayers(layerIndex)
	}

	f := func() {
		node.Draw(target, op)
	}
	cb.stacks[layerIndex].Push(f)
}

// drawAll executes all callbacks in order, from lowest to highest layer index.
// Within a layer, nodes are drawn in LIFO order (last pushed = drawn first).
// NB: After drawAll is called, all callbacks are removed from the stacks.
func (cb *callbackStacks) drawAll() {
	for layer := range slices.Values(cb.stacks) {
		for !layer.IsEmpty() {
			if f, ok := layer.Pop(); ok {
				f()
			}
		}
	}
}

// addLayers ensures that the callback slice has at least idx+1 layers.
func (cb *callbackStacks) addLayers(idx int) {
	for idx >= len(cb.stacks) {
		cb.stacks = append(cb.stacks, stack.New[func()]())
	}
}
