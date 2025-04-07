package raytracer

type Sphere struct {
	transform Matrix
}

func NewSphere() *Sphere {
	return &Sphere{
		transform: IdentityMatrix(),
	}
}

func (s *Sphere) SetTransform(m Matrix) {
	s.transform = m
}

func (s *Sphere) transformMatrix() Matrix {
	return s.transform
}
