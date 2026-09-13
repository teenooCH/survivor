package transform_test

import (
	"math"
	"testing"

	"survivor/internal/entities/transform"
	"survivor/internal/entities/vector"
)

func TestTransform_Concatenate(t *testing.T) {
	tests := []struct {
		name       string
		tr         transform.Transform
		trScale    vector.Vector
		other      transform.Transform
		otherScale vector.Vector
		want       transform.Transform
		wantScale  vector.Vector
	}{
		{
			name:  "All zero transforms",
			tr:    transform.NewZero(),
			other: transform.NewZero(),
			want:  transform.NewZero(),
		},
		{
			name:  "tr at 0/0, other at 1/1",
			tr:    transform.NewZero(),
			other: transform.New(vector.New(1, 1), vector.New(0, 0), 0),
			want:  transform.New(vector.New(1, 1), vector.New(0, 0), 0),
		},
		{
			name:      "tr scaled to 2/2",
			tr:        transform.NewZero(),
			trScale:   vector.New(2, 2),
			other:     transform.New(vector.New(1, 1), vector.New(0, 0), 0),
			want:      transform.New(vector.New(2, 2), vector.New(0, 0), 0),
			wantScale: vector.New(2, 2),
		},
		{
			name:  "rotate tr by 90 degrees",
			tr:    transform.New(vector.New(0, 0), vector.New(0, 0), math.Pi/2),
			other: transform.New(vector.New(1, 0), vector.New(0, 0), 0),
			want:  transform.New(vector.New(0, 1), vector.New(0, 0), math.Pi/2),
		},
		{
			name:  "rotate other by 90 degrees",
			tr:    transform.NewZero(),
			other: transform.New(vector.New(1, 0), vector.New(0, 0), math.Pi/2),
			want:  transform.New(vector.New(1, 0), vector.New(0, 0), math.Pi/2),
		},
		{
			name:  "tr at 0/0, other pivot at 1/1",
			tr:    transform.NewZero(),
			other: transform.New(vector.New(0, 0), vector.New(1, 1), 0),
			want:  transform.New(vector.New(0, 0), vector.New(1, 1), 0),
		},
		{
			name:      "scale and rotate tr, rotate other",
			tr:        transform.New(vector.New(0, 0), vector.New(0, 0), math.Pi/2),
			trScale:   vector.New(2, 2),
			other:     transform.New(vector.New(1, 0), vector.New(3, 3), math.Pi/2),
			want:      transform.New(vector.New(0, 2), vector.New(3, 3), math.Pi),
			wantScale: vector.New(2, 2),
		},
		{
			name:       "scale and rotate tr and other",
			tr:         transform.New(vector.New(0, 0), vector.New(0, 0), math.Pi/2),
			trScale:    vector.New(2, 2),
			other:      transform.New(vector.New(1, 0), vector.New(3, 3), math.Pi/2),
			otherScale: vector.New(4, 4),
			want:       transform.New(vector.New(0, 2), vector.New(3, 3), math.Pi),
			wantScale:  vector.New(8, 8),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setScale(&tt.tr, tt.trScale)
			setScale(&tt.other, tt.otherScale)
			setScale(&tt.want, tt.wantScale)

			ttr := tt.tr

			ttr.Concatenate(tt.other)

			if got := ttr; !compare(got, tt.want) {
				t.Errorf("Transform.Concatenate() =\ngot  %v\nwant %v", got, tt.want)
			}
		})
	}
}

// helper function to compare two Transform instances for equality.
func compare(a, b transform.Transform) bool {
	if !a.Position().IsEqual(b.Position()) {
		return false
	}

	if !a.Pivot().IsEqual(b.Pivot()) {
		return false
	}

	if a.Rotation() != b.Rotation() {
		return false
	}

	if !a.Scale().IsEqual(b.Scale()) {
		return false
	}

	return true
}

// setScale sets the scale of a Transform if the provided scale vector is not zero.
func setScale(t *transform.Transform, scale vector.Vector) {
	if scale.X() != 0 && scale.Y() != 0 {
		t.SetScale(scale.X(), scale.Y())
	}
}
