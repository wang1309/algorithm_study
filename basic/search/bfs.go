package search

type NodeQueue struct {
	Items []Node
}

func NewQueue() *NodeQueue {
	return &NodeQueue{
		Items: []Node{},
	}
}

func (n *NodeQueue) Enqueue(node Node) {
	n.Items = append(n.Items, node)
}

func (n *NodeQueue) Dequeue() *Node {
	if len(n.Items) <= 0 {
		return nil
	}

	node := n.Items[0]
	n.Items = n.Items[1:]

	return &node
}

func (n *NodeQueue) IsEmpty() bool {
	return len(n.Items) == 0
}

// Bfs 广度优先遍历
func (g *Graph) Bfs(f func(node *Node)) {
	if len(g.Nodes) <= 0 {
		return
	}

	queue := NewQueue()

	n := g.Nodes[0]
	queue.Enqueue(*n)

	visited := make(map[*Node]bool)
	visited[n] = true

	for {
		if queue.IsEmpty() {
			break
		}

		node := queue.Dequeue()
		near := g.Edges[*node]

		for i := 0; i < len(near); i++ {
			if !visited[near[i]] {
				queue.Enqueue(*near[i])
				visited[near[i]] = true
			}
		}

		if f != nil {
			f(node)
		}
	}
}
