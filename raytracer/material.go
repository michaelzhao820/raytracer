package raytracer

import "math"

type Light struct {
	Position  Tuple
	Intensity Color
}

type Material struct {
	color     Color
	ambient   float64
	diffuse   float64
	specular  float64
	shininess float64
	Pattern   Pattern
}

func DefaultMaterial() *Material {
	return &Material{
		color:     NewColor(1, 1, 1),
		ambient:   0.1,
		diffuse:   0.9,
		specular:  0.9,
		shininess: 200.0,
	}
}

func (m *Material) SetColor(r, g, b float64) {
	m.color = NewColor(r, g, b)
}

func (m *Material) SetDiffuse(f float64) {
	m.diffuse = f
}

func (m *Material) SetSpecular(f float64) {
	m.specular = f

}

func (m *Material) SetAmbient(f float64) {
	m.ambient = f

}

func Lighting(material Material, object Shape, light Light, point, eyev, normalv Tuple, inShadow bool) Color {
	var diffuse, specular, ambient, color Color

	if material.Pattern != nil {
		color = material.Pattern.PatternAtObject(object, point)
	} else {
		color = material.color
	}

	effectiveColor := color.MultiplyOtherColor(light.Intensity)

	lightv, _ := light.Position.Subtract(point)
	lightv, _ = lightv.Normalize()

	ambient = effectiveColor.MultiplyByScalar(material.ambient)
	if inShadow {
		return ambient
	}

	lightDotNormal, _ := Dot(lightv, normalv)
	if lightDotNormal < 0 {
		diffuse = NewColor(0, 0, 0)
		specular = NewColor(0, 0, 0)
	} else {
		diffuse = effectiveColor.MultiplyByScalar(material.diffuse).MultiplyByScalar(lightDotNormal)

		reflectv, _ := Reflect(lightv.Multiply(-1), normalv)
		reflectDotEye, _ := Dot(reflectv, eyev)

		if reflectDotEye < 0 {
			specular = NewColor(0, 0, 0)
		} else {
			factor := math.Pow(reflectDotEye, material.shininess)
			specular = light.Intensity.MultiplyByScalar(material.specular).MultiplyByScalar(factor)
		}

	}
	return ambient.AddColor(diffuse).AddColor(specular)
}
