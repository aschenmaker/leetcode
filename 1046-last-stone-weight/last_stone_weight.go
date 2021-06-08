package laststoneweight

import (
	"container/heap"
	"sort"
)

// 维护一个始终保持大小顺序的数据结构，
// 取最大值
// 使用 最大堆 来完成这道题目
// 将所给数组初始化为堆

func lastStoneWeight(stones []int) int {
	q := &hp{stones}
	heap.Init(q)
	for q.Len() > 1 {
		x, y := q.pop(), q.pop()
		if x > y {
			q.push(x - y)
		}
	}
	if q.Len() > 0 {
		return q.IntSlice[0]
	}
	return 0
}

type hp struct {
	sort.IntSlice
}

// sortslice 已经实现了Len Less Swap
func (h hp) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

// 实现堆的接口
func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}
func (h *hp) Pop() interface{} {
	v := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return v
}

func (h *hp) push(v int) {
	heap.Push(h, v)
}
func (h *hp) pop() int {
	return heap.Pop(h).(int)
}
