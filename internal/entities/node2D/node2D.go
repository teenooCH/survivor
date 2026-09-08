package node2D

import (
	"iter"
	"slices"
	"sync/atomic"

	"github.com/teenooCH/survivor/internal/entities/node"
)

type Node2D struct {
	id       uint64
	name     string
	children []node.Node
	parent   node.Node
}

var globalNodeID atomic.Uint64

func New(name string) *Node2D {
	return &Node2D{id: globalNodeID.Add(1), name: name, children: make([]node.Node, 0)}
}

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
		n.children[i] = n.children[len(n.children)-1]
		n.children = n.children[:len(n.children)-1]

		node.AttachParent(nil)

		return true
	}

	return false
}

func (n *Node2D) MarkDirty() {
	for c := range n.GetChildren() {
		c.MarkDirty()
	}
}
