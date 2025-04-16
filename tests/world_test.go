package tests

import (
	. "github.com/michaelzhao820/raytracer/raytracer"
	"testing"
)

func TestWorld(t *testing.T) {
	t.Run("Creating a world", func(t *testing.T) {
		w := NewWorld()
		if len(w.GetObjects()) != 0 {
			t.Errorf("Expected world to have 0 GetObjects, got %d", len(w.GetObjects()))
		}
		if w.GetLight() != nil {
			t.Errorf("Expected world to have no light source, got %+v", w.GetLight())
		}
	})

	t.Run("The default world", func(t *testing.T) {
		w := NewWorld()
		w.DefaultWorld()

		light := Light{
			Position:  NewPoint(-10, 10, -10),
			Intensity: NewColor(1, 1, 1),
		}

		if !w.GetLight().Position.Equals(light.Position) {
			t.Errorf("Expected default world to have light %v, got %v", light, w.GetLight())
		}

		if !w.GetLight().Intensity.Equals(light.Intensity) {
			t.Errorf("Expected default world to have light %v, got %v", light, w.GetLight())
		}

		if len(w.GetObjects()) != 2 {
			t.Fatalf("Expected default world to have 2 GetObjects, got %d", len(w.GetObjects()))
		}
		// We won't compare deep equality here since s1/s2 are set up by DefaultWorld.
		// But you can if needed using Equals on materials and transforms.
	})

	t.Run("Intersect a world with a ray", func(t *testing.T) {
		w := NewWorld()
		w.DefaultWorld()

		r := NewRay(
			NewPoint(0, 0, -5),
			NewVector(0, 0, 1),
		)

		xs := w.IntersectWorld(r)
		if len(xs) != 4 {
			t.Fatalf("Expected 4 intersections, got %d", len(xs))
		}

		expectedTs := []float64{4, 4.5, 5.5, 6}
		for i, expected := range expectedTs {
			if !almostEqual(xs[i].GetTime(), expected) {
				t.Errorf("Expected xs[%d].t = %v, got %v", i, expected, xs[i].GetTime())
			}
		}
	})
	t.Run("Shading an intersection", func(t *testing.T) {
		w := NewWorld()
		w.DefaultWorld()

		r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
		shape := w.GetObjects()[0]
		i := NewIntersection(4, shape)

		comps := PrepareComputations(i, r)
		c := w.ShadeHits(comps)

		expected := NewColor(0.38066, 0.047583, 0.2855)
		if !c.Equals(expected) {
			t.Errorf("Expected shaded color = %v, got %v", expected, c)
		}
	})

	t.Run("Shading an intersection from the inside", func(t *testing.T) {
		w := NewWorld()
		w.DefaultWorld()

		// override the light position for this test
		w.GetLight().Position = NewPoint(0, 0.25, 0)

		r := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
		shape := w.GetObjects()[1]
		i := NewIntersection(0.5, shape)

		comps := PrepareComputations(i, r)
		c := w.ShadeHits(comps)

		expected := NewColor(0.90498, 0.90498, 0.90498)
		if !c.Equals(expected) {
			t.Errorf("Expected shaded color = %v, got %v", expected, c)
		}
	})
}

func TestWorldShadows(t *testing.T) {
	t.Run("There is no shadow when nothing is collinear with point and light", func(t *testing.T) {
		w := NewWorld()
		w.DefaultWorld()

		p := NewPoint(0, 10, 0)
		if w.IsShadowed(p) {
			t.Errorf("Expected point %v not to be in shadow", p)
		}
	})

	t.Run("The shadow when an object is between the point and the light", func(t *testing.T) {
		w := NewWorld()
		w.DefaultWorld()

		p := NewPoint(10, -10, 10)
		if !w.IsShadowed(p) {
			t.Errorf("Expected point %v to be in shadow", p)
		}
	})

	t.Run("There is no shadow when an object is behind the light", func(t *testing.T) {
		w := NewWorld()
		w.DefaultWorld()

		p := NewPoint(-20, 20, -20)
		if w.IsShadowed(p) {
			t.Errorf("Expected point %v not to be in shadow", p)
		}
	})

	t.Run("There is no shadow when an object is behind the point", func(t *testing.T) {
		w := NewWorld()
		w.DefaultWorld()

		p := NewPoint(-2, 2, -2)
		if w.IsShadowed(p) {
			t.Errorf("Expected point %v not to be in shadow", p)
		}
	})
}

func almostEqual(a, b float64) bool {
	const epsilon = 1e-5
	return (a-b) < epsilon && (b-a) < epsilon
}
