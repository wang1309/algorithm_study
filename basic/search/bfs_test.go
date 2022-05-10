package search

import (
	"fmt"
	"testing"
)

var g Graph

func fillGraph() {
	nA := Node{"A"}
	nB := Node{"B"}
	nC := Node{"C"}
	nD := Node{"D"}
	nE := Node{"E"}
	nF := Node{"F"}

	g.AddNodes(&nA, &nB, &nC, &nD, &nE, &nF)

	g.AddEdge(nA, nB)
	g.AddEdge(nA, nC)
	g.AddEdge(nB, nE)
	g.AddEdge(nC, nE)
	g.AddEdge(nE, nF)
	g.AddEdge(nD, nA)
}

func TestBfs(t *testing.T) {
	fillGraph()
	g.Bfs(func(node *Node) {
		fmt.Printf("DFS visiting... %v\n", node)
	})
}