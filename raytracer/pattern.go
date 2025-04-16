package raytracer

import "math"

type Pattern interface {
	PatternAtObject(obj Shape, point Tuple) Color
	PatternAt(point Tuple) Color
	GetTransform() Matrix
	SetTransform(m Matrix)
}

// Helper to handle world → object → pattern space transform
func patternPointFor(p Pattern, obj Shape, worldPoint Tuple) Tuple {
	objInv, _ := obj.GetTransformMatrix().Inverse()
	objPoint, _ := objInv.MultiplyWithTuple(worldPoint)
	patInv, _ := p.GetTransform().Inverse()
	ret, _ := patInv.MultiplyWithTuple(objPoint)
	return ret
}

////////////////////////////////////////////////////////////////////////////////

type StripePattern struct {
	a, b      Color
	transform Matrix
}

func NewStripePattern(a, b Color) *StripePattern {
	return &StripePattern{
		a:         a,
		b:         b,
		transform: IdentityMatrix(),
	}
}

func (sp *StripePattern) PatternAtObject(obj Shape, point Tuple) Color {
	return sp.PatternAt(patternPointFor(sp, obj, point))
}

func (sp *StripePattern) PatternAt(point Tuple) Color {
	if int(math.Floor(point[X]))%2 == 0 {
		return sp.a
	}
	return sp.b
}

func (sp *StripePattern) GetTransform() Matrix {
	return sp.transform
}

func (sp *StripePattern) SetTransform(m Matrix) {
	sp.transform = m
}

////////////////////////////////////////////////////////////////////////////////

type GradientPattern struct {
	a, b      Color
	transform Matrix
}

func NewGradientPattern(a, b Color) *GradientPattern {
	return &GradientPattern{
		a:         a,
		b:         b,
		transform: IdentityMatrix(),
	}
}

func (gp *GradientPattern) PatternAtObject(obj Shape, point Tuple) Color {
	return gp.PatternAt(patternPointFor(gp, obj, point))
}

func (gp *GradientPattern) PatternAt(point Tuple) Color {
	distance := gp.b.SubtractColor(gp.a)
	fraction := point[X] - math.Floor(point[X])
	return gp.a.AddColor(distance.MultiplyByScalar(fraction))
}

func (gp *GradientPattern) GetTransform() Matrix {
	return gp.transform
}

func (gp *GradientPattern) SetTransform(m Matrix) {
	gp.transform = m
}

////////////////////////////////////////////////////////////////////////////////

type RingPattern struct {
	a, b      Color
	transform Matrix
}

func NewRingPattern(a, b Color) *RingPattern {
	return &RingPattern{
		a:         a,
		b:         b,
		transform: IdentityMatrix(),
	}
}

func (rp *RingPattern) PatternAtObject(obj Shape, point Tuple) Color {
	return rp.PatternAt(patternPointFor(rp, obj, point))
}

func (rp *RingPattern) PatternAt(point Tuple) Color {
	dist := math.Sqrt(point[X]*point[X] + point[Z]*point[Z])
	if int(math.Floor(dist))%2 == 0 {
		return rp.a
	}
	return rp.b
}

func (rp *RingPattern) GetTransform() Matrix {
	return rp.transform
}

func (rp *RingPattern) SetTransform(m Matrix) {
	rp.transform = m
}

////////////////////////////////////////////////////////////////////////////////

type Checker3DPattern struct {
	a, b      Color
	transform Matrix
}

func NewChecker3DPattern(a, b Color) *Checker3DPattern {
	return &Checker3DPattern{
		a:         a,
		b:         b,
		transform: IdentityMatrix(),
	}
}

func (cp *Checker3DPattern) PatternAtObject(obj Shape, point Tuple) Color {
	return cp.PatternAt(patternPointFor(cp, obj, point))
}

func (cp *Checker3DPattern) PatternAt(point Tuple) Color {
	sum := int(math.Floor(point[X]) + math.Floor(point[Y]) + math.Floor(point[Z]))
	if sum%2 == 0 {
		return cp.a
	}
	return cp.b
}

func (cp *Checker3DPattern) GetTransform() Matrix {
	return cp.transform
}

func (cp *Checker3DPattern) SetTransform(m Matrix) {
	cp.transform = m
}
