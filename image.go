package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

// WriteImage out a framebuffer to a PPM image file.
func WriteImage(filename string, options Options, framebuffer []Vec3f) {
	f, err := os.OpenFile(
		filename,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	header := createHeader(options.width, options.height)

	f.WriteString(header)

	for _, pixel := range framebuffer {
		r, g, b := createColor(pixel)
		f.Write([]byte{r, g, b})
	}
}

func createHeader(width, height int) string {
	var header strings.Builder

	sWidth := strconv.Itoa(width)
	sHeight := strconv.Itoa(height)

	header.WriteString("P6\n")
	header.WriteString(sWidth)
	header.WriteString(" ")
	header.WriteString(sHeight)
	header.WriteString("\n255\n")

	return header.String()
}

func createColor(pixel Vec3f) (r, g, b byte) {
	r = byte(255 * Clamp(0, 1, pixel.x))
	g = byte(255 * Clamp(0, 1, pixel.y))
	b = byte(255 * Clamp(0, 1, pixel.z))

	return
}
