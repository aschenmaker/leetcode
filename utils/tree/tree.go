package tree

import "fmt"

type TreeNode struct {
	Data  string
	Left  *TreeNode
	Right *TreeNode
}

// Print
func (t *TreeNode) Print() {
	fmt.Print(t.Data, " ")
}

// PreOrder
func (t *TreeNode) PreOrder() {
	if t == nil {
		return
	}
	t.Print()
	// print left and right
	t.Left.PreOrder()
	t.Right.PreOrder()
}

// MidOrder
func (t *TreeNode) MidOrder() {
	if t == nil {
		return
	}

	t.Left.MidOrder()
	t.Print()
	// print left and right
	t.Right.MidOrder()
}

// PostOrder
func (t *TreeNode) PostOrder() {
	if t == nil {
		return
	}

	t.Left.PostOrder()
	t.Print()
	// print left and right
	t.Right.PostOrder()
}
