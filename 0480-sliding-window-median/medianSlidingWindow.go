package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3, 4, 2, 3, 1, 4, 2}
	k := 3
	fmt.Println(medianSlidingWindow(nums, k))
}

func medianSlidingWindow(nums []int, k int) []float64 {
	if k > len(nums) {
		return nil
	}
	if k == 0 {
		return nil
	}
	res := []float64{}

	if k == 1 {
		for _, v := range nums {
			res = append(res, float64(v))
		}
		return res
	}

	windows := make([]int, k)
	copy(windows, nums[:k])

	sort.Ints(windows)

	res = append(res, getMid(windows, k))

	for i := k; i < len(nums); i++ {
		id := sort.SearchInts(windows, nums[i-k])
		windows[id] = nums[i]

		// 冒泡排序
		for id < k-1 && windows[id] > windows[id+1] {
			windows[id], windows[id+1] = windows[id+1], windows[id]
			id++
		}
		// 向下冒泡
		for id > 0 && windows[id] < windows[id-1] {
			windows[id], windows[id-1] = windows[id-1], windows[id]
			id--
		}

		res = append(res, getMid(windows, k))
	}
	return res
}

func getMid(windows []int, k int) float64 {
	res := (float64(windows[k/2]) + float64(windows[(k-1)/2])) / 2
	return res
}
