/*
	*
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

package main	
 
import "fmt"

type TreeNode struct {
	Val int
	Left, Right *TreeNode
}
var res []int

func rightSideView(root *TreeNode) []int {
	res = []int{}
	dfs(root, 1)
	fmt.Println("1")
	return res
}

func dfs(root *TreeNode, level int)  {
	if root == nil {
		return 
	}
	if(level>len(res)){
		res = append(res,root.Val)
	}
	dfs(root.Right,level+1)
	dfs(root.Left,level+1)
}

func main()  {
	root := TreeNode{Val: 3}
	rightSideView(root.Left)
}