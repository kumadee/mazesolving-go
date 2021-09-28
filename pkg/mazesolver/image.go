package mazesolver

import (
	"fmt"
	"image"
	"log"
	"os"

	// Package image/jpeg is not used explicitly in the code below,
	// but is imported for its initialization side-effect, which allows
	// image.Decode to understand JPEG formatted images. Uncomment these
	// two lines to also understand GIF and PNG images:
	// _ "image/gif"
	"image/color"
	"image/png"
	// _ "image/jpeg"
)

func LoadImage(path string) image.Image {
	reader, err := os.Open(path)
	CheckError(err)
	defer reader.Close()
	img, _, err := image.Decode(reader)
	CheckError(err)
	return img
}

func ConvertToRGBA(i image.Image) *image.RGBA {
	b := i.Bounds()
	m := image.NewRGBA(b)
	for y := 0; y < b.Dy(); y++ {
		for x := 0; x < b.Dx(); x++ {
			m.Set(x, y, i.At(x, y))
		}
	}
	return m
}

func TraverseImage(path string) {
	img := LoadImage(path)
	b := img.Bounds()
	for y := 0; y < b.Dy(); y++ {
		for x := 0; x < b.Dx(); x++ {
			r, g, b, a := img.At(x, y).RGBA()
			if r&g&b&a == 0 {
				fmt.Println("Wall: ", x, y)
			}
		}
	}
}

func HighlightPathImage(m image.Image, s []*Node, path string, highlightNode bool) {
	var prev *Node
	im, ok := m.(*image.Paletted)
	if !ok {
		log.Fatal("Unable to convert image to image.Paletted.")
	}
	img := ConvertToRGBA(im)
	var p color.RGBA
	for i, n := range s {
		r := uint8((i * 255) / len(s))
		c := color.RGBA{r, 0, 255 - r, 255}
		if highlightNode {
			p = color.RGBA{255 - r, 255, r, 255}
		} else {
			p = c
		}
		img.Set(n.position.x, n.position.y, p)
		switch {
		case prev == nil:
		case prev.position.x == n.position.x:
			for yi := min(prev.position.y, n.position.y) + 1; yi < max(prev.position.y, n.position.y); yi++ {
				img.Set(n.position.x, yi, c)
			}
		case prev.position.y == n.position.y:
			for xi := min(prev.position.x, n.position.x) + 1; xi < max(prev.position.x, n.position.x); xi++ {
				img.Set(xi, n.position.y, c)
			}
		}
		prev = n
	}
	SaveImage(img, path)
}

func SaveImage(img image.Image, path string) {
	f, _ := os.Create(path)
	CheckError(png.Encode(f, img))
	defer f.Close()
}
