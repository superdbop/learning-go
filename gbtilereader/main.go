package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()

	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	const width, height = 8, 8

	palette := []color.Color{
		color.RGBA{255, 255, 255, 255}, // White
		color.RGBA{191, 191, 191, 255}, // Light Gray
		color.RGBA{127, 127, 127, 255}, // Dark Gray
		color.RGBA{0, 0, 0, 255},       // Black
	}

	// Create a colored image of the given width and height.
	rect := image.Rect(0, 0, width, height)
	img := image.NewPaletted(rect, palette)

	for y := 0; y < height; y++ {
		b1 := data[y*2]
		b2 := data[(y*2)+1]
		for x := 0; x < width; x++ {
			bit1 := (b1 >> uint8(x)) & 0x01
			bit2 := (b2 >> uint8(x)) & 0x01
			p := bit1 | (bit2 << 1)

			img.SetColorIndex(x, y, p)
		}
	}

	f, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Time: %v\n", time.Since(start))
}
