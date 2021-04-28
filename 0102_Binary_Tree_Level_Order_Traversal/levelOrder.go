package main

import (
	"fmt"

	"github.com/aschenmaker/leetcode/utils"
)

func main() {
	list := []int{3, 9, 20, 0, 0, 15, 7}
	tr := utils.MakeTree(list, 0, len(list))

	fmt.Println(levelOrder(tr))
}

func levelOrder(root *utils.TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var queue = []*utils.TreeNode{}
	queue = append(queue, root)

	ans := [][]int{}

	for len(queue) > 0 {
		length := len(queue)
		var layer = []int{}

		for length != 0 {
			element := queue[0]
			queue = queue[1:]

			layer = append(layer, element.Val)

			if element.Left != nil {
				queue = append(queue, element.Left)
			}
			if element.Right != nil {
				queue = append(queue, element.Right)
			}
			length--
		}
		ans = append(ans, layer)
	}
	return ans
}
