// Package world owns the scene graph and camera for the game world.
// It manages the scene graph and also provides a layer-based draw system.
// The Draw method handles drawing the scene graph by preparing and executing
// draw callbacks for all drawable nodes.
package world

import (
	"iter"
	"slices"
	"strconv"

	"survivor/internal/domain/camera"
	"survivor/internal/domain/graph"
	"survivor/internal/domain/node"
	"survivor/internal/domain/node2D"
	"survivor/internal/domain/transform"
)

type World struct {
	rootNode   node.Node
	layerRoots []node.Node
	camera     *camera.Camera
}

func New(camera *camera.Camera) *World {
	return &World{
		rootNode:   node2D.New("root"),
		layerRoots: make([]node.Node, 0),
		camera:     camera,
	}
}

// AddNode adds a node to the specified layer.
// If the layer does not exist, it will be created.
func (w *World) AddNode(n node.Node, layerIndex int) {
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

// Updatable is implemented by nodes that need per-frame logic (e.g. reading
// input and moving). World.Update walks the scene graph and calls Update on
// every node that implements it, mirroring how Draw treats graph.Drawable.
type Updatable interface {
	Update()
}

func (w *World) Update() {
	for _, layerRoot := range w.layerRoots {
		updateNode(layerRoot)
	}
}

// updateNode recursively updates a node's children before the node itself.
func updateNode(n node.Node) {
	for child := range n.GetChildren() {
		updateNode(child)
	}

	if u, ok := n.(Updatable); ok {
		u.Update()
	}
}

// Draw draws the world scene graph onto the target image.
// It prepares the draw callbacks for all drawable nodes and
// executes them in the correct order.
func (w *World) Draw(target graph.Image) {
	w.camera.Update() // TODO - Check if it should be called in Update() instead of Draw()

	cb := newCallbackStacks()
	for i, layerRoot := range w.layerRoots {
		prepareCallbacks(cb, w.camera, layerRoot, i, target)
	}

	cb.excecuteAll()
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
		pos := camera.GetPosition()
		// translate world coordinates to camera-relative coordinates
		tr.Translate(-pos.X(), -pos.Y())

		f := func() {
			drawable.Draw(target, graph.DrawOpt{Transform: tr})
		}
		cb.addCallback(layerIndex, f)
	}
}

// Sort the nodes by their layer in descending order.
func sortByLayer(children iter.Seq[node.Node]) iter.Seq[node.Node] {
	getLayer := func(n node.Node) int {
		if d, ok := n.(graph.Drawable); ok {
			return d.GetLayer()
		}

		return 0
	}

	ch := slices.Collect(children)
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
