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

func HighlightNodesOnlyImage(i image.Image, m *Maze, path string) {
	im, ok := i.(*image.Paletted)
	if !ok {
		log.Fatal("Unable to convert image to image.Paletted.")
	}
	img := ConvertToRGBA(im)
	c := color.RGBA{255, 0, 0, 255}
	for _, row := range m.nodes {
		for _, n := range row {
			if n != nil {
				img.Set(int(n.position.x), int(n.position.y), c)
			}
		}
	}
	SaveImage(img, path)
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
		nX := int(n.position.x)
		nY := int(n.position.y)
		img.Set(nX, nY, p)
		switch {
		case prev == nil:
		case prev.position.x == n.position.x:
			pY := int(prev.position.y)
			for yi := min(pY, nY) + 1; yi < max(pY, nY); yi++ {
				img.Set(nX, yi, c)
			}
		case prev.position.y == n.position.y:
			pX := int(prev.position.x)
			for xi := min(pX, nX) + 1; xi < max(pX, nX); xi++ {
				img.Set(xi, nY, c)
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
