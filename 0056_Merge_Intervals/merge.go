package main

import (
	"fmt"
	"sort"
)

func main() {
	// var intervals = [][]int{{1, 3}, {8, 14}, {2, 6}, {14, 15}}
	// var intervals = [][]int{{1, 4}, {2, 3}}\
	var intervals = [][]int{{}}
	fmt.Println(intervals)
	fmt.Println(merge(intervals))
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var result = [][]int{}
	prev := intervals[0]

	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if prev[1] >= cur[0] {

			prev[1] = max(prev[1], cur[1])
		} else {
			result = append(result, prev)
			prev = cur
		}
	}
	result = append(result, prev)

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
