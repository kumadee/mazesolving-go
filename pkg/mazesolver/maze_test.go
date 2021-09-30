package mazesolver

import "testing"

var testCases = []struct {
	path      string
	nodeCount int
}{
	{
		path:      "../../images/tiny.png",
		nodeCount: 23,
	},
	{
		path:      "../../images/normal.png",
		nodeCount: 325,
	},
}

func TestCreateMaze(t *testing.T) {
	for _, tc := range testCases {
		img := LoadImage(tc.path)
		m := NewMaze(img)
		a := m.CountNodes()
		if a != tc.nodeCount {
			t.Error("Node count mismatch. Actual:", a, "Expected:", tc.nodeCount)
		}
	}
}

func BenchmarkCreateMaze(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			img := LoadImage(tc.path)
			NewMaze(img)
		}
	}
}
