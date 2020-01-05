package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 6, 1}
	res := countSmaller(nums)
	fmt.Println(res)
}

func countSmaller(nums []int) []int {
	type Node struct {
		left      *Node
		right     *Node
		val       int // 值
		leftChild int // 左子树个数
		dup       int // 重复次数
	}

	// root节点记录其左子树个数,及时计算结果
	res := make([]int, len(nums))
	var root *Node
	var insert func(root *Node, val int, i, preSum int) *Node
	insert = func(root *Node, val int, i, preSum int) *Node {
		if root == nil {
			root = &Node{
				val: val,
				dup: 1,
			}
			res[i] = preSum
			return root
		} else if root.val == val {
			root.dup++
			res[i] = preSum + root.leftChild
		} else if root.val < val {
			root.right = insert(root.right, val, i, preSum+root.leftChild+root.dup)
		} else if root.val > val {
			root.leftChild++
			root.left = insert(root.left, val, i, preSum)
		}

		return root
	}
	for i := len(nums) - 1; i >= 0; i-- {
		root = insert(root, nums[i], i, 0)
	}
	return res
}

/*
func showNode(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(root)
	showNode(root.left)
	showNode(root.right)
}
*/
