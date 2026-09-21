// Package tile provides an effectively infinite tilemap built from small,
// generated chunks. A ChunkGenerator produces the tile grid for any chunk
// coordinate on demand (e.g. by repeating a fixed Pattern), and TileMap
// streams chunks in and out of memory as a Viewport (typically the camera)
// moves through the world.
package tile

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// Pattern is a fixed rectangular grid of tile indices, e.g. parsed from a
// map file. It is the building block used by generators such as
// PatternGenerator to produce chunks.
type Pattern struct {
	width, height int
	tiles         []int
}

func (p Pattern) Width() int  { return p.width }
func (p Pattern) Height() int { return p.height }

// At returns the tile index at the given local coordinate.
func (p Pattern) At(x, y int) int {
	return p.tiles[y*p.width+x]
}

// ParsePattern parses a text grid of whitespace-separated tile indices, one
// row per line. Empty lines and lines starting with '#' are ignored.
func ParsePattern(data []byte) (Pattern, error) {
	var rows [][]int

	width := -1

	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		fields := strings.Fields(line)

		row := make([]int, len(fields))
		for i, f := range fields {
			v, err := strconv.Atoi(f)
			if err != nil {
				return Pattern{}, fmt.Errorf("tile: parse pattern: invalid tile index %q: %w", f, err)
			}

			row[i] = v
		}

		if width == -1 {
			width = len(row)
		} else if len(row) != width {
			return Pattern{}, fmt.Errorf("tile: parse pattern: row width %d, want %d", len(row), width)
		}

		rows = append(rows, row)
	}

	if err := scanner.Err(); err != nil {
		return Pattern{}, fmt.Errorf("tile: parse pattern: %w", err)
	}

	if len(rows) == 0 {
		return Pattern{}, fmt.Errorf("tile: parse pattern: empty map")
	}

	tiles := make([]int, 0, len(rows)*width)
	for _, row := range rows {
		tiles = append(tiles, row...)
	}

	return Pattern{width: width, height: len(rows), tiles: tiles}, nil
}
