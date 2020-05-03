package main

import "fmt"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	result := maxSubArray(nums)
	fmt.Println(result)
}

func maxSubArray(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	dp := nums[0]
	re := dp
	for _, each := range nums[1:] {
		if dp > 0 {
			dp += each
		} else {
			dp = each
		}
		if dp > re {
			re = dp
		}
	}
	return re

}
