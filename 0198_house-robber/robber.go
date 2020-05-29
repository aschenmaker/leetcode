package main

import "fmt"

func main() {
	house := []int{1, 2, 3, 1}
	fmt.Println(rob(house))
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func rob2(nums []int) int {
	preMax := 0
	curMax := 0
	for _, v := range nums {
		temp := curMax
		if preMax+v > curMax {
			curMax = preMax + v
		}
		preMax = temp
	}
	return curMax
}
