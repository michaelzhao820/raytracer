package tests

import (
	. "github.com/michaelzhao820/raytracer/raytracer"
	"testing"
)

func TestViewTransform(t *testing.T) {

	t.Run("Default orientation returns identity matrix", func(t *testing.T) {
		from := NewPoint(0, 0, 0)
		to := NewPoint(0, 0, -1)
		up := NewVector(0, 1, 0)

		result := ViewTransform(from, to, up)

		expected := IdentityMatrix()
		if !result.Equals(expected) {
			t.Errorf("Expected identity matrix, got %v", result)
		}
	})

	t.Run("Looking in positive Z direction reflects space", func(t *testing.T) {
		from := NewPoint(0, 0, 0)
		to := NewPoint(0, 0, 1)
		up := NewVector(0, 1, 0)

		result := ViewTransform(from, to, up)
		expected, _ := ScalingMatrix(-1, 1, -1)

		if !result.Equals(expected) {
			t.Errorf("Expected scaling(-1,1,-1), got %v", result)
		}
	})

	t.Run("View transform moves the world", func(t *testing.T) {
		from := NewPoint(0, 0, 8)
		to := NewPoint(0, 0, 0)
		up := NewVector(0, 1, 0)

		result := ViewTransform(from, to, up)
		expected, _ := TranslationMatrix(0, 0, -8)

		if !result.Equals(expected) {
			t.Errorf("Expected translation(0,0,-8), got %v", result)
		}
	})

	t.Run("Arbitrary view transform", func(t *testing.T) {
		from := NewPoint(1, 3, 2)
		to := NewPoint(4, -2, 8)
		up := NewVector(1, 1, 0)

		result := ViewTransform(from, to, up)

		expected := NewMatrix(4, 4)
		expected, _ = expected.Set(0, 0, -0.50709)
		expected, _ = expected.Set(0, 1, 0.50709)
		expected, _ = expected.Set(0, 2, 0.67612)
		expected, _ = expected.Set(0, 3, -2.36643)

		expected, _ = expected.Set(1, 0, 0.76772)
		expected, _ = expected.Set(1, 1, 0.60609)
		expected, _ = expected.Set(1, 2, 0.12122)
		expected, _ = expected.Set(1, 3, -2.82843)

		expected, _ = expected.Set(2, 0, -0.35857)
		expected, _ = expected.Set(2, 1, 0.59761)
		expected, _ = expected.Set(2, 2, -0.71714)
		expected, _ = expected.Set(2, 3, 0.0)

		expected, _ = expected.Set(3, 0, 0.0)
		expected, _ = expected.Set(3, 1, 0.0)
		expected, _ = expected.Set(3, 2, 0.0)
		expected, _ = expected.Set(3, 3, 1.0)

		if !result.Equals(expected) {
			t.Errorf("Expected view matrix = %v, got %v", expected, result)
		}
	})
}
