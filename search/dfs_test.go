package search

import (
	"fmt"
	"testing"
)

func TestDfs(t *testing.T) {
	fillGraph()
	g.Dfs(func(node *Node) {
		fmt.Printf("DFS visiting... %v\n", node)
	})
}