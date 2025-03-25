package tests

import (
	. "github.com/michaelzhao820/raytracer/raytracer"
	"testing"
)

func TestColorAddition(t *testing.T) {
	c1 := NewColor(0.9, 0.6, 0.75)
	c2 := NewColor(0.7, 0.1, 0.25)
	expected := NewColor(1.6, 0.7, 1.0)
	result := c1.AddColor(c2)

	if !result.Equals(expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestColorSubtraction(t *testing.T) {
	c1 := NewColor(0.9, 0.6, 0.75)
	c2 := NewColor(0.7, 0.1, 0.25)
	expected := NewColor(0.2, 0.5, 0.5)
	result := c1.SubtractColor(c2)

	if !result.Equals(expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestColorScalarMultiplication(t *testing.T) {
	c := NewColor(0.2, 0.3, 0.4)
	expected := NewColor(0.4, 0.6, 0.8)
	result := c.MultiplyByScalar(2)

	if !result.Equals(expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestColorMultiplication(t *testing.T) {
	c1 := NewColor(1, 0.2, 0.4)
	c2 := NewColor(0.9, 1, 0.1)
	expected := NewColor(0.9, 0.2, 0.04)
	result := c1.MultiplyOtherColor(c2)

	if !result.Equals(expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
