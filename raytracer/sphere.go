package raytracer

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
