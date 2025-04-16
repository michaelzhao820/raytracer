package raytracer

import "math"

const EPSILON float64 = 0.00001

type Plane struct {
	transform Matrix
	material  *Material
}

func NewPlane() *Plane {
	return &Plane{
		transform: IdentityMatrix(),
		material:  DefaultMaterial(),
	}
}

func (p *Plane) SetTransform(m Matrix) {
	p.transform = m
}

func (p *Plane) GetTransformMatrix() Matrix {
	return p.transform
}

func (p *Plane) GetMaterial() *Material {
	return p.material
}

func (p *Plane) NormalAt(worldPoint Tuple) Tuple {
	return NewVector(0, 1, 0)
}

func (p *Plane) Intersect(r Ray) []Intersection {
	if math.Abs(r.direction[Y]) < EPSILON {
		return nil
	}
	return []Intersection{{
		t: -r.origin[Y] / r.direction[Y],
		o: p,
	}}
}

func (p *Plane) SetMaterial(material *Material) {
	p.material = material
}
