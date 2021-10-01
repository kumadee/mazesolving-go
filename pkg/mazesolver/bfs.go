package mazesolver

import (
	"container/list"
	"log"
)

func BfsSolve(m *Maze) ([]*Node, error) {
	log.Println("Breadth First Search")
	// run the algo
	q := list.New()
	par := make(map[*Node]*Node)
	q.PushBack(m.start)
	m.start.visited = true
	for q.Len() != 0 {
		e := q.Front()
		v, ok := e.Value.(*Node)
		if !ok {
			return nil, NoSolution
		}
		q.Remove(e)
		if v == m.end {
			return Path(par, m), nil
		}
		for _, edg := range v.AdjacentEdges() {
			if !edg.visited {
				edg.visited = true
				par[edg] = v
				q.PushBack(edg)
			}
		}
	}
	return nil, NoSolution
}

func Path(p map[*Node]*Node, m *Maze) []*Node {
	var sol []*Node
	sol = append(sol, m.end)
	for sol[len(sol)-1] != m.start {
		sol = append(sol, p[sol[len(sol)-1]])
	}
	for _, n := range sol {
		log.Println(n.position)
	}
	for k, v := range p {
		log.Println(k.position, "->", v.position)
	}
	return sol
}
