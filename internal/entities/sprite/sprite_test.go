package sprite_test

import (
	"testing"

	"github.com/teenooCH/survivor/internal/entities/graph"
	"github.com/teenooCH/survivor/internal/entities/sprite"
)

func TestNew(t *testing.T) {
	tests := []struct {
		width, height float64
		wantX, wantY  float64
	}{
		{width: 100, height: 100, wantX: 50, wantY: 50},
		{width: 100, height: 10, wantX: 50, wantY: 5},
		{width: 99, height: 9, wantX: 49.5, wantY: 4.5},
		{width: 98, height: 9, wantX: 49, wantY: 4.5},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := sprite.New("test sprite", newMockImage(tt.width, tt.height), 0)

			gotPosition := got.GetPivot()

			if gotPosition.X() != tt.wantX || gotPosition.Y() != tt.wantY {
				t.Errorf("New(): Got pivot position %v, want {%v %v}", got.GetPivot(), tt.wantX, tt.wantY)
			}
		})
	}
}

type mockImage struct {
	w, h float64
}

func newMockImage(w, h float64) *mockImage {
	m := &mockImage{}
	m.SetDimensions(w, h)

	return m
}

func (m *mockImage) DrawImage(_ graph.Image, _ graph.DrawOpt) {}

func (m *mockImage) Dimensions() (width, height float64) {
	return m.w, m.h // Example dimensions
}

func (m *mockImage) SetDimensions(width, height float64) {
	m.w = width
	m.h = height
}
