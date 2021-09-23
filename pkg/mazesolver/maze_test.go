package mazesolver

import "testing"

var testCases = []struct {
	path string
}{
	{
		path: "../../images/tiny.png",
	},
	{
		path: "../../images/normal.png",
	},
}

func TestCreateMaze(t *testing.T) {
	for _, tc := range testCases {
		CreateMaze(tc.path)
	}
}

func BenchmarkCreateMaze(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			CreateMaze(tc.path)
		}
	}
}