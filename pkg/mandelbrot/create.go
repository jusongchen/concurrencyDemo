package mandelbrot

import (
	"image"
	"image/color"
	"sync"
)

const (
	destDir = "/home/jusong.chen/caddy/images"
	output  = "out.png"
	width   = 2048
	height  = 2048
)

// f, err := os.Create(path.Join(destDir, output))
// if err != nil {
// 	log.Fatal(err)
// }

// img := mandelbrot.CreateImage(width, height)

// if err = png.Encode(f, img); err != nil {
// 	log.Fatal(err)
// }

// CreateImage creates a mandelbrot image with specified size
func CreateImage(width, height int) image.Image {
	m := image.NewGray(image.Rect(0, 0, width, height))
	var w sync.WaitGroup
	w.Add(width)
	for i := 0; i < width; i++ {
		go func(i int) {
			for j := 0; j < height; j++ {
				m.Set(i, j, pixel(i, j, width, height))
			}
			w.Done()
		}(i)
	}
	w.Wait()
	return m
}

// pixel returns the color of a Mandelbrot fractal at the given point.
func pixel(i, j, width, height int) color.Color {
	// Play with this constant to increase the complexity of the fractal.
	// In the justforfunc.com video this was set to 4.
	const complexity = 1024

	xi := norm(i, width, -1.0, 2)
	yi := norm(j, height, -1, 1)

	const maxI = 1000
	x, y := 0., 0.

	for i := 0; (x*x+y*y < complexity) && i < maxI; i++ {
		x, y = x*x-y*y+xi, 2*x*y+yi
	}

	return color.Gray{uint8(x)}
}

func norm(x, total int, min, max float64) float64 {
	return (max-min)*float64(x)/float64(total) - max
}
