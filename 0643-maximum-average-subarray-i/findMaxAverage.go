package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{8860, -853, 6534, 4477, -4589, 8646, -6155, -5577, -1656, -5779, -2619, -8604, -1358, -8009, 4983, 7063, 3104, -1560, 4080, 2763, 5616, -2375, 2848, 1394, -7173, -5225, -8244, -809, 8025, -4072, -4391, -9579, 1407, 6700, 2421, -6685, 5481, -1732, -8892, -6645, 3077, 3287, -4149, 8701, -4393, -9070, -1777, 2237, -3253, -506, -4931, -7366, -8132, 5406, -6300, -275, -1908, 67, 3569, 1433, -7262, -437, 8303, 4498, -379, 3054, -6285, 4203, 6908, 4433, 3077, 2288, 9733, -8067, 3007, 9725, 9669, 1362, -2561, -4225, 5442, -9006, -429, 160, -9234, -4444, 3586, -5711, -9506, -79, -4418, -4348, -5891}
	k := 93
	fmt.Println(findMaxAverage(nums, k))
}

func findMaxAverage(nums []int, k int) float64 {
	if k == 1 {
		sort.Ints(nums)
		return float64(nums[len(nums)-1])
	}

	preSum := make([]int, len(nums)+1)
	preSum[0] = 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		preSum[i+1] = sum
	}
	fmt.Println(preSum)
	ans := -1 << 63
	for i, j := 0, k; j <= len(nums); {
		if preSum[j]-preSum[i] > ans {
			ans = preSum[j] - preSum[i]
		}
		i++
		j++
	}
	fmt.Println(ans)
	return float64(ans) / float64(k)
}

func moreBetterSolution(nums []int, k int) float64 {
	sum := 0
	for _, v := range nums[:k] {
		sum += v
	}
	maxSum := sum
	for i := k; i < len(nums); i++ {
		sum = sum - nums[i-k] + nums[i]
		maxSum = max(maxSum, sum)
	}
	return float64(maxSum) / float64(k)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
