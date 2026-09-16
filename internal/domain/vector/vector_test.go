package vector_test

import (
	"math"
	"testing"

	"survivor/internal/domain/vector"
)

func TestVector2D_Rotate(t *testing.T) {
	tests := []struct {
		x, y    float64
		radians float64
		want    vector.Vector
	}{
		{0, 1, 0, vector.New(0, 1)},
		{0, 1, math.Pi / 2, vector.New(-1, 0)},
		{0, 1, -math.Pi / 2, vector.New(1, 0)},
		{0, 1, math.Pi, vector.New(0, -1)},
		{0, 1, -math.Pi, vector.New(0, -1)},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			v := vector.New(tt.x, tt.y)

			got := v.Rotate(tt.radians)

			if !got.IsEqual(tt.want) {
				t.Errorf("Rotate() = %v, want %v", got, tt.want)
			}
		})
	}
}
