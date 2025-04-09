package tests

import (
	. "github.com/michaelzhao820/raytracer/raytracer"
	"math"
	"testing"
)

func TestLighting(t *testing.T) {
	m := *DefaultMaterial()
	position := NewPoint(0, 0, 0)

	t.Run("Lighting with the eye between the light and the surface", func(t *testing.T) {
		eyev := NewVector(0, 0, -1)
		normalv := NewVector(0, 0, -1)
		light := Light{Position: NewPoint(0, 0, -10), Intensity: NewColor(1, 1, 1)}
		result := Lighting(m, light, position, eyev, normalv)
		expected := NewColor(1.9, 1.9, 1.9)
		assertColorEqual(t, result, expected)
	})

	t.Run("Lighting with the eye between light and surface, eye offset 45°", func(t *testing.T) {
		eyev := NewVector(0, math.Sqrt2/2, -math.Sqrt2/2)
		normalv := NewVector(0, 0, -1)
		light := Light{Position: NewPoint(0, 0, -10), Intensity: NewColor(1, 1, 1)}
		result := Lighting(m, light, position, eyev, normalv)
		expected := NewColor(1.0, 1.0, 1.0)
		assertColorEqual(t, result, expected)
	})

	t.Run("Lighting with eye opposite surface, light offset 45°", func(t *testing.T) {
		eyev := NewVector(0, 0, -1)
		normalv := NewVector(0, 0, -1)
		light := Light{Position: NewPoint(0, 10, -10), Intensity: NewColor(1, 1, 1)}
		result := Lighting(m, light, position, eyev, normalv)
		intensity := 0.1 + 0.9*math.Sqrt2/2
		expected := NewColor(intensity, intensity, intensity)
		assertColorEqual(t, result, expected)
	})

	t.Run("Lighting with eye in the path of the reflection vector", func(t *testing.T) {
		eyev := NewVector(0, -math.Sqrt2/2, -math.Sqrt2/2)
		normalv := NewVector(0, 0, -1)
		light := Light{Position: NewPoint(0, 10, -10), Intensity: NewColor(1, 1, 1)}
		result := Lighting(m, light, position, eyev, normalv)
		intensity := 0.1 + 0.9*math.Sqrt2/2 + 0.9
		expected := NewColor(intensity, intensity, intensity)
		assertColorEqual(t, result, expected)
	})

	t.Run("Lighting with the light behind the surface", func(t *testing.T) {
		eyev := NewVector(0, 0, -1)
		normalv := NewVector(0, 0, -1)
		light := Light{Position: NewPoint(0, 0, 10), Intensity: NewColor(1, 1, 1)}
		result := Lighting(m, light, position, eyev, normalv)
		expected := NewColor(0.1, 0.1, 0.1)
		assertColorEqual(t, result, expected)
	})
}

// Helper function for comparing two colors with a small epsilon tolerance
func assertColorEqual(t *testing.T, got, want Color) {
	const epsilon = 1e-4
	if math.Abs(got.Tuple[R]-want.Tuple[R]) > epsilon ||
		math.Abs(got.Tuple[G]-want.Tuple[G]) > epsilon ||
		math.Abs(got.Tuple[B]-want.Tuple[B]) > epsilon {
		t.Errorf("Expected color %v, but got %v", want, got)
	}
}
