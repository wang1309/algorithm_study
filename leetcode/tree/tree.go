package tree

import "fmt"

/*
* 遇到一道二叉树的题目时的通用思考过程是：

* 1、是否可以通过遍历一遍二叉树得到答案？如果可以，用一个 traverse 函数配合外部变量来实现。

* 2、是否可以定义一个递归函数，通过子问题（子树）的答案推导出原问题的答案？如果可以，写出这个递归函数的定义，并充分利用这个函数的返回值。

* 3、无论使用哪一种思维模式，你都要明白二叉树的每一个节点需要做什么，需要在什么时候（前中后序）做。
 */

func init() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 生成二叉树
func dfs(p *TreeNode, depth int) {
	if depth < 3 {
		left := &TreeNode{Val: 2 * depth}
		right := &TreeNode{Val: 4 * depth}
		p.Left = left
		p.Right = right
		dfs(p.Left, depth+1)
		dfs(p.Right, depth+1)
	}
}

// 前序遍历
func PreorderTraversal(node *TreeNode) {
	if node == nil {
		return
	}

	fmt.Println(node.Val)
	PreorderTraversal(node.Left)
	PreorderTraversal(node.Right)
}

// 中序遍历
func MiddleorderTraversal(node *TreeNode) {
	if node == nil {
		return
	}

	MiddleorderTraversal(node.Left)
	fmt.Println(node.Val)
	MiddleorderTraversal(node.Right)
}

// 后序遍历
func PostorderTraversal(node *TreeNode) {
	if node == nil {
		return
	}

	PostorderTraversal(node.Left)
	PostorderTraversal(node.Right)
	fmt.Println(node.Val)
}

// 二叉树层序遍历
func levelOrder(node *TreeNode) {
	if node == nil {
		return
	}

	var nodeSlice []*TreeNode

	nodeSlice = append(nodeSlice, node)

	RecursionTraversal(nodeSlice)

}

func RecursionTraversal(nodeSlice []*TreeNode) {
	if len(nodeSlice) == 0 {
		return
	}

	var nexSlice []*TreeNode

	for i := 0; i < len(nodeSlice); i++ {
		node := nodeSlice[i]

		fmt.Println(node.Val)

		if node.Left != nil {
			nexSlice = append(nexSlice, node.Left)
		}

		if node.Right != nil {
			nexSlice = append(nexSlice, node.Right)
		}
	}

	RecursionTraversal(nexSlice)
}
