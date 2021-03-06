package mazesolver

import (
	"image"
	"image/color"
)

type Position struct {
	x, y int16
}

type Node struct {
	visited  bool
	position Position
	left     *Node
	right    *Node
	top      *Node
	down     *Node
}

func (n *Node) AdjacentEdges() []*Node {
	var e []*Node
	if n.left != nil {
		e = append(e, n.left)
	}
	if n.right != nil {
		e = append(e, n.right)
	}
	if n.top != nil {
		e = append(e, n.top)
	}
	if n.down != nil {
		e = append(e, n.down)
	}
	return e
}

type Maze struct {
	nodes [][]*Node
	start *Node
	end   *Node
}

// Converts the input image to a maze data structure
func NewMaze(img image.Image) *Maze {
	rec := img.Bounds()
	m := Maze{}
	m.nodes = make([][]*Node, rec.Dy())
	for y := range m.nodes {
		m.nodes[y] = make([]*Node, rec.Dx())
	inner:
		for x := range m.nodes[y] {
			path := !isWall(img.At(x, y))
			if !path {
				// This is a wall
				continue
			}
			p := Position{int16(x), int16(y)}
			switch {
			case y == 0:
				m.start = &Node{position: p}
				m.nodes[y][x] = m.start
				break inner
			case y == (rec.Dy() - 1):
				m.end = &Node{position: p}
				addNeighbours(m.end, &m, true, false)
				m.nodes[y][x] = m.end
				break inner
			default:
				// Check if new node has to be added in maze
				addNewNode(p, img, &m)
			}
		}
	}
	return &m
}

func (m *Maze) CountNodes() int {
	c := 0
	for _, row := range m.nodes {
		for _, n := range row {
			if n != nil {
				c++
			}
		}
	}
	return c
}

func isWall(c color.Color) bool {
	r, g, b, _ := c.RGBA()
	return r|g|b == 0
}

func addNeighbours(n *Node, m *Maze, l bool, t bool) {
	if !l {
		for x := n.position.x - 1; x >= 0; x-- {
			if m.nodes[n.position.y][x] != nil {
				n.left = m.nodes[n.position.y][x]
				m.nodes[n.position.y][x].right = n
				break
			}
		}
	}
	if !t {
		for y := n.position.y - 1; y >= 0; y-- {
			if m.nodes[y][n.position.x] != nil {
				n.top = m.nodes[y][n.position.x]
				m.nodes[y][n.position.x].down = n
				break
			}
		}
	}
}

func addNewNode(p Position, img image.Image, m *Maze) {
	x := int(p.x)
	y := int(p.y)
	l := isWall(img.At(x-1, y))
	r := isWall(img.At(x+1, y))
	t := isWall(img.At(x, y-1))
	d := isWall(img.At(x, y+1))

	validNode := (l && t && !r && !d) || (l && !t && !r && d) || (!l && !t && !r) || (!l && !t && r && d) || (!l && t && !r && !d) || (!l && t && r && !d) || (!l && !t && r && !d) || (l && !t && !r && !d)
	deadEndNode := (!l && t && r && d) || (l && !t && r && d) || (l && t && !r && d) || (l && t && r && !d)

	if !(validNode || deadEndNode) {
		return
	}
	m.nodes[p.y][p.x] = &Node{position: Position{p.x, p.y}}
	addNeighbours(m.nodes[p.y][p.x], m, l, t)

	if deadEndNode {
		// Marking them as already visited as they are deadends
		m.nodes[p.y][p.x].visited = true
	}
}
