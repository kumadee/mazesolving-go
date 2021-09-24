package mazesolver

type Position struct {
	x, y int
}

type Node struct {
	isWall   bool
	position Position
	left     *Node
	right    *Node
	up       *Node
	down     *Node
}

func (n *Node) New(p *Position) {
	n.position.x = p.x
	n.position.y = p.y
}

type Maze [][]Node

// Converts the input image to a maze data structure
func CreateMaze(imgPath string) *Maze {
	img := LoadImage(imgPath)
	bounds := img.Bounds()
	m := make(Maze, bounds.Dy())
	for y := range m {
		m[y] = make([]Node, bounds.Dx())
		for x, node := range m[y] {
			if r, g, b, _ := img.At(x, y).RGBA(); r&g&b == 0 {
				node.isWall = true
			}
		}
	}
	return &m
}
