package resource

import (
	"survivor/internal/domain/graph"
	"survivor/internal/ports"
)

type Manager struct {
	textures map[string]graph.Image
	provider ports.TextureProvider
}

func NewManager(provider ports.TextureProvider) *Manager {
	return &Manager{
		textures: make(map[string]graph.Image),
		provider: provider,
	}
}

// LoadTexture loads a texture from the specified path and
// associates it with the given name.
func (m *Manager) LoadTexture(path, name string) error {
	texture, err := m.provider.LoadTexture(path)
	if err != nil {
		return err
	}

	m.textures[name] = texture

	return nil
}

// GetTexture retrieves the texture associated with the given name.
// Returns false if the texture is not found.
func (m *Manager) GetTexture(name string) (graph.Image, bool) {
	img, ok := m.textures[name]
	return img, ok
}
