package search

type NodeStack struct {
	Items []Node
}

func NewNodeStack() *NodeStack {
	return &NodeStack{
		Items: make([]Node, 0),
	}
}

func (n *NodeStack) push(q Node) {
	n.Items = append(n.Items, q)
}

func (n *NodeStack) pop() *Node {

	item := n.Items[len(n.Items)-1]

	n.Items = n.Items[0 : len(n.Items)-1]
	return &item
}

func (n *NodeStack) IsEmpty() bool {
	return len(n.Items) == 0
}

func (n *NodeStack) Size() int {
	return len(n.Items)
}

func (g *Graph) Dfs(f func(node *Node)) {
	stack := NewNodeStack()

	n := g.Nodes[0]
	stack.push(*n)

	visited := make(map[*Node]bool)
	visited[n] = true

	for {
		if stack.IsEmpty() {
			break
		}

		node := stack.pop()
		if !visited[node] {
			visited[node] = true
		}

		near := g.Edges[*node]
		for i := 0; i < len(near); i++ {
			if !visited[near[i]] {
				visited[near[i]] = true
				stack.push(*near[i])
			}
		}

		if f != nil{
			f(node)
		}
	}
}
