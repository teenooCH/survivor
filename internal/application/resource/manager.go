// Package resource provides a manager for loading and
// retrieving game resources such as textures and tile patterns.
package resource

import (
	"fmt"

	"survivor/internal/domain/graph"
	"survivor/internal/domain/tile"
	"survivor/internal/ports"
)

type Manager struct {
	textures        map[string]graph.Image
	patterns        map[string]tile.Pattern
	textureProvider ports.TextureProvider
	mapProvider     ports.MapProvider
}

func NewManager(textureProvider ports.TextureProvider, mapProvider ports.MapProvider) *Manager {
	return &Manager{
		textures:        make(map[string]graph.Image),
		patterns:        make(map[string]tile.Pattern),
		textureProvider: textureProvider,
		mapProvider:     mapProvider,
	}
}

// LoadTexture loads a texture from the specified path and
// associates it with the given name.
func (m *Manager) LoadTexture(name, path string) error {
	texture, err := m.textureProvider.LoadTexture(path)
	if err != nil {
		return err
	}

	m.textures[name] = texture

	return nil
}

// LoadTileset loads a spritesheet and slices it into count tiles of
// tileWidth x tileHeight pixels (see ports.TextureProvider.LoadTileset),
// storing each one under "namePrefix_<index>".
func (m *Manager) LoadTileset(namePrefix, path string, tileWidth, tileHeight, spacing, count int) error {
	textures, err := m.textureProvider.LoadTileset(path, tileWidth, tileHeight, spacing, count)
	if err != nil {
		return err
	}

	for i, texture := range textures {
		m.textures[fmt.Sprintf("%s_%d", namePrefix, i)] = texture
	}

	return nil
}

// LoadPattern loads a tile index grid from the specified path and
// associates it with the given name.
func (m *Manager) LoadPattern(name, path string) error {
	data, err := m.mapProvider.LoadMap(path)
	if err != nil {
		return err
	}

	pattern, err := tile.ParsePattern(data)
	if err != nil {
		return err
	}

	m.patterns[name] = pattern

	return nil
}

// GetPattern retrieves the pattern associated with the given name.
// Returns false if the pattern is not found.
func (m *Manager) GetPattern(name string) (tile.Pattern, bool) {
	p, ok := m.patterns[name]
	return p, ok
}

// GetTexture retrieves the texture associated with the given name.
// Returns false if the texture is not found.
func (m *Manager) GetTexture(name string) (graph.Image, bool) {
	img, ok := m.textures[name]
	return img, ok
}
