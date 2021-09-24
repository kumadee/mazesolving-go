package mazesolver

import (
	"image"
	"testing"
)

func TestTraverseImage(t *testing.T) {
	for _, tc := range testCases {
		TraverseImage(tc.path)
	}
}

func TestHighlightPathImage(t *testing.T) {
	testCases := []struct {
		path string
		s    []*Node
		out  string
	}{
		{
			path: "../../images/normal.png",
			out:  "../../solutions/normal.png",
			s: []*Node{
				{position: Position{3, 0}},
				{position: Position{3, 1}},
				{position: Position{9, 1}},
				{position: Position{9, 3}},
				{position: Position{9, 9}},
				{position: Position{7, 9}},
				{position: Position{1, 9}},
				{position: Position{1, 11}},
				{position: Position{5, 11}},
				{position: Position{5, 13}},
				{position: Position{3, 13}},
				{position: Position{3, 15}},
				{position: Position{11, 15}},
				{position: Position{11, 13}},
				{position: Position{13, 13}},
				{position: Position{13, 11}},
				{position: Position{17, 11}},
				{position: Position{17, 13}},
				{position: Position{19, 13}},
				{position: Position{25, 13}},
				{position: Position{25, 9}},
				{position: Position{21, 9}},
				{position: Position{21, 7}},
				{position: Position{21, 5}},
				{position: Position{23, 5}},
				{position: Position{23, 3}},
				{position: Position{25, 3}},
				{position: Position{25, 5}},
				{position: Position{27, 5}},
				{position: Position{27, 7}},
				{position: Position{29, 7}},
				{position: Position{29, 9}},
				{position: Position{31, 9}},
				{position: Position{31, 5}},
				{position: Position{29, 5}},
				{position: Position{29, 3}},
				{position: Position{31, 3}},
				{position: Position{31, 1}},
				{position: Position{33, 1}},
				{position: Position{33, 3}},
				{position: Position{35, 3}},
				{position: Position{37, 3}},
				{position: Position{37, 5}},
				{position: Position{35, 5}},
				{position: Position{35, 7}},
				{position: Position{33, 7}},
				{position: Position{33, 9}},
				{position: Position{35, 9}},
				{position: Position{35, 11}},
				{position: Position{37, 11}},
				{position: Position{37, 13}},
				{position: Position{37, 15}},
				{position: Position{33, 15}},
				{position: Position{33, 21}},
				{position: Position{35, 21}},
				{position: Position{35, 19}},
				{position: Position{37, 19}},
				{position: Position{37, 21}},
				{position: Position{39, 21}},
				{position: Position{39, 23}},
			},
		},
	}
	for _, tc := range testCases {
		img := LoadImage(tc.path)
		HighlightPathImage(image.NewRGBA(img.Bounds()), tc.s, tc.out, true)
	}
}
