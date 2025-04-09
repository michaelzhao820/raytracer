package tests

import (
	. "github.com/michaelzhao820/raytracer/raytracer"
	"math"
	"testing"
)

func TestSphereNormals(t *testing.T) {
	type testCase struct {
		name     string
		sphere   *Sphere
		point    Tuple
		expected Tuple
	}

	sqrt3over3 := math.Sqrt(3) / 3
	sqrt2over2 := math.Sqrt(2) / 2

	// Transformed spheres
	translated := NewSphere()
	tm, _ := TranslationMatrix(0, 1, 0)
	translated.SetTransform(tm)

	rotatedScaled := NewSphere()
	m, _ := ScalingMatrix(1, 0.5, 1)
	n, _ := RotationZMatrix(math.Pi / 5)
	m, _ = m.MultiplyMatrices(n)
	rotatedScaled.SetTransform(m)

	tests := []testCase{
		{
			name:     "Normal on x axis",
			sphere:   NewSphere(),
			point:    NewPoint(1, 0, 0),
			expected: NewVector(1, 0, 0),
		},
		{
			name:     "Normal on y axis",
			sphere:   NewSphere(),
			point:    NewPoint(0, 1, 0),
			expected: NewVector(0, 1, 0),
		},
		{
			name:     "Normal on z axis",
			sphere:   NewSphere(),
			point:    NewPoint(0, 0, 1),
			expected: NewVector(0, 0, 1),
		},
		{
			name:     "Normal on non-axial point",
			sphere:   NewSphere(),
			point:    NewPoint(sqrt3over3, sqrt3over3, sqrt3over3),
			expected: NewVector(sqrt3over3, sqrt3over3, sqrt3over3),
		},
		{
			name:     "Normal on translated sphere",
			sphere:   translated,
			point:    NewPoint(0, 1.70711, -0.70711),
			expected: NewVector(0, 0.70711, -0.70711),
		},
		{
			name:     "Normal on transformed sphere (scaled and rotated)",
			sphere:   rotatedScaled,
			point:    NewPoint(0, sqrt2over2, -sqrt2over2),
			expected: NewVector(0, 0.97014, -0.24254),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			n := tc.sphere.NormalAt(tc.point)
			if !n.Equals(tc.expected) {
				t.Errorf("Expected normal %v but got %v", tc.expected, n)
			}
		})
	}

	t.Run("Normal is normalized", func(t *testing.T) {
		s := NewSphere()
		point := NewPoint(sqrt3over3, sqrt3over3, sqrt3over3)
		n := s.NormalAt(point)
		normalized, _ := n.Normalize()
		if !n.Equals(normalized) {
			t.Errorf("Expected normal to be normalized, got %v", n)
		}
	})
}
