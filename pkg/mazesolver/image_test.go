package mazesolver

import (
	"os"
	"testing"
)

func TestSaveImage(t *testing.T) {
	testCases := []struct {
		path string
		out  string
	}{
		{
			path: "../../images/normal.png",
			out:  "../../solutions/normal_save.png",
		},
	}
	for _, tc := range testCases {
		img := LoadImage(tc.path)
		SaveImage(img, tc.out)
		f, err := os.Open(tc.out)
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
	}
}

func TestHighlightNodesOnlyImage(t *testing.T) {
	testCases := []struct {
		path string
		out  string
	}{
		{
			path: "../../images/tiny.png",
			out:  "../../solutions/tiny_maze_nodes.png",
		},
		{
			path: "../../images/normal.png",
			out:  "../../solutions/normal_maze_nodes.png",
		},
	}
	for _, tc := range testCases {
		img := LoadImage(tc.path)
		m := NewMaze(img)
		HighlightNodesOnlyImage(img, m, tc.out)
	}
}

var testHighlightPathCases = []struct {
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
			{position: Position{37, 23}},
			{position: Position{37, 25}},
			{position: Position{37, 27}},
			{position: Position{39, 27}},
			{position: Position{39, 35}},
			{position: Position{37, 35}},
			{position: Position{37, 37}},
			{position: Position{37, 39}},
			{position: Position{35, 39}},
			{position: Position{35, 35}},
			{position: Position{31, 35}},
			{position: Position{31, 37}},
			{position: Position{29, 37}},
			{position: Position{29, 35}},
			{position: Position{27, 35}},
			{position: Position{27, 33}},
			{position: Position{27, 31}},
			{position: Position{25, 31}},
			{position: Position{25, 29}},
			{position: Position{23, 29}},
			{position: Position{23, 27}},
			{position: Position{23, 23}},
			{position: Position{15, 23}},
			{position: Position{15, 29}},
			{position: Position{17, 29}},
			{position: Position{17, 33}},
			{position: Position{15, 33}},
			{position: Position{13, 33}},
			{position: Position{13, 35}},
			{position: Position{11, 35}},
			{position: Position{11, 33}},
			{position: Position{9, 33}},
			{position: Position{9, 31}},
			{position: Position{9, 29}},
			{position: Position{5, 29}},
			{position: Position{3, 29}},
			{position: Position{3, 31}},
			{position: Position{3, 33}},
			{position: Position{3, 35}},
			{position: Position{3, 39}},
			{position: Position{5, 39}},
			{position: Position{5, 37}},
			{position: Position{9, 37}},
			{position: Position{9, 39}},
			{position: Position{11, 39}},
			{position: Position{13, 39}},
			{position: Position{15, 39}},
			{position: Position{17, 39}},
			{position: Position{17, 35}},
			{position: Position{19, 35}},
			{position: Position{19, 33}},
			{position: Position{21, 33}},
			{position: Position{21, 35}},
			{position: Position{23, 35}},
			{position: Position{23, 39}},
			{position: Position{25, 39}},
			{position: Position{27, 39}},
			{position: Position{31, 39}},
			{position: Position{31, 40}},
		},
	},
}

func TestHighlightPathImage(t *testing.T) {
	for _, tc := range testHighlightPathCases {
		img := LoadImage(tc.path)
		HighlightPathImage(img, tc.s, tc.out, true)
	}
}

func BenchmarkHighlightPathImage(b *testing.B) {
	for _, tc := range testHighlightPathCases {
		img := LoadImage(tc.path)
		for i := 0; i < b.N; i++ {
			HighlightPathImage(img, tc.s, tc.out, true)
		}
	}
}
