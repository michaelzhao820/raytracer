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
	Intersect(ray Ray) []Intersection
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
	reflectv  Tuple
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

	reflectv, _ := Reflect(ray.Direction(), comps.normalv)
	comps.reflectv = reflectv
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
