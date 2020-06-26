package main

import "fmt"

// ListNode is listnode
// 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	sli := []int{1, 2, 3, 3, 2, 1, 4}
	head := &ListNode{Val: 1}
	pointer := head
	for i := 1; i < len(sli); i++ {
		cur := &ListNode{Val: sli[i]}
		pointer.Next = cur
		pointer = pointer.Next
	}
	ans := removeDuplicateNodes(head)
	for ans != nil {
		fmt.Println(ans.Val)
		ans = ans.Next
	}
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	hasNode := map[int]bool{head.Val: true}
	pos := head
	for pos.Next != nil {
		cur := pos.Next
		if !hasNode[cur.Val] {
			hasNode[cur.Val] = true
			pos = pos.Next
		} else {
			pos.Next = pos.Next.Next
		}
	}
	// pos.Next = nil
	return head
}
