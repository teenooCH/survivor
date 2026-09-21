package assets

import (
	"fmt"
	"io/fs"
)

// MapProvider reads raw map/pattern files (e.g. a tile index grid) from an
// embedded or on-disk filesystem.
type MapProvider struct {
	fsys fs.FS
}

func NewMapProvider(fsys fs.FS) *MapProvider {
	return &MapProvider{fsys: fsys}
}

func (p *MapProvider) LoadMap(name string) ([]byte, error) {
	data, err := fs.ReadFile(p.fsys, name)
	if err != nil {
		return nil, fmt.Errorf("failed to read map %s: %w", name, err)
	}

	return data, nil
}
