package world

import (
	"slices"

	"survivor/internal/pkg/stack"
)

// callbackStacks is a helper structure that manages the draw callbacks for different layers.
// See executeAll() for the order in which callbacks are executed.
type callbackStacks struct {
	stacks []*stack.Stack[func()]
}

func newCallbackStacks() *callbackStacks {
	return &callbackStacks{make([]*stack.Stack[func()], 0)}
}

// addCallback adds a callback to the specified layer index.
// If the layer does not exist, it is created.
func (cb *callbackStacks) addCallback(layerIndex int, f func()) {
	if layerIndex >= len(cb.stacks) {
		cb.addLayers(layerIndex)
	}

	cb.stacks[layerIndex].Push(f)
}

// excecuteAll executes all callbacks in order, from lowest to highest layer index.
// Within a layer, functions are executed in LIFO order (last pushed = executed first).
// NB: After excecuteAll was called, all stacks are empty.
func (cb *callbackStacks) excecuteAll() {
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
