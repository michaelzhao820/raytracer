package raytracer

import (
	"fmt"
	"math"
)

type ColorIndex int

const (
	R ColorIndex = iota
	G
	B
)

type PositionIndex int

const (
	X PositionIndex = iota
	Y
	Z
	W
)

type Tuple []float64

func NewTuple(values ...float64) Tuple {
	return values
}

func NewPoint(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1.0}
}

func NewVector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0.0}
}

func (t Tuple) IsPoint() bool {
	if len(t) <= 3 {
		return false
	}
	return t[W] == 1.0
}

func (t Tuple) IsVector() bool {
	if len(t) <= 3 {
		return false
	}
	return t[W] == 0.0
}

func (t Tuple) Equals(other Tuple) bool {
	for i := range t {
		if !equalsWithMargin(t[i], other[i]) {
			return false
		}
	}
	return true
}

func (t Tuple) Add(other Tuple) (Tuple, error) {
	if t.IsPoint() && other.IsPoint() {
		return Tuple{}, fmt.Errorf("cannot add two points (w = 1) together")
	}
	if len(t) != len(other) {
		return Tuple{}, fmt.Errorf("cannot add tuples of different sizes")
	}
	rt := make([]float64, len(t))
	for i := range t {
		rt[i] = t[i] + other[i]
	}
	return rt, nil
}

func (t Tuple) Subtract(other Tuple) (Tuple, error) {
	if t.IsVector() && other.IsPoint() {
		return Tuple{}, fmt.Errorf("cannot subtract a point from a vector")
	}
	if len(t) != len(other) {
		return Tuple{}, fmt.Errorf("cannot add tuples of different sizes")
	}

	rt := make([]float64, len(t))
	for i := range t {
		rt[i] = t[i] - other[i]
	}
	return rt, nil
}

func (t Tuple) Multiply(scalar float64) Tuple {
	rt := make(Tuple, len(t))
	for i := range t {
		rt[i] = t[i] * scalar
	}
	return rt
}

func (t Tuple) Divide(scalar float64) Tuple {
	rt := make(Tuple, len(t))
	for i := range t {
		rt[i] = t[i] / scalar
	}
	return rt
}

func (t Tuple) Magnitude() (float64, error) {
	if !t.IsVector() {
		return 0, fmt.Errorf("magnitude only defined for vectors")
	}

	var sum float64
	sum += t[X] * t[X]
	sum += t[Y] * t[Y]
	sum += t[Z] * t[Z]
	return math.Sqrt(sum), nil
}

func (t Tuple) Normalize() (Tuple, error) {
	if !t.IsVector() {
		return nil, fmt.Errorf("normalization only defined for vectors")
	}

	mag, err := t.Magnitude()
	if err != nil {
		return nil, err
	}

	return t.Divide(mag), nil
}

func Dot(t, other Tuple) (float64, error) {
	if len(t) != len(other) {
		return 0, fmt.Errorf("tuples must be same length")
	}

	if t.IsPoint() || other.IsPoint() {
		return 0, fmt.Errorf("Both tuples must be vectors")
	}

	var sum float64
	sum += t[X] * other[X]
	sum += t[Y] * other[Y]
	sum += t[Z] * other[Z]
	return sum, nil
}

// Cross product (for 3D vectors)
func Cross(a, b Tuple) (Tuple, error) {
	if !a.IsVector() || !b.IsVector() {
		return nil, fmt.Errorf("cross product only defined for vectors")
	}

	return NewVector(
		a[Y]*b[Z]-a[Z]*b[Y],
		a[Z]*b[X]-a[X]*b[Z],
		a[X]*b[Y]-a[Y]*b[X],
	), nil
}

func equalsWithMargin(a, b float64) bool {
	epsilon := 0.00001
	if math.Abs(a-b) < epsilon {
		return true
	}
	return false
}

type Color struct {
	Tuple
}

func NewColor(r, g, b float64) Color {
	return Color{NewTuple(r, g, b)}
}

func (c Color) AddColor(other Color) Color {
	t, _ := c.Add(other.Tuple)
	return Color{t}
}

func (c Color) SubtractColor(other Color) Color {
	t, _ := c.Subtract(other.Tuple)
	return Color{t}
}

func (c Color) MultiplyByScalar(scalar float64) Color {
	return Color{c.Multiply(scalar)}
}

func (c Color) MultiplyOtherColor(other Color) Color {
	r, g, b := c.Tuple[R]*other.Tuple[R], c.Tuple[G]*other.Tuple[G], c.Tuple[B]*other.Tuple[B]
	return Color{NewTuple(r, g, b)}
}

func (c Color) Equals(other Color) bool {
	return c.Tuple.Equals(other.Tuple)
}
