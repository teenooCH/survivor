package tile_test

import (
	"testing"

	"survivor/internal/domain/tile"
)

func TestParsePattern(t *testing.T) {
	data := []byte(`
# comment line, should be ignored
0 1 2
1 0 1
`)

	p, err := tile.ParsePattern(data)
	if err != nil {
		t.Fatalf("ParsePattern() error = %v", err)
	}

	if p.Width() != 3 || p.Height() != 2 {
		t.Fatalf("ParsePattern() size = %dx%d, want 3x2", p.Width(), p.Height())
	}

	want := [2][3]int{{0, 1, 2}, {1, 0, 1}}

	for y := range 2 {
		for x := range 3 {
			if got := p.At(x, y); got != want[y][x] {
				t.Errorf("At(%d,%d) = %d, want %d", x, y, got, want[y][x])
			}
		}
	}
}

func TestParsePatternInconsistentRowWidth(t *testing.T) {
	_, err := tile.ParsePattern([]byte("0 1 2\n0 1\n"))
	if err == nil {
		t.Fatal("ParsePattern() error = nil, want error for inconsistent row width")
	}
}

func TestParsePatternEmpty(t *testing.T) {
	_, err := tile.ParsePattern([]byte("# only a comment\n"))
	if err == nil {
		t.Fatal("ParsePattern() error = nil, want error for empty map")
	}
}

func TestParsePatternInvalidIndex(t *testing.T) {
	_, err := tile.ParsePattern([]byte("0 x 2\n"))
	if err == nil {
		t.Fatal("ParsePattern() error = nil, want error for invalid tile index")
	}
}
