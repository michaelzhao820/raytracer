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

func (c *Camera) rayForPixel(px, py float64) Ray {
	xoffset := (px + 0.5) * c.pixelSize
	yoffset := (py + 0.5) * c.pixelSize

	worldX := c.halfWidth - xoffset
	worldY := c.halfHeight - yoffset

	pointTM, _ := c.transform.Inverse()
	pixel, _ := pointTM.MultiplyWithTuple(NewPoint(worldX, worldY, -1))
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
			color := w.ColorAt(ray)
			image.WritePixel(x, y, color)
		}
	}
	err := image.CanvasToPPM("scene.ppm")
	return err
}

func (c *Camera) SetTransform(transform Matrix) {
	c.transform = transform

}
