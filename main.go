package main

import (
	. "github.com/michaelzhao820/raytracer/raytracer"
	"math"
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
	s, _ := ScalingMatrix(0.5, 1, 1)
	ry, _ := RotationZMatrix(math.Pi / 4)
	t, _ := ry.MultiplyMatrices(s)
	shape.SetTransform(t)

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
			if Hit(xs) != nil {
				c.WritePixel(x, y, NewColor(1, 0, 0))
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
