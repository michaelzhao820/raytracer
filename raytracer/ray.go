package raytracer

import "math"

type Ray struct {
	origin    Tuple
	direction Tuple
}

type Shape interface {
	GetTransformMatrix() Matrix
	NormalAt(x Tuple) Tuple
	GetMaterial() *Material
}

type Intersection struct {
	t float64
	o Shape
}

type Computation struct {
	t         float64
	o         Shape
	point     Tuple
	eyev      Tuple
	normalv   Tuple
	inside    bool
	overpoint Tuple
}

func (i Intersection) GetTime() float64 {
	return i.t
}

func (i Intersection) GetObject() Shape {
	return i.o
}

func NewRay(origin, direction Tuple) Ray {
	return Ray{origin: origin, direction: direction}
}

func NewIntersection(T float64, o Shape) Intersection {
	return Intersection{t: T, o: o}
}

func (r Ray) Position(t float64) (Tuple, error) {
	return r.origin.Add(r.direction.Multiply(t))
}

func Intersections(args ...Intersection) []Intersection {
	return args
}

func PrepareComputations(intersection Intersection, ray Ray) Computation {
	epsilon := 0.00001

	comps := Computation{}
	comps.t = intersection.t
	comps.o = intersection.o
	r, _ := ray.Position(comps.t)
	comps.point = r
	eye := ray.Direction().Multiply(-1)
	comps.eyev = eye

	normal := comps.o.NormalAt(comps.point)
	comps.normalv = normal

	dot, _ := Dot(comps.normalv, comps.eyev)
	if dot < 0 {
		comps.inside = true
		comps.normalv = comps.normalv.Multiply(-1)
	} else {
		comps.inside = false
	}

	add, _ := comps.point.Add(comps.normalv.Multiply(epsilon))
	comps.overpoint = add

	return comps
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
	m, _ := s.GetTransformMatrix().Inverse()
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

// Transform applies the given matrix to the ray, returning a new ray.
// This is used to convert a world-space ray into object space
// by applying the inverse of the object's transform matrix.

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
