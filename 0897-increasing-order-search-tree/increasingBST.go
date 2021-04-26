package main

func main() {

}

func increasingBST(root *TreeNode) *TreeNode {
	vl := []int{}
	var midOrder = func(*TreeNode) {}
	midOrder = func(node *TreeNode) {
		if node != nil {
			midOrder(node.Left)
			vl = append(vl, node.Val)
			midOrder(node.Right)
		}
	}
	midOrder(root)

	newTreeRoot := &TreeNode{}
	curNode := newTreeRoot

	for _, v := range vl {
		curNode.Right = &TreeNode{Val: v}
		curNode = curNode.Right
	}

	return newTreeRoot.Right
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST2(root *TreeNode) *TreeNode {
	var newRoot = &TreeNode{}
	resNode := newRoot

	var midOrder func(*TreeNode)
	midOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		midOrder(node.Left)
		resNode.Right = node
		node.Left = nil
		resNode = node

		midOrder(node.Right)
	}
	midOrder(root)
	return newRoot.Right
}
