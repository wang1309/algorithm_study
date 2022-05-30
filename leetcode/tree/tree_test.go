package tree

import "testing"

func GenerateTree() *TreeNode {
	node := &TreeNode{
		Val: 1,
		Left: nil,
		Right: nil,
	}

	return node
}


//根节点
var rootNode *TreeNode = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 4,
			Right: &TreeNode{
				Val:   6,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 8},
			},
		},
	},
	Right: &TreeNode{
		Val:   3,
		Right: &TreeNode{Val: 5},
	},
}

// 前序遍历
func TestPreorderTraversal(t *testing.T)  {
	//root := GenerateTree()
	PreorderTraversal(rootNode)
}

// 中序遍历
func TestMiddleorderTraversal(t *testing.T)  {
	//root := GenerateTree()
	MiddleorderTraversal(rootNode)
}

// 后序遍历
func TestPostorderTraversal(t *testing.T)  {
	//root := GenerateTree()
	PostorderTraversal(rootNode)
}

// 层序遍历
func TestLevelOrder(t *testing.T)  {
	levelOrder(rootNode)
}