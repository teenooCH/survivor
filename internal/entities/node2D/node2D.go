// package node2D provides a basic node.Node implementation for a scene graph.
package node2D

import (
	"iter"
	"slices"
	"sync/atomic"

	"survivor/internal/entities/node"
	"survivor/internal/entities/transform"
	"survivor/internal/entities/vector"
)

// Node2D implements the node.Node and the transform.Transformable interface.
type Node2D struct {
	id             uint64
	name           string
	children       []node.Node
	parent         node.Node
	localTransform transform.Transform
	worldTransform transform.Transform
	isDirty        bool
}

var globalNodeID atomic.Uint64

// New creates a new Node2D instance with the given name and a unique ID.
func New(name string) *Node2D {
	return &Node2D{
		id:             globalNodeID.Add(1),
		name:           name,
		children:       make([]node.Node, 0),
		localTransform: transform.NewZero(),
		worldTransform: transform.NewZero(),
		isDirty:        true,
	}
}

// Implementation of the node.Node interface

func (n *Node2D) GetID() uint64               { return n.id }
func (n *Node2D) GetName() string             { return n.name }
func (n *Node2D) GetParent() node.Node        { return n.parent }
func (n *Node2D) AttachParent(node node.Node) { n.parent = node }

func (n *Node2D) AddChild(child node.Node) {
	child.AttachParent(n)
	n.children = append(n.children, child)
}

func (n *Node2D) GetChildren() iter.Seq[node.Node] {
	return func(yield func(node.Node) bool) {
		for _, child := range n.children {
			if !yield(child) {
				return
			}
		}
	}
}

func (n *Node2D) RemoveChild(node node.Node) bool {
	i := slices.Index(n.children, node)
	if i != -1 {
		n.children = slices.Delete(n.children, i, i+1)

		node.AttachParent(nil)

		return true
	}

	return false
}

func (n *Node2D) MarkDirty() {
	if n.isDirty {
		return
	}

	n.isDirty = true
	for c := range n.GetChildren() {
		c.MarkDirty()
	}
}

// Implementation of the transform.Transformable interface

func (n *Node2D) GetTransform() transform.Transform {
	return n.localTransform
}

func (n *Node2D) SetTransform(t transform.Transform) {
	n.localTransform = t
	n.MarkDirty()
}

// GetWorldTransform the concatenated transform of the node and all
// its parents up to the root of the scene graph. It is cached and
// only recalculated when the node or any of its parents are marked dirty.
func (n *Node2D) GetWorldTransform() transform.Transform {
	if !n.isDirty {
		return n.worldTransform
	}

	world := transform.NewZero()

	if n.parent != nil {
		if pt, ok := n.parent.(transform.Transformable); ok {
			world = pt.GetWorldTransform()
		}
	}

	world.Concatenate(n.localTransform)
	n.worldTransform = world
	n.isDirty = false

	return n.worldTransform
}

// Getters and setters for position, pivot, rotation, and scale

func (n *Node2D) SetPosition(x, y float64)   { n.localTransform.SetPosition(x, y); n.MarkDirty() }
func (n *Node2D) GetPosition() vector.Vector { return n.localTransform.Position() }
func (n *Node2D) SetRotation(r float64)      { n.localTransform.SetRotation(r); n.MarkDirty() }
func (n *Node2D) GetRotation() float64       { return n.localTransform.Rotation() }
func (n *Node2D) SetScale(x, y float64)      { n.localTransform.SetScale(x, y); n.MarkDirty() }
func (n *Node2D) GetScale() vector.Vector    { return n.localTransform.Scale() }
func (n *Node2D) GetPivot() vector.Vector    { return n.localTransform.Pivot() }
func (n *Node2D) SetPivot(x, y float64)      { n.localTransform.SetPivot(x, y); n.MarkDirty() }

// GetWorldPosition returns the world-space position (convenience for collision etc.).
func (n *Node2D) GetWorldPosition() vector.Vector {
	wt := n.GetWorldTransform()
	return wt.Position()
}
