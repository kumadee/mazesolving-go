package mazesolver

import "testing"

func TestSolve(t *testing.T) {
	testCases := []struct {
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
	}
	for _, tc := range testCases {
		tc.opt.Solve()
	}
}
