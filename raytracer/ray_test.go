package raytracer

import (
	"math"
	"testing"
)

func TestRayCreation(t *testing.T) {

	origin := NewPoint(1, 2, 3)
	direction := NewVector(4, 5, 6)

	r := NewRay(origin, direction)

	// Then
	if !r.Origin().Equals(origin) {
		t.Errorf("Ray origin does not match expected value. Got %v, want %v", r.Origin(), origin)
	}

	if !r.Direction().Equals(direction) {
		t.Errorf("Ray direction does not match expected value. Got %v, want %v", r.Direction(), direction)
	}
}

func TestRayPosition(t *testing.T) {

	origin := NewPoint(2, 3, 4)
	direction := NewVector(1, 0, 0)
	r := NewRay(origin, direction)

	// Test cases
	testCases := []struct {
		t        float64
		expected Tuple
	}{
		{0, NewPoint(2, 3, 4)},
		{1, NewPoint(3, 3, 4)},
		{-1, NewPoint(1, 3, 4)},
		{2.5, NewPoint(4.5, 3, 4)},
	}

	for _, tc := range testCases {
		position, err := r.Position(tc.t)

		if err != nil {
			t.Errorf("Error calculating position at t=%v: %v", tc.t, err)
			continue
		}

		if !position.Equals(tc.expected) {
			t.Errorf("Position at t=%v incorrect. Got %v, want %v", tc.t, position, tc.expected)
		}
	}
}

func TestIntersection(t *testing.T) {
	t.Run("An intersection encapsulates t and object", func(t *testing.T) {
		s := NewSphere()
		i := Intersection{
			t: 3.5,
			o: s,
		}

		if i.t != 3.5 {
			t.Errorf("Expected i.T = 3.5, got %f", i.t)
		}

		if i.o != s {
			t.Errorf("Expected i.Object = s")
		}
	})
}
func TestIntersections(t *testing.T) {
	t.Run("Aggregating intersections", func(t *testing.T) {
		s := NewSphere()
		i1 := Intersection{
			t: 1,
			o: s,
		}
		i2 := Intersection{
			t: 2,
			o: s,
		}

		xs := Intersections(i1, i2)

		if len(xs) != 2 {
			t.Errorf("Expected xs.count = 2, got %d", len(xs))
		}

		if xs[0].t != 1 {
			t.Errorf("Expected xs[0].T = 1, got %f", xs[0].t)
		}

		if xs[1].t != 2 {
			t.Errorf("Expected xs[1].T = 2, got %f", xs[1].t)
		}
	})
}

func TestRaySphereIntersection(t *testing.T) {
	// Small value for floating point comparisons
	const epsilon = 0.0001

	t.Run("A ray intersects a sphere at two points", func(t *testing.T) {
		r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
		s := NewSphere()

		xs := r.Intersect(s)

		if len(xs) != 2 {
			t.Errorf("Expected 2 intersections, got %d", len(xs))
		}

		if math.Abs(xs[0].t-4.0) > epsilon {
			t.Errorf("Expected xs[0].T = 4.0, got %f", xs[0].t)
		}

		if math.Abs(xs[1].t-6.0) > epsilon {
			t.Errorf("Expected xs[1].T = 6.0, got %f", xs[1].t)
		}
	})

	t.Run("A ray intersects a sphere at a tangent", func(t *testing.T) {
		r := NewRay(NewPoint(0, 1, -5), NewVector(0, 0, 1))
		s := NewSphere()

		xs := r.Intersect(s)

		if len(xs) != 2 {
			t.Errorf("Expected 2 intersections, got %d", len(xs))
		}

		if math.Abs(xs[0].t-5.0) > epsilon {
			t.Errorf("Expected xs[0].T = 5.0, got %f", xs[0].t)
		}

		if math.Abs(xs[1].t-5.0) > epsilon {
			t.Errorf("Expected xs[1].T = 5.0, got %f", xs[1].t)
		}
	})

	t.Run("A ray misses a sphere", func(t *testing.T) {
		r := NewRay(NewPoint(0, 2, -5), NewVector(0, 0, 1))
		s := NewSphere()

		xs := r.Intersect(s)

		if len(xs) != 0 {
			t.Errorf("Expected 0 intersections, got %d", len(xs))
		}
	})

	t.Run("A ray originates inside a sphere", func(t *testing.T) {
		r := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
		s := NewSphere()

		xs := r.Intersect(s)

		if len(xs) != 2 {
			t.Errorf("Expected 2 intersections, got %d", len(xs))
		}

		if math.Abs(xs[0].t-(-1.0)) > epsilon {
			t.Errorf("Expected xs[0].T = -1.0, got %f", xs[0].t)
		}

		if math.Abs(xs[1].t-1.0) > epsilon {
			t.Errorf("Expected xs[1].T = 1.0, got %f", xs[1].t)
		}
	})

	t.Run("A sphere is behind a ray", func(t *testing.T) {
		r := NewRay(NewPoint(0, 0, 5), NewVector(0, 0, 1))
		s := NewSphere()

		xs := r.Intersect(s)

		if len(xs) != 2 {
			t.Errorf("Expected 2 intersections, got %d", len(xs))
		}

		if math.Abs(xs[0].t-(-6.0)) > epsilon {
			t.Errorf("Expected xs[0].T = -6.0, got %f", xs[0].t)
		}

		if math.Abs(xs[1].t-(-4.0)) > epsilon {
			t.Errorf("Expected xs[1].T = -4.0, got %f", xs[1].t)
		}
	})

	t.Run("Intersect sets the object on the intersection", func(t *testing.T) {
		r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
		s := NewSphere()

		xs := r.Intersect(s)

		if len(xs) != 2 {
			t.Errorf("Expected 2 intersections, got %d", len(xs))
		}

		if xs[0].o != s {
			t.Errorf("Expected xs[0].object = s")
		}

		if xs[1].o != s {
			t.Errorf("Expected xs[1].object = s")
		}
	})
}

