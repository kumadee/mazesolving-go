package mazesolver

import "testing"

var testMazeCases = []struct {
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

func TestNewMaze(t *testing.T) {
	for _, tc := range testMazeCases {
		img := LoadImage(tc.path)
		m := NewMaze(img)
		a := m.CountNodes()
		if a != tc.nodeCount {
			t.Error("Node count mismatch. Actual:", a, "Expected:", tc.nodeCount)
		}
	}
}

func benchmarkNewMaze(i int, b *testing.B) {
	tc := testMazeCases[i]
	img := LoadImage(tc.path)
	for i := 0; i < b.N; i++ {
		NewMaze(img)
	}
}

func BenchmarkTinyNewMaze(b *testing.B) {
	benchmarkNewMaze(0, b)
}

func BenchmarkNormalNewMaze(b *testing.B) {
	benchmarkNewMaze(1, b)
}
