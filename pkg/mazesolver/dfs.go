package mazesolver

import (
	"container/list"
	"log"
)

func DfsSolve(m *Maze) ([]*Node, error) {
	log.Println("Depth First Search")
	stk := list.New()
	par := make(map[Position]Position)
	stk.PushBack(m.start)
	m.start.visited = true
	for stk.Len() != 0 {
		e := stk.Back()
		v, ok := e.Value.(*Node)
		if !ok {
			return nil, NoSolution
		}
		stk.Remove(e)
		if v == m.end {
			return Path(par, m), nil
		}
		for _, edg := range v.AdjacentEdges() {
			if !edg.visited {
				edg.visited = true
				par[edg.position] = v.position
				stk.PushBack(edg)
			}
		}
	}
	return nil, NoSolution
}
