package tree

func MakeTree() *TreeNode {
	t := &TreeNode{Data: "A"}
	t.Left = &TreeNode{Data: "B"}
	t.Right = &TreeNode{Data: "C"}
	t.Left.Left = &TreeNode{Data: "D"}
	t.Left.Right = &TreeNode{Data: "E"}
	t.Right.Left = &TreeNode{Data: "F"}
	return t
}
