package ports

// MapProvider loads raw map/pattern data (e.g. a tile index grid) by name.
type MapProvider interface {
	LoadMap(name string) ([]byte, error)
}
