package mazesolver

import (
	"image"
	"image/color"
)

type Position struct {
	x, y int
}

type Node struct {
	visited  bool
	position Position
	left     *Node
	right    *Node
	top      *Node
	down     *Node
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
			switch {
			case !path:
				// This is a wall
				continue
			case path && y == 0:
				m.start = &Node{position: Position{x, y}}
				m.nodes[y][x] = m.start
				break inner
			case path && y == (rec.Dy()-1):
				m.end = &Node{position: Position{x, y}}
				m.nodes[y][x] = m.end
				break inner
			case path:
				// Check if new node has to be added in maze
				addNewNode(Position{x, y}, img, &m)
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
	l := isWall(img.At(p.x-1, p.y))
	r := isWall(img.At(p.x+1, p.y))
	t := isWall(img.At(p.x, p.y-1))
	d := isWall(img.At(p.x, p.y+1))
	// left != wall & top == wall & right == wall & down == wall: (deadend node)
	// left == wall & top == wall & right != wall & down == wall: (deadend node)
	// left == wall & top != wall & right == wall & down == wall: (deadend node)
	// left == wall & top == wall & right == wall & down != wall: (deadend node)
	if (l && t && !r && !d) || (l && !t && !r && d) || (!l && !t && !r) || (!l && !t && r && d) || (!l && t && !r && !d) || (!l && t && r && !d) || (!l && !t && r && !d) || (l && !t && !r && !d) {
		m.nodes[p.y][p.x] = &Node{position: Position{p.x, p.y}}
		addNeighbours(m.nodes[p.y][p.x], m, l, t)
	}
}
