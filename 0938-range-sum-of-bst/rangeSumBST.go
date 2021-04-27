package main

import (
	"fmt"

	"github.com/aschenmaker/leetcode/utils"
)

func main() {
	// var list = []int{10, 5, 15, 3, 7, 0, 18}
	var list = []int{10, 5, 15, 3, 7, 13, 18, 1, 0, 6}
	tree := utils.MakeTree(list, 0, len(list))

	fmt.Println(rangeSumBST(tree, 6, 10))
}

func rangeSumBST(root *utils.TreeNode, low int, high int) int {
	// 中序遍历 在范围内
	if root == nil {
		return 0
	}
	sum := 0

	var midOrder func(*utils.TreeNode)
	midOrder = func(root *utils.TreeNode) {
		if root == nil {
			return
		}

		midOrder(root.Left)
		fmt.Println(root.Val)
		if root.Val >= low && root.Val <= high {
			sum += root.Val
		}
		if root.Val > high {
			return
		}
		midOrder(root.Right)
	}

	midOrder(root)

	return sum
}

func byDFS(root *utils.TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	if root.Val > high {
		return byDFS(root.Left, low, high)
	}

	if root.Val < low {
		return byDFS(root.Right, low, high)
	}

	return root.Val + byDFS(root.Left, low, high) + byDFS(root.Right, low, high)
}

func byBFS(root *utils.TreeNode, low int, high int) (sum int) {
	q := []*utils.TreeNode{root}

	for len(q) > 0 {
		node := q[0]
		q = q[1:]

		if node == nil {
			continue
		}
		if node.Val > high {
			q = append(q, node.Left)
		} else if node.Val < low {
			q = append(q, node.Right)
		} else {
			sum += node.Val
			q = append(q, node.Left, node.Right)
		}

	}
	return
}
