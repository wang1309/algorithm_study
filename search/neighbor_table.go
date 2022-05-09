package search

import "fmt"

type Node struct {
	Value interface{}
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.Value)
}

type Graph struct {
	Nodes []*Node
	Edges map[Node][]*Node
}

// AddNodes 添加节点
func (g *Graph) AddNodes(nodes ...*Node) {
	g.Nodes = append(g.Nodes, nodes...)
}

// AddEdge 添加边
func (g *Graph) AddEdge(node1, node2 Node) {
	// 懒加载初始化
	if g.Edges == nil {
		g.Edges = make(map[Node][]*Node)
	}

	// 无向图
	// 设置 node1 指向 node2 的边
	g.Edges[node1] = append(g.Edges[node1], &node2)
	// 设置 node2 指向 node2 的边
	g.Edges[node2] = append(g.Edges[node2], &node1)
}

func (g *Graph) String() {
	s := ""

	for i := 0; i < len(g.Nodes); i++ {
		s += g.Nodes[i].String() + "->"
		near := g.Edges[*g.Nodes[i]]

		for j := 0; j < len(near); j++ {
			s += near[j].String() + " "
		}
		s += "\n"
	}

	fmt.Println(s)
}
