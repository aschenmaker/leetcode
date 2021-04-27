package utils

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MakeTree(list []int, i, l int) *TreeNode {
	if i >= l || list[i] == 0 {
		return nil
	}

	root := TreeNode{Val: list[i]}

	root.Left = MakeTree(list, 2*i+1, l)
	root.Right = MakeTree(list, 2*i+2, l)

	return &root
}
