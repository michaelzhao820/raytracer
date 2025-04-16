package raytracer

import "math"

type Sphere struct {
	transform Matrix
	material  *Material
}

func NewSphere() *Sphere {
	return &Sphere{
		transform: IdentityMatrix(),
		material:  DefaultMaterial(),
	}
}

func (s *Sphere) SetTransform(m Matrix) {
	s.transform = m
}

func (s *Sphere) GetTransformMatrix() Matrix {
	return s.transform
}

func (s *Sphere) GetMaterial() *Material {
	return s.material
}

func (s *Sphere) NormalAt(worldPoint Tuple) Tuple {

	tm, _ := s.GetTransformMatrix().Inverse()
	objectSpacePoint, _ := tm.MultiplyWithTuple(worldPoint)

	objectSpaceNormal, _ := objectSpacePoint.Subtract(NewPoint(0, 0, 0))
	tmt, _ := tm.Transpose()

	worldNormal, _ := tmt.MultiplyWithTuple(objectSpaceNormal)
	worldNormal[W] = 0
	worldNormal, _ = worldNormal.Normalize()
	return worldNormal
}

func (s *Sphere) Intersect(r Ray) []Intersection {
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

func (s *Sphere) SetMaterial(material *Material) {
	s.material = material
}
