package raytracer

import "math"

type Camera struct {
	hsize       float64
	vsize       float64
	fieldOfView float64
	transform   Matrix
	pixelSize   float64
	halfWidth   float64
	halfHeight  float64
}

func NewCamera(hsize, vsize, fieldOfView float64) *Camera {
	return &Camera{hsize, vsize, fieldOfView, IdentityMatrix(),
		calculatePixelSize(hsize, vsize, fieldOfView), computeHalfWidth(hsize, vsize, fieldOfView),
		computeHalfHeight(hsize, vsize, fieldOfView)}
}

func computeHalfHeight(hsize float64, vsize float64, fieldOfView float64) float64 {
	halfView := math.Tan(fieldOfView / 2)
	aspect := hsize / vsize

	if aspect >= 1 {
		return halfView / aspect
	} else {
		return halfView
	}
}

func computeHalfWidth(hsize float64, vsize float64, fieldOfView float64) float64 {
	halfView := math.Tan(fieldOfView / 2)
	aspect := hsize / vsize

	if aspect >= 1 {
		return halfView
	} else {
		return halfView * aspect
	}

}

func calculatePixelSize(hsize float64, vsize float64, fieldOfView float64) float64 {
	return (computeHalfWidth(hsize, vsize, fieldOfView) * 2) / hsize
}

// rayForPixel generates a ray in world space that starts at the camera's position
// and passes through the center of the specified pixel (px, py) on the virtual canvas.
//
// The camera assumes a default orientation where it is positioned at the origin,
// looking straight down the negative Z axis, with +Y as "up". The canvas is placed
// at z = -1 in this camera-local coordinate system.
//
// The pixel loop (in Render) iterates over camera-local canvas coordinates, with
// (0,0) referring to the upper-left corner of the canvas. Each (px, py) refers to a
// pixel in this local grid.
//
// To determine where in the world this pixel lies, we compute its position in camera
// space (worldX, worldY, -1), and then apply the inverse of the camera's view
// transformation. This effectively transforms the point from camera space into
// world space — i.e., it answers: "Where is this pixel in the actual world, given the
// camera’s position and orientation?"
//
// Similarly, the camera's local origin (0,0,0) is transformed into world space to find
// the ray's true world-space origin.
//
// Finally, the ray is created with the world-space origin and a direction pointing
// toward the world-space pixel, normalized to ensure a unit-length direction vector.

func (c *Camera) rayForPixel(px, py float64) Ray {
	xoffset := (px + 0.5) * c.pixelSize
	yoffset := (py + 0.5) * c.pixelSize

	cameraX := c.halfWidth - xoffset
	cameraY := c.halfHeight - yoffset

	pointTM, _ := c.transform.Inverse()

	pixel, _ := pointTM.MultiplyWithTuple(NewPoint(cameraX, cameraY, -1))
	origin, _ := pointTM.MultiplyWithTuple(NewPoint(0, 0, 0))

	direction, _ := pixel.Subtract(origin)
	direction, _ = direction.Normalize()

	return NewRay(origin, direction)
}

func (c *Camera) Render(w World) error {

	image := NewCanvas(int(c.hsize), int(c.vsize))

	for y := 0; y < int(c.vsize); y++ {
		for x := 0; x < int(c.hsize); x++ {
			ray := c.rayForPixel(float64(x), float64(y))
			color := w.ColorAt(ray, 4)
			image.WritePixel(x, y, color)
		}
	}
	err := image.CanvasToPPM("scene.ppm")
	return err
}

func (c *Camera) SetTransform(transform Matrix) {
	c.transform = transform

}
