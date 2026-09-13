// Package world owns the scene graph and camera for the game world.
// It manages the scene graph and also provides a layer-based rendering system.
// The Draw method handles rendering the world by preparing and executing
// draw callbacks for all drawable nodes.
package world

import (
	"iter"
	"slices"
	"strconv"

	"survivor/internal/entities/camera"
	"survivor/internal/entities/graph"
	"survivor/internal/entities/node"
	"survivor/internal/entities/node2D"
	"survivor/internal/entities/transform"
)

type World struct {
	rootNode   node.Node
	layerRoots []node.Node
	camera     *camera.Camera
}

func New(surface graph.Image) *World {
	return &World{
		rootNode:   node2D.New("root"),
		layerRoots: make([]node.Node, 0),
		camera:     camera.New(surface),
	}
}

// AddNode adds a node to the specified layer.
// If the layer does not exist, it will be created.
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

// RemoveNode removes the specified node from its parent.
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

func (w *World) Update() {
	// I don't know yet what this function should do
}

// Draw renders the world to the target image.
// It prepares the draw callbacks for all drawable nodes and
// executes them in the correct order.
func (w *World) Draw(target graph.Image) {
	w.camera.Update()

	cb := newCallbackStacks()
	for i, layerRoot := range w.layerRoots {
		prepareCallbacks(cb, w.camera, layerRoot, i, target)
	}

	cb.excecuteAll()
	w.camera.DrawToScreen(target)
}

// prepare the callbacks for the given node and its children recursively.
// The callbacks are added to the appropriate layer in the callback stack.
func prepareCallbacks(cb *callbackStacks, camera *camera.Camera,
	node node.Node, layerIndex int, target graph.Image,
) {
	for child := range sortByLayer(node.GetChildren()) {
		prepareCallbacks(cb, camera, child, layerIndex, target)
	}

	if drawable, ok := node.(graph.Drawable); ok {
		tr := prepareTransform(drawable)
		camera.ApplyOffset(&tr)
		op := graph.DrawOpt{Transform: tr}
		f := func() {
			drawable.Draw(target, op)
		}
		cb.addCallback(layerIndex, f)
	}
}

// Remove non-drawable nodes and sort the remaining nodes
// by their layer in descending order.
func sortByLayer(children iter.Seq[node.Node]) iter.Seq[node.Node] {
	getLayer := func(n node.Node) int {
		if d, ok := n.(graph.Drawable); ok {
			return d.GetLayer()
		}

		return 0
	}

	ch := slices.Collect(children)

	// filter out non-drawable nodes
	ch = slices.DeleteFunc(ch, func(n node.Node) bool {
		if _, ok := n.(graph.Drawable); !ok {
			return true
		}

		return false
	})

	slices.SortFunc(ch, func(a, b node.Node) int {
		return getLayer(b) - getLayer(a)
	})

	return slices.Values(ch)
}

// prepareTransform calculates and returns the
// transformation for the given drawable.
// Move pivot to the origin, apply scale and rotation,
// then move to the world position.
func prepareTransform(d graph.Drawable) transform.Transform {
	op := transform.NewZero()

	if tr, ok := d.(transform.Transformable); ok {
		wt := tr.GetWorldTransform()
		pivot := wt.Pivot()
		pos := wt.Position()
		scale := wt.Scale()
		rot := wt.Rotation()

		op.Translate(-pivot.X(), -pivot.Y())
		op.SetScale(scale.X(), scale.Y())
		op.Rotate(rot)
		op.Translate(pos.X(), pos.Y())
	}

	return op
}
