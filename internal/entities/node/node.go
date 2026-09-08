package node

import "iter"

type Node interface {
	// GetID returns the unique identifier of the node.
	GetID() uint64

	// GetName returns the name of the node.
	GetName() string

	// GetType returns the type of the node.
	// GetType() string

	// GetChildren returns the child nodes of the current node.
	GetChildren() iter.Seq[Node]

	// AddChild adds a child node to the current node.
	AddChild(child Node)

	// RemoveChild removes a child node from the current node.
	RemoveChild(child Node) bool

	// AttachParent attaches a parent node to the current node.
	AttachParent(parent Node)

	// GetParent returns the parent node of the current node.
	GetParent() Node

	// MarkDirty marks the node as dirty, indicating that it has been modified.
	MarkDirty()
}
