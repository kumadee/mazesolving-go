package mazesolver

import "testing"

var testSolveCases = []struct {
	opt Options
}{
	{
		opt: Options{
			Method:         Bfs,
			InImg:          "../../images/tiny.png",
			OutImg:         "../../solutions/tiny.png",
			HighlightNodes: true,
		},
	},
	{
		opt: Options{
			Method:         Bfs,
			InImg:          "../../images/normal.png",
			OutImg:         "../../solutions/normal.png",
			HighlightNodes: true,
		},
	},
	{
		opt: Options{
			Method:         Dfs,
			InImg:          "../../images/tiny.png",
			OutImg:         "../../solutions/tiny_dfs.png",
			HighlightNodes: true,
		},
	},
	{
		opt: Options{
			Method:         Dfs,
			InImg:          "../../images/normal.png",
			OutImg:         "../../solutions/normal_dfs.png",
			HighlightNodes: true,
		},
	},
}

func TestSolve(t *testing.T) {
	for _, tc := range testSolveCases {
		tc.opt.Solve()
	}
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < 1; i++ {
		for _, tc := range testSolveCases {
			tc.opt.Solve()
		}
	}
}
