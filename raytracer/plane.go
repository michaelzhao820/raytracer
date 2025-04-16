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

	localNormal := NewVector(0, 1, 0)

	//convert back into world space
	tm, _ := p.GetTransformMatrix().Inverse()
	tmt, _ := tm.Transpose()
	worldNormal, _ := tmt.MultiplyWithTuple(localNormal)
	worldNormal[W] = 0
	worldNormal, _ = worldNormal.Normalize()
	return worldNormal
}

func (p *Plane) Intersect(r Ray) []Intersection {
	//Convert ray into object space
	inv, _ := p.transform.Inverse()
	localRay := r.Transform(inv)

	if math.Abs(localRay.direction[Y]) < EPSILON {
		return nil
	}
	return []Intersection{{
		t: -localRay.origin[Y] / localRay.direction[Y],
		o: p,
	}}
}

func (p *Plane) SetMaterial(material *Material) {
	p.material = material
}
