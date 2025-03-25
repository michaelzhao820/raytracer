package raytracer

import (
	"fmt"
	"os"
)

type Canvas struct {
	width, Height int
	pixels        [][]Color
}

func NewCanvas(width, height int) Canvas {
	pixels := make([][]Color, height)
	for i := range pixels {
		pixels[i] = make([]Color, width)
		for j := range pixels[i] {
			pixels[i][j] = NewColor(0, 0, 0) // Default color (black)
		}
	}
	return Canvas{width, height, pixels}
}

func (c *Canvas) WritePixel(x, y int, color Color) {
	if x >= 0 && x < c.width && y >= 0 && y < c.Height {
		c.pixels[y][x] = color
	}
}

func (c *Canvas) CanvasToPPM(filename string) error {
	// Open the file for writing
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the PPM header
	_, err = fmt.Fprintf(file, "P3\n%d %d\n255\n", c.width, c.Height)
	if err != nil {
		return err
	}
	for i := 0; i < c.Height; i++ {
		lineLength := 0
		for j := 0; j < c.width; j++ {
			color := c.pixels[i][j]
			pixelStr := fmt.Sprintf("%d %d %d",
				int(clamp(color.Tuple[R]*255, 0, 255)),
				int(clamp(color.Tuple[G]*255, 0, 255)),
				int(clamp(color.Tuple[B]*255, 0, 255)))

			if lineLength+len(pixelStr)+1 > 70 {
				fmt.Fprintln(file)
				lineLength = 0
			}

			if lineLength > 0 {
				fmt.Fprint(file, " ")
				lineLength++
			}
			fmt.Fprint(file, pixelStr)
			lineLength += len(pixelStr)
		}
		fmt.Fprintln(file)
	}
	fmt.Fprintln(file)
	return nil
}

// Helper function to clamp color values between min and max
func clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
