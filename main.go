package main

import (
	. "github.com/michaelzhao820/raytracer/raytracer"
)

func main() {

	rayOrigin := NewPoint(0, 0, -5)

	wallZPosition := 10.0
	wallSize := 7.0
	half := wallSize / 2

	canvasPixels := 100
	c := NewCanvas(canvasPixels, canvasPixels)

	pixelSize := wallSize / float64(canvasPixels)

	shape := NewSphere()
	shape.GetMaterial().SetColor(1, 0.2, 1)

	light := Light{
		Position:  NewPoint(-10, 10, -10),
		Intensity: NewColor(1, 1, 1),
	}

	for y := 0; y < canvasPixels; y++ {
		for x := 0; x < canvasPixels; x++ {
			worldX := -half + pixelSize*float64(x)
			worldY := half - pixelSize*float64(y)
			position := NewPoint(
				worldX,
				worldY,
				wallZPosition,
			)
			direction, _ := position.Subtract(rayOrigin)
			direction, _ = direction.Normalize()
			r := NewRay(rayOrigin, direction)
			xs := r.Intersect(shape)
			hit := Hit(xs)
			if hit != nil {
				point, _ := r.Position(hit.GetTime())
				normal := hit.GetObject().NormalAt(point)
				eye := r.Direction().Multiply(-1)

				color := Lighting(*hit.GetObject().GetMaterial(), light, point, eye, normal)
				c.WritePixel(x, y, color)
			}
		}
	}
	err := c.CanvasToPPM("sphere.ppm")
	if err != nil {
		return
	}

	/*
		canvasPixels := 500
		c := NewCanvas(canvasPixels, canvasPixels)

		clockRadius := float64(canvasPixels) * 3 / 8
		centerX := canvasPixels / 2
		centerY := canvasPixels / 2

		p := NewPoint(0, 0, 1)

		for i := 0; i < TotalHours; i++ {
			r, _ := RotationYMatrix(math.Pi / 6 * float64(i))
			n, _ := r.MultiplyWithTuple(p)

			scaledX := n[X] * clockRadius
			scaledZ := n[Z] * clockRadius

			c.WritePixel(centerX+int(scaledX), centerY+int(scaledZ), NewColor(1, 0, 0))
		}

		err := c.CanvasToPPM("clock.ppm")
		if err != nil {
			return
		}
	*/

	/*
		projectile := struct {
			position Tuple
			velocity Tuple
		}{position: NewPoint(0, 1, 0),
			velocity: NewVector(1, 1.8, 0),
		}
		environment := struct {
			gravity Tuple
			wind    Tuple
		}{
			gravity: NewVector(0, -0.1, 0),
			wind:    NewVector(-0.01, 0, 0),
		}

		projectile.velocity, _ = projectile.velocity.Normalize()
		projectile.velocity = projectile.velocity.Multiply(11.25)
		c := NewCanvas(900, 550)
		for projectile.position[Y] > 0 {
			tick(&projectile, &environment, &c)
		}
		err := c.CanvasToPPM("projectile.ppm")
		if err != nil {
			return
		}

	*/

}

func tick(proj *struct {
	position Tuple
	velocity Tuple
}, env *struct {
	gravity Tuple
	wind    Tuple
}, c *Canvas,
) {
	proj.position, _ = proj.position.Add(proj.velocity)
	proj.velocity, _ = proj.velocity.Add(env.gravity)
	proj.velocity, _ = proj.velocity.Add(env.wind)

	c.WritePixel(int(proj.position[X]), c.Height-int(proj.position[Y])-1, NewColor(1, 0, 0))

}
