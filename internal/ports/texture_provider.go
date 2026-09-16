package ports

import "survivor/internal/domain/graph"

type TextureProvider interface {
	LoadTexture(name string) (graph.Image, error)
}
