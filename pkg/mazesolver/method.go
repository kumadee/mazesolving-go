package mazesolver

import (
	"errors"
)

const (
	Bfs = iota
	Dfs
)

var (
	NoSolution = errors.New("No solution found for the given maze")
)

type Options struct {
	method         int
	inImg          string
	outImg         string
	highlightNodes bool
}

func (opt *Options) Solve() {
	m := map[int]func(*Maze) ([]*Node, error){
		Bfs: BfsSolve,
		Dfs: DfsSolve,
	}
	// make maze data structure out of the inImg
	img := LoadImage(opt.inImg)
	mz := NewMaze(img)
	// find solution for the maze
	sol, err := m[opt.method](mz)
	CheckError(err)
	// make outImg out of maze data structure
	if sol != nil {
		HighlightPathImage(img, sol, opt.outImg, opt.highlightNodes)
	}
}
