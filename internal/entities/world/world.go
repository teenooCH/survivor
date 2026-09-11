package world

import (
	"strconv"

	"github.com/teenooCH/survivor/internal/entities/camera"
	"github.com/teenooCH/survivor/internal/entities/graph"
	"github.com/teenooCH/survivor/internal/entities/node"
	"github.com/teenooCH/survivor/internal/entities/node2D"
)

type World struct {
	rootNode   node.Node
	layerRoots []node.Node
	camera     *camera.Camera

	callbacks *callbackStacks
}

func NewWorld(surface graph.Image) *World {
	return &World{
		rootNode:   node2D.New("root"),
		layerRoots: make([]node.Node, 0),
		camera:     camera.New(surface),
		callbacks:  newCallbackStacks(),
	}
}

func (w *World) AddNode(layerIndex int, n node.Node) {
	if layerIndex >= len(w.layerRoots) {
		w.addLayerRoots(layerIndex)
	}

	w.layerRoots[layerIndex].AddChild(n)
}

func (w *World) addLayerRoots(layerIndex int) {
	for layerIndex >= len(w.layerRoots) {
		root := node2D.New("layerRoot_" + strconv.Itoa(layerIndex))
		w.layerRoots = append(w.layerRoots, root)
		w.rootNode.AddChild(root)
	}
}

func (w *World) RemoveNode(n node.Node) bool {
	parent := n.GetParent()
	if parent == nil {
		return false
	}

	if !parent.RemoveChild(n) {
		return false
	}

	n.AttachParent(nil)

	return true
}
