package mazesolver

import (
	"fmt"
	"image"
	"os"

	// Package image/jpeg is not used explicitly in the code below,
	// but is imported for its initialization side-effect, which allows
	// image.Decode to understand JPEG formatted images. Uncomment these
	// two lines to also understand GIF and PNG images:
	// _ "image/gif"
	"image/color"
	"image/png"
	_ "image/png"
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

func HighlightPathImage(img *image.RGBA, s []*Node, path string, highlightNode bool) {
	var prev *Node
	blue := color.RGBA{0, 0, 255, 0xff}
	yellow := color.RGBA{0, 255, 255, 0xff}
	c := blue
	if highlightNode {
		c = yellow
	}
	for _, n := range s {
		img.SetRGBA(n.position.x, n.position.y, c)
		if prev == nil {
			prev = n
			continue
		}
		if prev.position.x == n.position.x {
			for i := min(prev.position.y, n.position.y) + 1; i < max(prev.position.y, n.position.y); i++ {
				img.SetRGBA(n.position.x, i, blue)
			}
			continue
		}
		if prev.position.y == n.position.y {
			for i := min(prev.position.x, n.position.x) + 1; i < max(prev.position.x, n.position.x); i++ {
				img.SetRGBA(i, n.position.y, blue)
			}
		}
	}
	f, _ := os.Create(path)
	CheckError(png.Encode(f, img))
}
