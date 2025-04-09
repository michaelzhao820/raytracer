package tests

import (
	"math"
	"testing"

	. "github.com/michaelzhao820/raytracer/raytracer"
)

func TestTupleW(t *testing.T) {
	tests := []struct {
		name     string
		tuple    Tuple
		isPoint  bool
		isVector bool
	}{
		{
			name:     "w = 1.0 is a point",
			tuple:    NewTuple(4.3, -4.2, 3.1, 1.0),
			isPoint:  true,
			isVector: false,
		},
		{
			name:     "w = 0.0 is a vector",
			tuple:    NewTuple(4.3, -4.2, 3.1, 0.0),
			isPoint:  false,
			isVector: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.tuple.IsPoint() != tt.isPoint {
				t.Errorf("Expected IsPoint() to be %v, got %v", tt.isPoint, tt.tuple.IsPoint())
			}
			if tt.tuple.IsVector() != tt.isVector {
				t.Errorf("Expected IsVector() to be %v, got %v", tt.isVector, tt.tuple.IsVector())
			}
		})
	}
}

func TestConstructors(t *testing.T) {
	tests := []struct {
		name        string
		constructor func() Tuple
		expected    Tuple
	}{
		{
			name:        "Point() constructor creates a point",
			constructor: func() Tuple { return NewPoint(4, -4, 3) },
			expected:    NewTuple(4, -4, 3, 1),
		},
		{
			name:        "Vector() constructor creates a vector",
			constructor: func() Tuple { return NewVector(4, -4, 3) },
			expected:    NewTuple(4, -4, 3, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.constructor()
			if !result.Equals(tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestAddTuples(t *testing.T) {
	tests := []struct {
		name      string
		tuple     Tuple
		tuple2    Tuple
		expected  Tuple
		expectErr bool
	}{
		{
			name:      "Add point with a vector",
			tuple:     NewTuple(3, -2, 5, 1),
			tuple2:    NewTuple(-2, 3, 1, 0),
			expected:  NewTuple(1, 1, 6, 1),
			expectErr: false,
		},
		{
			name:      "Should error when trying to add two points",
			tuple:     NewTuple(3, -2, 5, 1),
			tuple2:    NewTuple(-2, 3, 1, 1),
			expected:  Tuple{},
			expectErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.tuple.Add(tt.tuple2)

			if tt.expectErr {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !result.Equals(tt.expected) {
				t.Errorf("expected tuple %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestSubtractTuples(t *testing.T) {
	tests := []struct {
		name      string
		tuple     Tuple
		tuple2    Tuple
		expected  Tuple
		expectErr bool
	}{
		{
			name:      "Subtract point with a point",
			tuple:     NewPoint(3, 2, 1),
			tuple2:    NewPoint(5, 6, 7),
			expected:  NewVector(-2, -4, -6),
			expectErr: false,
		},
		{
			name:      "Subtracting vector from a point",
			tuple:     NewPoint(3, 2, 1),
			tuple2:    NewVector(5, 6, 7),
			expected:  NewPoint(-2, -4, -6),
			expectErr: false,
		},
		{
			name:      "Subtracting two vectors",
			tuple:     NewVector(3, 2, 1),
			tuple2:    NewVector(5, 6, 7),
			expected:  NewVector(-2, -4, -6),
			expectErr: false,
		},
		{
			name:      "Should error when trying to subtract a point from a vector",
			tuple:     NewVector(3, 2, 1),
			tuple2:    NewPoint(-2, -4, -6),
			expected:  Tuple{},
			expectErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.tuple.Subtract(tt.tuple2)

			if tt.expectErr {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !result.Equals(tt.expected) {
				t.Errorf("expected tuple %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestMultiplyTuples(t *testing.T) {
	tests := []struct {
		name     string
		tuple    Tuple
		scalar   float64
		expected Tuple
	}{
		{
			name:     "multiplying a tuple by a scalar",
			tuple:    NewTuple(1, -2, 3, -4),
			scalar:   3.5,
			expected: NewTuple(3.5, -7, 10.5, -14),
		},
		{
			name:     "multiplying a tuple by a fraction",
			tuple:    NewTuple(1, -2, 3, -4),
			scalar:   0.5,
			expected: NewTuple(0.5, -1, 1.5, -2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.tuple.Multiply(tt.scalar)
			if !result.Equals(tt.expected) {
				t.Errorf("expected tuple %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestDivideTuples(t *testing.T) {
	tests := []struct {
		name     string
		tuple    Tuple
		scalar   float64
		expected Tuple
	}{
		{
			name:     "Dividing a tuple by a scalar",
			tuple:    NewTuple(1, -2, 3, -4),
			scalar:   2,
			expected: NewTuple(0.5, -1, 1.5, -2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.tuple.Divide(tt.scalar)
			if !result.Equals(tt.expected) {
				t.Errorf("expected tuple %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestMagnitude(t *testing.T) {
	tests := []struct {
		name      string
		tuple     Tuple
		expected  float64
		expectErr bool
	}{
		{
			name:      "Computing the magnitude of vector(0, 1, 0)",
			tuple:     NewVector(0, 1, 0),
			expected:  1,
			expectErr: false,
		},
		{
			name:      "Computing the magnitude of vector(0, 0, 1)",
			tuple:     NewVector(0, 0, 1),
			expected:  1,
			expectErr: false,
		},
		{
			name:      "Computing the magnitude of vector(1, 2, 3)",
			tuple:     NewVector(1, 2, 3),
			expected:  math.Sqrt(14),
			expectErr: false,
		},
		{
			name:      "Computing the magnitude of vector(-1, -2, -3)",
			tuple:     NewVector(-1, -2, -3),
			expected:  math.Sqrt(14),
			expectErr: false,
		},
		{
			name:      "Attempting to compute the magnitude of a point(1, 2, 3)",
			tuple:     NewPoint(1, 2, 3),
			expected:  0,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			magnitude, err := tt.tuple.Magnitude()

			if tt.expectErr {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Magnitude() error = %v, wantErr %v", err, nil)
				}
				if math.Abs(magnitude-tt.expected) > 1e-5 {
					t.Errorf("Magnitude() = %v, want %v", magnitude, tt.expected)
				}
			}
		})
	}
}

func TestNormalize(t *testing.T) {
	tests := []struct {
		name      string
		tuple     Tuple
		expected  Tuple
		expectErr bool
	}{
		{
			name:      "Normalizing vector(4, 0, 0) gives (1, 0, 0)",
			tuple:     NewVector(4, 0, 0),
			expected:  NewVector(1, 0, 0),
			expectErr: false,
		},
		{
			name:      "Normalizing vector(1, 2, 3) gives approximately (0.26726, 0.53452, 0.80178)",
			tuple:     NewVector(1, 2, 3),
			expected:  NewVector(0.26726, 0.53452, 0.80178),
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			normalized, err := tt.tuple.Normalize()

			if tt.expectErr {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Normalize() error = %v, wantErr %v", err, nil)
				}

				// Check that the normalized tuple is approximately equal to the expected value
				if !normalized.Equals(tt.expected) {
					t.Errorf("Normalize() = %v, want %v", normalized, tt.expected)
				}
			}
		})
	}
}

func TestMagnitudeOfNormalizedVector(t *testing.T) {
	tuple := NewVector(1, 2, 3)

	// Normalize the vector
	normalized, err := tuple.Normalize()
	if err != nil {
		t.Fatalf("Error normalizing vector: %v", err)
	}

	// Calculate the magnitude of the normalized vector
	magnitude, err := normalized.Magnitude()
	if err != nil {
		t.Fatalf("Error calculating magnitude: %v", err)
	}

	// The magnitude of a normalized vector should be 1
	if math.Abs(magnitude-1) > 1e-5 {
		t.Errorf("Magnitude of normalized vector = %v, want 1", magnitude)
	}
}

func TestDot(t *testing.T) {
	// Test the dot product of two vectors
	a := NewVector(1, 2, 3)
	b := NewVector(2, 3, 4)

	result, err := Dot(a, b)
	if err != nil {
		t.Errorf("Dot product calculation failed: %v", err)
	}
	if result != 20.0 {
		t.Errorf("Expected dot product to be 20, but got %v", result)
	}

	// Test when one of the tuples is a point (invalid case)
	p := NewPoint(1, 2, 3)
	_, err = Dot(a, p)
	if err == nil {
		t.Error("Expected error when computing dot product with a point, but got nil")
	}
}

func TestCross(t *testing.T) {
	// Test the cross product of two vectors
	a := NewVector(1, 2, 3)
	b := NewVector(2, 3, 4)

	result, err := Cross(a, b)
	if err != nil {
		t.Errorf("Cross product calculation failed: %v", err)
	}
	expected := NewVector(-1, 2, -1)
	if !result.Equals(expected) {
		t.Errorf("Expected cross product to be vector(-1, 2, -1), but got %v", result)
	}

	// Test the cross product when order is reversed
	result, err = Cross(b, a)
	if err != nil {
		t.Errorf("Cross product calculation failed: %v", err)
	}
	expected = NewVector(1, -2, 1)
	if !result.Equals(expected) {
		t.Errorf("Expected cross product to be vector(1, -2, 1), but got %v", result)
	}

	// Test when one of the tuples is a point (invalid case)
	p := NewPoint(1, 2, 3)
	_, err = Cross(a, p)
	if err == nil {
		t.Error("Expected error when computing cross product with a point, but got nil")
	}
}

func TestReflect(t *testing.T) {
	t.Run("Reflecting a vector approaching at 45°", func(t *testing.T) {
		v := NewVector(1, -1, 0)
		n := NewVector(0, 1, 0)
		r, _ := Reflect(v, n)
		expected := NewVector(1, 1, 0)

		if !r.Equals(expected) {
			t.Errorf("Expected %v but got %v", expected, r)
		}
	})

	t.Run("Reflecting a vector off a slanted surface", func(t *testing.T) {
		v := NewVector(0, -1, 0)
		n := NewVector(math.Sqrt(2)/2, math.Sqrt(2)/2, 0)
		r, _ := Reflect(v, n)
		expected := NewVector(1, 0, 0)

		if !r.Equals(expected) {
			t.Errorf("Expected %v but got %v", expected, r)
		}
	})
}
