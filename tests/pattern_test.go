package tests

import (
	. "github.com/michaelzhao820/raytracer/raytracer"
	"testing"
)

func TestStripePattern(t *testing.T) {
	white := NewColor(1, 1, 1)
	black := NewColor(0, 0, 0)
	pattern := NewStripePattern(white, black)

	t.Run("stripe pattern is constant in y", func(t *testing.T) {
		if !pattern.PatternAt(NewPoint(0, 0, 0)).Equals(white) {
			t.Error("Expected white at (0, 0, 0)")
		}
		if !pattern.PatternAt(NewPoint(0, 1, 0)).Equals(white) {
			t.Error("Expected white at (0, 1, 0)")
		}
		if !pattern.PatternAt(NewPoint(0, 2, 0)).Equals(white) {
			t.Error("Expected white at (0, 2, 0)")
		}
	})

	t.Run("stripe pattern is constant in z", func(t *testing.T) {
		if !pattern.PatternAt(NewPoint(0, 0, 0)).Equals(white) {
			t.Error("Expected white at (0, 0, 0)")
		}
		if !pattern.PatternAt(NewPoint(0, 0, 1)).Equals(white) {
			t.Error("Expected white at (0, 0, 1)")
		}
		if !pattern.PatternAt(NewPoint(0, 0, 2)).Equals(white) {
			t.Error("Expected white at (0, 0, 2)")
		}
	})

	t.Run("stripe pattern alternates in x", func(t *testing.T) {
		if !pattern.PatternAt(NewPoint(0.0, 0, 0)).Equals(white) {
			t.Error("Expected white at (0.0, 0, 0)")
		}
		if !pattern.PatternAt(NewPoint(0.9, 0, 0)).Equals(white) {
			t.Error("Expected white at (0.9, 0, 0)")
		}
		if !pattern.PatternAt(NewPoint(1.0, 0, 0)).Equals(black) {
			t.Error("Expected black at (1.0, 0, 0)")
		}
		if !pattern.PatternAt(NewPoint(-0.1, 0, 0)).Equals(black) {
			t.Error("Expected black at (-0.1, 0, 0)")
		}
		if !pattern.PatternAt(NewPoint(-1.0, 0, 0)).Equals(black) {
			t.Error("Expected black at (-1.0, 0, 0)")
		}
		if !pattern.PatternAt(NewPoint(-1.1, 0, 0)).Equals(white) {
			t.Error("Expected white at (-1.1, 0, 0)")
		}
	})
}

func TestLightingWithPatternApplied(t *testing.T) {
	white := NewColor(1, 1, 1)
	black := NewColor(0, 0, 0)

	// Pattern: stripe alternating between white and black
	pattern := NewStripePattern(white, black)

	// Material with the pattern
	m := DefaultMaterial()
	m.Pattern = pattern
	m.SetAmbient(1)
	m.SetDiffuse(0)
	m.SetSpecular(0)

	// Shared vectors
	eyev := NewVector(0, 0, -1)
	normalv := NewVector(0, 0, -1)
	light := Light{
		Position:  NewPoint(0, 0, -10),
		Intensity: NewColor(1, 1, 1),
	}

	// When: lighting at two points
	c1 := Lighting(*m, NewSphere(), light, NewPoint(0.9, 0, 0), eyev, normalv, false)
	c2 := Lighting(*m, NewSphere(), light, NewPoint(1.1, 0, 0), eyev, normalv, false)

	if !c1.Equals(white) {
		t.Errorf("Expected c1 to be white, got %v", c1)
	}
	if !c2.Equals(black) {
		t.Errorf("Expected c2 to be black, got %v", c2)
	}
}
