package main

import (
	"fmt"
	"sort"
)

func main() {
	weights := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	D := 5

	fmt.Println(shipWithinDays(weights, D))
}

func shipWithinDays(weights []int, D int) int {
	left, right := 0, 0
	// 先计算左右边界
	// 左边界为单个最小货物的值
	// 右边界为所有货物的重量
	for _, w := range weights {
		if w > left {
			left = w
		}
		right += w
	}
	// right-left 优化数值大小，进行放缩
	ans := left + sort.Search(right-left, func(i int) bool {
		// fmt.Println("current ans:=", left+i)
		i += left // 从下边界进行查找
		day := 1  // 从第一日计算
		sum := 0  //
		//如果已经运送货物量 大于了当前货运吨数，
		//则天数需要增加
		for _, w := range weights {
			if sum+w > i {
				day++
				sum = 0
			}
			sum += w
		}
		return day <= D
	})
	return ans
}
