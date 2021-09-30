package mazesolver

import (
	"container/list"
	"log"
)

func BfsSolve(m *Maze) ([]*Node, error) {
	log.Println("Breadth First Search")
	// run the algo
	var q *list.List
	q.PushBack(m.start)
	m.start.visited = true
	for q.Len() != 0 {
		e := q.Back()
		v, ok := e.Value.(*Node)
		if !ok {
			return nil, NoSolution
		}
		q.Remove(e)
		if v == m.end {
			return Path(), nil
		}
		if v.left != nil && !v.left.visited {
			v.left.visited = true
			q.PushFront(v.left)
		}
		if v.right != nil && !v.right.visited {
			v.right.visited = true
			q.PushFront(v.right)
		}
		if v.top != nil && !v.top.visited {
			v.top.visited = true
			q.PushFront(v.top)
		}
		if v.down != nil && !v.down.visited {
			v.down.visited = true
			q.PushFront(v.down)
		}
	}
	return nil, NoSolution
}

func Path() []*Node {
	var n []*Node
	return n
}
