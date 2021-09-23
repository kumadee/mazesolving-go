package mazesolver

const (
	Bfs = iota
	Dfs
)

type Options struct {
	method int
	inImg  string
	outImg string
}

func (opt *Options) Solve() {
	m := map[int]func(inImg string, outImg string){
		Bfs: BfsSolve,
		Dfs: DfsSolve,
	}
	// make maze data structure out of the inImg
	m[opt.method](opt.inImg, opt.outImg)
	// make outImg out of maze data structure
}