func TestHitFunction(t *testing.T) {
	t.Run("All intersections have positive t", func(t *testing.T) {
		s := NewSphere()
		i1 := Intersection{t: 1, o: s}
		i2 := Intersection{t: 2, o: s}
		xs := Intersections(i2, i1)
		hit := Hit(xs)

		if *hit != i1 {
			t.Errorf("Expected hit to be i1, got %+v", hit)
		}
	})

	t.Run("Some intersections have negative t", func(t *testing.T) {
		s := NewSphere()
		i1 := Intersection{t: -1, o: s}
		i2 := Intersection{t: 1, o: s}
		xs := Intersections(i2, i1)
		hit := Hit(xs)

		if *hit != i2 {
			t.Errorf("Expected hit to be i2, got %+v", hit)
		}
	})

	t.Run("All intersections have negative t", func(t *testing.T) {
		s := NewSphere()
		i1 := Intersection{t: -2, o: s}
		i2 := Intersection{t: -1, o: s}
		xs := Intersections(i2, i1)
		hit := Hit(xs)

		if hit != nil {
			t.Errorf("Expected hit to be nothing (t = +Inf), got %+v", hit)
		}
	})

	t.Run("Hit is lowest non-negative t", func(t *testing.T) {
		s := NewSphere()
		i1 := Intersection{t: 5, o: s}
		i2 := Intersection{t: 7, o: s}
		i3 := Intersection{t: -3, o: s}
		i4 := Intersection{t: 2, o: s}
		xs := Intersections(i1, i2, i3, i4)
		hit := Hit(xs)

		if *hit != i4 {
			t.Errorf("Expected hit to be i4, got %+v", hit)
		}
	})
}

func TestRayTransform(t *testing.T) {
	t.Run("Translating a ray", func(t *testing.T) {
		r := NewRay(NewPoint(1, 2, 3), NewVector(0, 1, 0))
		m, err := TranslationMatrix(3, 4, 5)
		if err != nil {
			t.Fatalf("failed to create translation matrix: %v", err)
		}
		r2 := r.Transform(m)

		expectedOrigin := NewPoint(4, 6, 8)
		expectedDirection := NewVector(0, 1, 0)

		if !r2.origin.Equals(expectedOrigin) {
			t.Errorf("expected origin %v, got %v", expectedOrigin, r2.origin)
		}
		if !r2.direction.Equals(expectedDirection) {
			t.Errorf("expected direction %v, got %v", expectedDirection, r2.direction)
		}
	})

	t.Run("Scaling a ray", func(t *testing.T) {
		r := NewRay(NewPoint(1, 2, 3), NewVector(0, 1, 0))
		m, err := ScalingMatrix(2, 3, 4)
		if err != nil {
			t.Fatalf("failed to create scaling matrix: %v", err)
		}
		r2 := r.Transform(m)

		expectedOrigin := NewPoint(2, 6, 12)
		expectedDirection := NewVector(0, 3, 0)

		if !r2.origin.Equals(expectedOrigin) {
			t.Errorf("expected origin %v, got %v", expectedOrigin, r2.origin)
		}
		if !r2.direction.Equals(expectedDirection) {
			t.Errorf("expected direction %v, got %v", expectedDirection, r2.direction)
		}
	})
}

func TestRaySphereTransformIntersection(t *testing.T) {
	t.Run("Intersecting a scaled sphere with a ray", func(t *testing.T) {
		r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
		s := NewSphere()
		sm, _ := ScalingMatrix(2, 2, 2)
		s.SetTransform(sm)

		xs := r.Intersect(s)

		if len(xs) != 2 {
			t.Fatalf("Expected 2 intersections, got %d", len(xs))
		}
		if math.Abs(xs[0].t-3.0) > 1e-5 {
			t.Errorf("Expected xs[0].t = 3.0, got %f", xs[0].t)
		}
		if math.Abs(xs[1].t-7.0) > 1e-5 {
			t.Errorf("Expected xs[1].t = 7.0, got %f", xs[1].t)
		}
	})

	t.Run("Intersecting a translated sphere with a ray", func(t *testing.T) {
		r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
		s := NewSphere()
		tm, _ := TranslationMatrix(5, 0, 0)
		s.SetTransform(tm)

		xs := r.Intersect(s)

		if len(xs) != 0 {
			t.Fatalf("Expected 0 intersections, got %d", len(xs))
		}
	})
}

