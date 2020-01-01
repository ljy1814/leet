package main

func main() {
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	m := 0
	maxPathSum1(root, &m)
	return m
}

func maxPathSum1(root *TreeNode, m *int) int {
	if root == nil {
		return 0
	}

	ml := maxPathSum1(root.Left, m)
	mr := maxPathSum1(root.Right, m)

	if ml < 0 {
		ml = 0
	}
	if mr < 0 {
		mr = 0
	}

	t := root.Val + ml + mr
	if t > *m {
		*m = t
	}
	if ml > mr {
		return root.Val + ml
	}
	return root.Val + mr
}
