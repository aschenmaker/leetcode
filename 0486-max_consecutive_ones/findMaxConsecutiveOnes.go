package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 1, 0}
	fmt.Println(findMaxConsecutiveOnes(nums))
}

func findMaxConsecutiveOnes(nums []int) int {
	res, cur := 0, 0
	for _, v := range nums {
		if v == 1 {
			cur++
		} else {
			res = max(res, cur)
			cur = 0
		}
	}
	res = max(res, cur)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
