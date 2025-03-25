package main

import . "github.com/michaelzhao820/raytracer/raytracer"

func main() {
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
