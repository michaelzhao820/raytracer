package raytracer

import "math"

type Ray struct {
	origin    Tuple
	direction Tuple
}

type Shape interface {
	transformMatrix() Matrix
}

type Intersection struct {
	t float64
	o Shape
}

func NewRay(origin, direction Tuple) Ray {
	return Ray{origin: origin, direction: direction}
}

func (r Ray) Position(t float64) (Tuple, error) {
	return r.origin.Add(r.direction.Multiply(t))
}

func Intersections(args ...Intersection) []Intersection {
	return args
}

func Hit(args []Intersection) *Intersection {
	currentMin := math.Inf(1)
	var returnIntersection *Intersection = nil
	for i := 0; i < len(args); i++ {
		if args[i].t < currentMin && args[i].t > 0 {
			currentMin = args[i].t
			returnIntersection = &args[i]
		}
	}
	return returnIntersection
}

func (r Ray) Intersect(s Shape) []Intersection {
	m, _ := s.transformMatrix().Inverse()
	r2 := r.Transform(m)
	sphereToRay, _ := r2.origin.Subtract(NewPoint(0, 0, 0))
	a, _ := Dot(r2.Direction(), r2.Direction())
	b, _ := Dot(r2.Direction(), sphereToRay)
	b *= 2
	c, _ := Dot(sphereToRay, sphereToRay)
	c -= 1
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return nil
	} else {
		t1 := (-1*b - math.Sqrt(discriminant)) / (2 * a)
		t2 := (-1*b + math.Sqrt(discriminant)) / (2 * a)
		return []Intersection{{t: t1, o: s}, {t: t2, o: s}}
	}
}

func (r Ray) Transform(m Matrix) Ray {
	origin, _ := m.MultiplyWithTuple(r.origin)
	direction, _ := m.MultiplyWithTuple(r.direction)
	return NewRay(origin, direction)
}

func (r Ray) Origin() Tuple {
	return r.origin
}

func (r Ray) Direction() Tuple {
	return r.direction
}