func TestPrepareComputations(t *testing.T) {
	t.Run("Precomputing the state of an intersection", func(t *testing.T) {
		r := NewRay(
			NewPoint(0, 0, -5),
			NewVector(0, 0, 1),
		)
		shape := NewSphere()

		i := Intersection{
			t: 4,
			o: shape,
		}

		comps := PrepareComputations(i, r)

		if comps.t != i.GetTime() {
			t.Errorf("Expected comps.t = %v, got %v", i.GetTime(), comps.t)
		}

		if comps.o != i.GetObject() {
			t.Errorf("Expected comps.object to equal i.object")
		}

		expectedPoint := NewPoint(0, 0, -1)
		if !comps.point.Equals(expectedPoint) {
			t.Errorf("Expected comps.point = %v, got %v", expectedPoint, comps.point)
		}

		expectedEyeV := NewVector(0, 0, -1)
		if !comps.eyev.Equals(expectedEyeV) {
			t.Errorf("Expected comps.eyev = %v, got %v", expectedEyeV, comps.eyev)
		}

		expectedNormalV := NewVector(0, 0, -1)
		if !comps.normalv.Equals(expectedNormalV) {
			t.Errorf("Expected comps.normalv = %v, got %v", expectedNormalV, comps.normalv)
		}
	})
	t.Run("The hit, when an intersection occurs on the outside", func(t *testing.T) {
		r := NewRay(
			NewPoint(0, 0, -5),
			NewVector(0, 0, 1),
		)
		shape := NewSphere()

		i := Intersection{
			t: 4,
			o: shape,
		}

		comps := PrepareComputations(i, r)

		if comps.inside {
			t.Errorf("Expected comps.inside = false, got true")
		}
	})

	t.Run("The hit, when an intersection occurs on the inside", func(t *testing.T) {
		r := NewRay(
			NewPoint(0, 0, 0),
			NewVector(0, 0, 1),
		)
		shape := NewSphere()

		i := Intersection{
			t: 1,
			o: shape,
		}

		comps := PrepareComputations(i, r)

		expectedPoint := NewPoint(0, 0, 1)
		if !comps.point.Equals(expectedPoint) {
			t.Errorf("Expected comps.point = %v, got %v", expectedPoint, comps.point)
		}

		expectedEyeV := NewVector(0, 0, -1)
		if !comps.eyev.Equals(expectedEyeV) {
			t.Errorf("Expected comps.eyev = %v, got %v", expectedEyeV, comps.eyev)
		}

		if !comps.inside {
			t.Errorf("Expected comps.inside = true, got false")
		}

		expectedNormalV := NewVector(0, 0, -1)
		if !comps.normalv.Equals(expectedNormalV) {
			t.Errorf("Expected comps.normalv = %v, got %v", expectedNormalV, comps.normalv)
		}
	})
}

func TestWorldRay(t *testing.T) {
	t.Run("The color when a ray misses", func(t *testing.T) {
		w := NewWorld()
		w.DefaultWorld()

		r := NewRay(NewPoint(0, 0, -5), NewVector(0, 1, 0)) // ray goes up, misses both spheres
		c := w.ColorAt(r)

		expected := NewColor(0, 0, 0)
		if !c.Equals(expected) {
			t.Errorf("Expected color = %v, got %v", expected, c)
		}
	})

	t.Run("The color when a ray hits", func(t *testing.T) {
		w := NewWorld()
		w.DefaultWorld()

		r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1)) // ray hits the first sphere
		c := w.ColorAt(r)

		expected := NewColor(0.38066, 0.047583, 0.2855)
		if !c.Equals(expected) {
			t.Errorf("Expected color = %v, got %v", expected, c)
		}
	})

	t.Run("The color with an intersection behind the ray", func(t *testing.T) {
		w := NewWorld()
		w.DefaultWorld()

		// Get references to the two spheres
		outer := w.GetObjects()[0]
		inner := w.GetObjects()[1]

		// Set both ambient to 1 (force full material color contribution)
		outer.GetMaterial().ambient = 1
		inner.GetMaterial().ambient = 1

		r := NewRay(NewPoint(0, 0, 0.75), NewVector(0, 0, -1)) // intersects both, inner closer
		c := w.ColorAt(r)

		expected := inner.GetMaterial().color
		if !c.Equals(expected) {
			t.Errorf("Expected color = %v, got %v", expected, c)
		}
	})
}
