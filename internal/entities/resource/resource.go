package resource

import "survivor/internal/entities/graph"

type Manager interface {
	// LoadTexture loads a texture from the specified path and
	// associates it with the given name.
	LoadTexture(path, name string) error

	// GetTexture retrieves the texture associated with the given name.
	// Returns false if the texture is not found.
	GetTexture(name string) (graph.Image, bool)
}
