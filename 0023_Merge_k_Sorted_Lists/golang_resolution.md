
```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
    if 0==len(lists) {return nil}
    return merge(lists,0,len(lists)-1)
}
func merge(lists []*ListNode, left int,right int) *ListNode{
    //分治框架
    if left==right {return lists[left]}
    mid:=(left+right)>>1
    left_list:= merge(lists,left,mid)   //merge左
    right_list:=merge(lists,mid+1,right)  //merge右
    return merge2Lists(left_list,right_list) //左右合并
   
}

func merge2Lists(l1,l2 *ListNode) *ListNode{
    //合并两个链表
    if nil==l1 || nil==l2 {
        if nil != l1 {return l1}
        if nil != l2 {return l2}
        return nil
    }//先判空
    head:=&ListNode{-1,nil}
    cur:=head
    for nil!=l1 && nil!=l2 {
        if l1.Val<=l2.Val {
            cur.Next = l1
            l1=l1.Next
        }else{
            cur.Next = l2
            l2=l2.Next
        }
        cur=cur.Next
    }
    if nil!=l1 {cur.Next=l1}
    if nil!=l2 {cur.Next=l2}
    return head.Next
}
```

