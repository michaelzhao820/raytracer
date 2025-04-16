package raytracer

import "sort"

type World struct {
	objects []Shape
	light   *Light
}

func NewWorld() *World {
	return &World{
		objects: []Shape{},
		light:   nil,
	}
}

func (w *World) GetLight() *Light {
	return w.light
}

func (w *World) GetObjects() []Shape {
	return w.objects
}

func (w *World) DefaultWorld() {
	w.light = &Light{
		Position:  NewPoint(-10, 10, -10),
		Intensity: NewColor(1, 1, 1),
	}
	s1 := NewSphere()
	s1.material.color = NewColor(0.8, 0.1, 0.6)
	s1.material.diffuse = 0.7
	s1.material.specular = 0.2
	s2 := NewSphere()
	m, _ := ScalingMatrix(0.5, 0.5, 0.5)
	s2.SetTransform(m)
	w.objects = append(w.objects, s1, s2)
}

func (w *World) ColorAt(r Ray) Color {
	xs := w.IntersectWorld(r)
	hit := Hit(xs)
	if hit == nil {
		return NewColor(0.0, 0.0, 0.0)
	}
	c := PrepareComputations(*hit, r)
	return w.ShadeHits(c)
}

func (w *World) IntersectWorld(r Ray) []Intersection {
	var xs []Intersection
	for _, object := range w.objects {
		//We are in object space here to calculate the intersections!
		xs = append(xs, r.Intersect(object)...)
	}
	sort.Slice(xs, func(i, j int) bool {
		return xs[i].t < xs[j].t
	})
	return xs
}

func (w *World) IsShadowed(p Tuple) bool {
	v, _ := w.light.Position.Subtract(p)
	distance, _ := v.Magnitude()
	direction, _ := v.Normalize()

	r := NewRay(p, direction)

	xs := w.IntersectWorld(r)
	h := Hit(xs)
	if h != nil && h.GetTime() < distance {
		return true
	}
	return false
}

func (w *World) ShadeHits(comps Computation) Color {
	return Lighting(*comps.o.GetMaterial(), *w.light, comps.point, comps.eyev, comps.normalv, w.IsShadowed(comps.overpoint))
}

func (w *World) SetLight(l *Light) {
	w.light = l
}

func (w *World) AddObject(o Shape) {
	w.objects = append(w.objects, o)
}
